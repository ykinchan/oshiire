package	service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/ykinchan/Oshiire/backend/model"
	"log"
	"os"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := os.Getenv("DB_USER") + ":" +  os.Getenv("DB_PASS") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/oshiire?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Block))
    fmt.Println("init data base ok")
}