/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/cmd"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/db"
)

func main() {

	//loading the config
	cfg:= config.LoadConfig()
	if cfg == nil {
		fmt.Fprintln(os.Stderr,"failed to load the config")
	}

	//setting up mysql
	database, err:= db.MySQL(&cfg.MySQL)
	if err!=nil {
		fmt.Fprintln(os.Stderr,"failed to connect to the database")
		os.Exit(1)
	}
	cmd.SetDB(database)
	cmd.SetConfig(cfg)

	//setting cli 
	cmd.Execute()
}
