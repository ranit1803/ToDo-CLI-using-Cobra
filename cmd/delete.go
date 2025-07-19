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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by its ID",
	Long: `Deletes a task from your to-do list using its unique ID.

You must provide the ID of the task you wish to delete using the --id flag.
For example:
  todo delete --id 3

This action is permanent and cannot be undone.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _:= cmd.Flags().GetUint("id")

		err := db.DeleteTask(cmd.Context(), DB, id)
		if err!=nil {
			log.Fatalf("failed to delete the task: %v\n", err)
		}
		fmt.Println("Deleted the task Successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintP("id", "i", 0, "deletes a task based on id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
