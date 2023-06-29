package bootstrap

import (
	"editor/config"
	"editor/internal/api"
	"editor/internal/db"
)

func Serve() {
	config.Set()
	db.Connect()
	server := api.NewServer()
	server.Run()
}
