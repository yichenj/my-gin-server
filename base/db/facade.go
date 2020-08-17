package db

import (
	"database/sql"

	"github.com/my-gin-server/base/appconfig"
	"github.com/my-gin-server/base/applog"
)

var (
	mysql *MysqlDB
)

type DataAccessObject interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	BeginTx() (*sql.Tx, error)
}

func Init(config *appconfig.Config) {
	mysql = NewMysqlDB()

	mysqlCfg := config.Database.Mysql
	err := mysql.Init(mysqlCfg.Url, mysqlCfg.MaxIdleConns, mysqlCfg.MaxOpenConns)
	if err != nil {
		applog.Error.Panicf("Mysql init failed: %s", err)
	}
}

func Close() {
	mysql.Close()
}

func MysqlAccessObj() DataAccessObject {
	return mysql
}
