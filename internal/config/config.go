package config

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct{
	Env 		string 		`yaml:"env" env:"ENV" env-required:"true"`
	MySQL 		MySQL		`yaml:"mysql" env-required:"true"`
}

type MySQL struct{
	Host 		string 	`yaml:"host" env:"HOST" env-default:"localhost"`
	Port 		int 	`yaml:"port" env:"PORT" env-default:"3306"`
	User 		string 	`yaml:"user" env:"USER" env-default:"root"`
	Password 	string 	`yaml:"password" env:"PASSWORD"`
	DBname 		string 	`yaml:"dbname" env-required:"true"`
}


var (
	cfg *Config
	once sync.Once
)

func LoadConfig() *Config{
	once.Do(
		func () {
			configpath:= os.Getenv("ConfigPath")
			//setting the configuration path
			if configpath == ""{
				flagConfig:= flag.String("config", "", "path to the configuration")
				flag.Parse()
				configpath = *flagConfig

				if configpath == "" {
					log.Fatal("Set the Configuration Path or use --config in the cli")
				}
			}

			//checking the configuration exists
			_, err:= os.Stat(configpath)
			if err!=nil {
				log.Fatal("Config file doesn't exits")
			}

			//loading the config
			var load Config
			err = cleanenv.ReadConfig(configpath, &load)
			if err!= nil {
				log.Fatal("Failed to load the config",err.Error())
			}
			slog.Info("Configuration Loaded Successfully")
			cfg = &load
	})
	return cfg
}