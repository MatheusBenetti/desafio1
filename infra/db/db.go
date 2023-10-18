package db

import (
	"github.com/MatheusBenetti/desafio1/domain/model"
	"gorm.io/driver/sqlite"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(basePath + "/../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectToDB(env string) *gorm.DB {

	var db *gorm.DB
	var err error

	db, err = gorm.Open(sqlite.Open("DBTYPE"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("AUTOMIGRATE") == "true" {
		db.AutoMigrate(&model.Product{})
	}

	return db
}
