/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	Config *config.Config
)

func SetDB(db *gorm.DB){
	DB = db
}

func SetConfig(cfg *config.Config){
	Config = cfg
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI ToDo application to manage tasks efficiently.",
	Long: `ToDo CLI is a terminal-based task manager built with Golang and Cobra.
It allows you to add, delete, update, and complete tasks directly from your terminal.

You can also view pending or completed tasks in a clean and organized format.

Example usage:

  todo add --title "Buy groceries" --desc "Milk, eggs, bread"
  todo markcomplete --id 3
  taskcli list            # Show all tasks
  taskcli list --pending  # Show only pending tasks
  todo delete --id 3`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ToDo-CLI-using-Cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


