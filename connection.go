package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
)

type DBConfig struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func getConnectionString() (string, string) {
	config, err := loadConfig("./config")
	if err != nil {
		log.Fatal("Cannot load config file!!! ", err)
	}
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	return str, config.DBDriver
}

func loadConfig(filepath string) (config DBConfig, err error) {
	viper.AddConfigPath(filepath)
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func MySQLConn() *gorm.DB {
	dbURL, dbDriver := getConnectionString()
	conn, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		log.Fatal("Connection Not Established!! ", err)
	}
	fmt.Print("Connected to :- ", dbDriver)
	return conn
}
