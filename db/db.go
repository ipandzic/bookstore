package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"

	"main/config"
)

var Db *sql.DB
var err error

func ConnectDB() {
	boil.DebugMode = true

	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, config.DbSSLMode)
	Db, err = sql.Open(config.DbEngine, connectionInfo)
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	Db.Close()
}

func GetDB() *sql.DB {
	return Db
}
