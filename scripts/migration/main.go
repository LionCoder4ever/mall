package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mall/internal/app/account"
	"mall/internal/pkg/conf"
	"mall/internal/pkg/database/sql"
)

func main() {
	flag.Parse()
	if err := conf.Load(); err != nil {
		panic(fmt.Sprintf("conf load failed %s", err.Error()))
	}
	// init service
	//&conf.Conf
	db := sql.NewSql(conf.Conf.MySQL)
	db.AutoMigrate(&account.Account{})
	db.Model(&account.Account{}).AddUniqueIndex("uid", "uid")

	db.Close()
}
