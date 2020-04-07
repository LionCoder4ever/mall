package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mall/internal/pkg/database/sql"
)

var Db *gorm.DB

func NewDao(c *sql.Config) {
	Db = sql.NewSql(c)
}

func Close() {
	Db.Close()
}
