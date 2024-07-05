package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBName     string
    DBUser     string
    DBPassword string
}

var AppConfig Config

func LoadConfig() error {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return err
    }

    AppConfig = Config{
        DBHost:     viper.GetString("DB_HOST"),
        DBPort:     viper.GetString("DB_PORT"),
        DBName:     viper.GetString("DB_DATABASE"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
    }

    return nil
}

func GetDBConnectionString() string {
    return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        AppConfig.DBUser,
        AppConfig.DBPassword,
        AppConfig.DBHost,
        AppConfig.DBPort,
        AppConfig.DBName,
    )
}