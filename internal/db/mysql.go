package db

//to run this program
/*
$env:ConfigPath="config/config.yaml"
go run main.go
*/

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/config"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySQL(cfg *config.MySQL) (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host,cfg.Port,cfg.DBname)

	//creating a custom logger
	NewLogger:= logger.New(slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
	logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel: logger.Error,
		IgnoreRecordNotFoundError: true,
		Colorful: false,
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewLogger,
	})
	if err!=nil {
		log.Panic("Failed to connect to database")
	}
	// slog.Info("Connected to the Database", "DataBase", cfg.DBname)

	err = db.AutoMigrate(&models.Task{})
	if err!=nil{
		fmt.Println("failed to migrate the database")
		return nil,err
	}
	// fmt.Println("migration complete")
	return db, nil
}