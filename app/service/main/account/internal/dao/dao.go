package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mall/library/database/sql"
)

type Dao struct {
	db *gorm.DB
}

func New(c *sql.Config) (dao *Dao) {
	dao = &Dao{
		db: sql.NewSql(c),
	}
	return
}

func (d *Dao) Close() {
	d.db.Close()
}
