/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/cmd"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
)

func main() {
	cmd.Execute()

	//loading the config
	cfg:= config.LoadConfig()
	log.Println("Environment",cfg.Env)
	log.Println("Database ",cfg.MySQL.DBname)

	//setting up mysql
	_, err:= db.MySQL(&cfg.MySQL)
	if err!=nil {
		slog.Error("failed to connect to the database","error", err)
		os.Exit(1)
	}
}
