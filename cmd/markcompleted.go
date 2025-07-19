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

// markcompletedCmd represents the markcompleted command
var markcompletedCmd = &cobra.Command{
	Use:   "markcompleted",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		id, _:= cmd.Flags().GetUint("id")
		
		cfg := config.LoadConfig()
		database, err := db.MySQL(&cfg.MySQL)
		if err!=nil {
			log.Fatalf("failed to connect the database: %v", err)
		}

		err = db.MarkComplete(cmd.Context(),database, id)
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
