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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySQL(cfg *config.MySQL) (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host,cfg.Port,cfg.DBname)

	//creating a custom logger
	NewLogger:= logger.New(slog.NewLogLogger(slog.Default().Handler(), slog.LevelInfo),
	logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel: logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful: true,
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewLogger,
	})
	if err!=nil {
		log.Panic("Failed to connect to database")
	}
	slog.Info("Connected to the Database", "DataBase", cfg.DBname)
	return db, nil
}