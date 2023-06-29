package bootstrap

import (
	"gold/config"
	"gold/internal/api"
	"gold/internal/db"
)

func Serve() {
	config.Set()
	db.Connect()
	server := api.NewServer()
	server.Run()
}
