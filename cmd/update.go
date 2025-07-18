/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"


	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the title or description of a task by its ID",
	Long: `The update command lets you modify the title or description of a task using its ID.

You can choose to update either or both fields by providing the appropriate flags.

Examples:
  To update only the title:
    todo update --id 3 --title "New Title"

  To update only the description
	todo update --id 3 --desc "Updated description"

  To update both title and description:
    todo update --id 3 --title "New Title" --desc "Updated description"

Note: If neither flag is provided, nothing will be updated.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetUint("id")
	title, _ := cmd.Flags().GetString("title")
	description, _ := cmd.Flags().GetString("desc")

	if title == "" && description == "" {
		fmt.Println("Please provide at least one field to update: --title or --desc")
		return
	}

	cfg := config.LoadConfig()
	database, err := db.MySQL(&cfg.MySQL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	err = db.UpdateTask(cmd.Context(), database, id, title, description)
	if err != nil {
		log.Fatalf("Failed to update the task: %v\n", err)
	}

	fmt.Println("Updated the task successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().UintP("id", "i", 0, "gets the id")
	updateCmd.Flags().StringP("title","t","","title to be updated")
	updateCmd.Flags().StringP("desc","d","","description to be updated")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
