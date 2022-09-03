package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	Id        int64     `gorm:"column:Id"`
	Price     int64     `gorm:"column:Price"`
	Title     string    `gorm:"column:Title"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
}

// creates Order table if not exists when app is starting
func (o Order) Create(db *sql.DB) error {
	statement := `CREATE TABLE IF NOT EXISTS Orders 
		(Id int, Price bigint, Title varchar(255), CreatedAt timestamp);`

	if _, err := db.Exec(statement); err != nil {
		return err
	}
	return nil
}

// inserts new row to Order table
func (o Order) Insert(db *sql.DB) error {
	statement := fmt.Sprintf(`INSERT INTO Orders(Id, Price, Title, CreatedAt) VALUES(%d, %d, '%s', '%s')`,
		o.Id, o.Price, o.Title, time.Now().Format(time.RFC3339))

	if _, err := db.Exec(statement); err != nil {
		return err
	}
	return nil
}
