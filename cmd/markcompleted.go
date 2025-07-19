/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
	"github.com/spf13/cobra"
)

// markcompletedCmd represents the markcompleted command
var markcompletedCmd = &cobra.Command{
	Use:   "markcompleted",
	Short: "Mark a task as completed using its ID.",
	Long: `The markcompleted command allows you to mark a task as completed by providing its task ID.

This helps you update the task status when you've finished working on it. 
Once marked, the task will appear as completed when listed.

Examples:
  taskcli markcompleted 3    # Marks the task with ID 3 as completed
  taskcli markcompleted 12   # Marks task ID 12 as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		
		id, _:= cmd.Flags().GetUint("id")

		err := db.MarkComplete(cmd.Context(),DB, id)
		if err!=nil {
			log.Fatalf("failed to mark the task: %v\n",err)
		}
		fmt.Println("Task Completed")
	},
}

func init() {
	rootCmd.AddCommand(markcompletedCmd)

	markcompletedCmd.Flags().UintP("id", "i", 0, "marks the task completed based on id (id required)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markcompletedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markcompletedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
