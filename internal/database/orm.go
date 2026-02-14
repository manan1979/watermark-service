package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	TicketID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(400)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermark string `gorm:"type:varchar(100)"`
}

func Init(host, port, user, dbname, pass string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
