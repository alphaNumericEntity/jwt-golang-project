package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	// postgres://ttiemmga:lHJ5ATj7bUrdd1VfEygSXEbn-5PdAE7l@tiny.db.elephantsql.com/ttiemmga
	//dsn := "host=tiny.db.elephantsql.com user=ttiemmga password=lHJ5ATj7bUrdd1VfEygSXEbn-5PdAE7l dbname=ttiemmga port=5432 sslmode=disable"
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database %v", err.Error())
		os.Exit(2)
	}
}
