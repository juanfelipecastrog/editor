package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

var Configs Config

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&Configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}

func Get() Config {
	return Configs
}
