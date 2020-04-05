package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mall/app/conf"
	"mall/app/internal/model"
	"mall/library/database/sql"
)

func main() {
	flag.Parse()
	if err := conf.Load(); err != nil {
		panic(fmt.Sprintf("conf load failed %s", err.Error()))
	}
	// init service
	//&conf.Conf
	db := sql.NewSql(conf.Conf.MySQL)
	db.AutoMigrate(&model.Account{})
	db.Model(&model.Account{}).AddUniqueIndex("uid", "uid")

	db.Close()
}
