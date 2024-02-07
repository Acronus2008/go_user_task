package datasource

import (
	"api_rest_postgresql/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=root password=root dbname=db port=5432 sslmode=disable TimeZone=America/Santiago"
var DB *gorm.DB

func DatasourceConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Datasource connected successfully!")
	}
}

func ExecuteMigrations() {
	DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})
}
