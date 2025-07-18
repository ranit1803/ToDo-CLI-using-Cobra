/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your ToDo list",
	Long: `Adds a new task to your ToDo list with a title and optional description.

Example:
  todo add --title "Buy groceries" --desc "Milk, eggs, bread"

You can later view your tasks using:
  todo list     - to view all tasks
  todo pending  - to view only pending tasks
  todo done     - to view completed tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		title,_:= cmd.Flags().GetString("title")
		description,_:= cmd.Flags().GetString("desc")

		tasks:= &models.Task{
			Title: title,
			Description: description,
			Completed: false,
			CreatedAt: time.Now(),
			CompletedAt: nil,
			UpdatedAt: nil,
		}

		database, err:= db.MySQL(&config.LoadConfig().MySQL)
		if err!=nil {
			fmt.Printf("failed to connect to database: %v\n",err)
		}

		err= db.AddTask(cmd.Context(), database, tasks)
		if err!=nil {
			log.Fatalf("failed to add the task: %v", err)
		}

		fmt.Println("Task Added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("title","t", "", "Title of the task (required)")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().StringP("desc", "d", "", "Description of the task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
