/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/models"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks, including completed and pending.",
	Long: `The list command shows all tasks in your task manager, displaying both completed and pending tasks.
Use this command to quickly view the status of your to-dos and keep track of your progress.

example:
	taskcli list            # Show all tasks
	taskcli list --pending  # Show only pending tasks
`,
	Run: func(cmd *cobra.Command, args []string) {
		showpending,_:= cmd.Flags().GetBool("pending")
		
		var task []models.Task
		var err error

		if showpending {
			task, err = db.PendingTasks(cmd.Context(),DB)
		}else {
			task, err = db.GetAllTasks(cmd.Context(),DB)
		}

		if err != nil {
			log.Fatalf("failed to fetch tasks: %v", err)
		}
		utils.PrintTasks(task)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("pending", "p", false, "gets all pending tasks")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
