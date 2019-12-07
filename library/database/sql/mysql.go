package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Config struct {
	DSN string
}

func NewSql(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Fatal(fmt.Sprintf("db from %s opened failed", c.DSN))
	}
	return
}
