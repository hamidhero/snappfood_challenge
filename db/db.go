package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"snappfood_challenge/models"
)

var db *sql.DB
var err error

// Init establishes a connection to postgres database
func Init() {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("USERNAME"), viper.GetString("dev.db.db_name"),
		os.Getenv("PASSWORD"),
		viper.GetString("dev.db.host"), viper.GetString("dev.db.port"))
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// gives an instance of db
func GetDB() *sql.DB {
	return db
}

// create table if not exists
func CreateOrderTable() {
	orderModel := models.Order{}
	if err := orderModel.Create(db); err != nil {
		panic(err)
	}
}
