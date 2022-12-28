package	service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"/run/model"
	"log"
)

var DbEngine *xorm.DbEngine

func init() {
	driverName := "mysql"
	DsName := "root:root@(database:3306)/gin?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		lob.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Block))
    fmt.Println("init data base ok")
}