package bootstrap

import (
	"editor/config"
	"editor/internal/db"
)

func Migrate() {
	config.Set()

	db.Connect()

	db.Migrate()
}
