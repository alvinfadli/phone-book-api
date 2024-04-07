package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	"fmt"

	"phone-book-api/helpers"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {

	helpers.LoadEnvVariables()
    dbUser := helpers.GetEnvVariable("DB_USER")
	dbPassword := helpers.GetEnvVariable("DB_PASSWORD")
	dbHost := helpers.GetEnvVariable("DB_HOST")
	dbPort := helpers.GetEnvVariable("DB_PORT")
	dbName := helpers.GetEnvVariable("DB_NAME")
	fmt.Println(dbUser)

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return db, nil
}