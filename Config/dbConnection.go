package config

import (
	"fmt"
	"os"

	// "github.com/gofiber/fiber/v2"
	middleware "github.com/juanfmange/fiberapp/Middleware"
	models "github.com/juanfmange/fiberapp/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Db struct {
	dbhost     string
	dbuser     string
	dbpassword string
	dbname     string
	dbport     string
}

func Connection() {
	middleware.LoadEnv()
	dbc := Db{
		dbhost:     os.Getenv("HOSTNAME"),
		dbuser:     os.Getenv("USR"),
		dbpassword: os.Getenv("PWD"),
		dbname:     os.Getenv("DB_NAME"),
		dbport:     os.Getenv("DB_PORT"),
	}

	fmt.Println(dbc.dbuser, dbc.dbpassword, dbc.dbhost, dbc.dbport, dbc.dbname)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbc.dbuser, dbc.dbpassword, dbc.dbhost, dbc.dbport, dbc.dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection failed")
	}

	DB = db
	fmt.Println("db connection succesfully")

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
	)
}
