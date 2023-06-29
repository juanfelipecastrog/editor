package db

import (
	"fmt"
	"gold/internal/model"
	"log"
)

func Migrate() {
	db := Connection()
	err := db.AutoMigrate(&model.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
	}

	fmt.Println("Migration done...")

}
