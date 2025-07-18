package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/models"
	"gorm.io/gorm"
)


/*
AddTask adds a new task to the database using GORM.
It uses context to allow better control (e.g., timeout, cancellation).
Returns an error if the task is nil or if the DB insert fails. */

func AddTask(ctx context.Context, db *gorm.DB, task *models.Task) error{
	// Defensive check to ensure task object is not nil.
	if task == nil {
		return errors.New("cannot add a nil task")
	}
	
	// db.WithContext ensures context is passed to GORM — helpful for timeouts, cancellation, etc.
	err:= db.WithContext(ctx).Create(task).Error
	if err!=nil {
		// Wrap the error with more context so it’s easier to debug.
		return fmt.Errorf("error in adding task: %w", err)
	}
	return nil
}

/*
DeleteTask removes a task from the database by its ID.
Uses context for potential timeout/cancel and GORM for DB interaction.
Returns an error if ID is 0 or if the delete operation fails. */

func DeleteTask(ctx context.Context, db *gorm.DB, id uint) error {
	//ID can't be zero as the DB is set to autoincrement of ID's
	if id == 0 {
		return errors.New("id doesn't exist")
	}
	// Perform the delete operation for the given ID.
	err:= db.WithContext(ctx).Delete(&models.Task{}, id).Error
	if err!= nil{
		// Wrap the error with more context so it’s easier to debug.
		return fmt.Errorf("error in deleting task %d: %w",id, err)
	}
	return nil
}

/*
MarkComplete marks a task as completed by updating two fields:
1. "completed" set to true
2. "completed_at" set to current timestamp
Takes the task ID and applies an update using GORM's map update feature.
*/

func MarkComplete(ctx context.Context, db *gorm.DB, id uint) error {
	
	// Prevent invalid ID usage
	if id == 0 {
		return errors.New("id doesn't exist")
	}
	
	// Prepare the fields to update as a map[string]any
	update:= map[string] any{
		"completed": true,
		"completed_at": time.Now(),
	}

	// Apply the update only to the task with matching ID
	err:= db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Updates(update).Error
	
	if err!=nil {
		// Wrap the error for better error tracing
		return fmt.Errorf("error in task %d marked complete: %w", id, err)
	}
	return nil
}


// GetAllTasks fetches all tasks from the database, regardless of completion status.
// It uses GORM's Find method to retrieve all records from the tasks table.
func GetAllTasks(ctx context.Context, db *gorm.DB) ([]models.Task,error){
	var tasks []models.Task

	// Use context-aware DB call to safely fetch all task records
	err:= db.WithContext(ctx).Find(&tasks).Error
	if err!=nil {
		return nil, fmt.Errorf("getting tasks: %w",err)
	}
	return tasks, nil
}

// GetPendingTasks fetches only the tasks that are not marked as completed (i.e., pending).
// This is useful for showing only the active tasks to the user.
func PendingTasks(ctx context.Context, db *gorm.DB) ([]models.Task, error){
	var tasks []models.Task

	// Fetch tasks where the "completed" field is false
	err:= db.WithContext(ctx).Where("completed = ?", false).Find(&tasks).Error
	if err!= nil{
		return nil, fmt.Errorf("error in getting pending tasks: %w",err)
	}
	return tasks, nil
}

// UpdateTask updates the title and/or description of a task identified by its ID.
// It dynamically builds an update map based on non-empty inputs, allowing partial updates.
// If neither title nor description is provided, it returns an error.
func UpdateTask(ctx context.Context, db *gorm.DB, id uint, title string, description string) error {
	if id == 0{
		return errors.New("id cannot be zero")
	}
	// Dynamically build the update map based on what user wants to change
	update:= make(map[string] interface{})

	if title != ""{
		update["title"] = title
	}

	if description!= ""{
		update["description"] = description
	}

	// If no fields provided to update, return an error
	if len(update) == 0{
		return errors.New("no updates have been made")
	}
	// Perform the update query where task ID matches
	err:= db.WithContext(ctx).Model(&models.Task{}).Where("id = ?",id).Updates(update).Error
	if err!= nil{
		return fmt.Errorf("update %d failed: %w", id, err)
	}
	return nil
}