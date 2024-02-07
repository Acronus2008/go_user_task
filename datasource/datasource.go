package datasource

import (
	"api_rest_postgresql/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=cl-saas-db-dev.c3u65yjjhoev.us-east-1.rds.amazonaws.com user=saas_root password=3st3t1p03s3lr00t dbname=saas port=5432 sslmode=disable TimeZone=America/Santiago"
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
