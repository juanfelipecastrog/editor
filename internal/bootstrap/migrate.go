package bootstrap

import (
	"gold/config"
	"gold/internal/db"
)

func Migrate() {
	config.Set()

	db.Connect()

	db.Migrate()
}
