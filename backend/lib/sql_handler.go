package lib

import (
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type SQLHandler struct {
	DB *gorm.DB
	Err error
}

var dbConn *SQLHandler

func DBOpen() {
	dbConn = NewSQLHandler()
}

func DBClose() {
	sqlDB, _ = dbConn.DB.DB()
	sqlDB.Close()
}

func NewSQLHandler() *SQLHandler {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(user, password, host, port)

	var db *gorm.DB
	var err error

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(100)
    //接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
    sqlDB.SetMaxOpenConns(100)
    //接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
    sqlDB.SetConnMaxLifetime(100 * time.Second)

	sqlHandler := new(SQLHandler)
    db.Logger.LogMode(4)
    sqlHandler.DB = db

    return sqlHandler
}

// GetDBConn ...
func GetDBConn() *SQLHandler {
    return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
    dbConn.DB = dbConn.DB.Begin()
    return dbConn.DB
}

// Rollback ...
func RollBack() {
    dbConn.DB.Rollback()
}