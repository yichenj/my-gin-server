package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	db *sql.DB
}

func NewMysqlDB() *MysqlDB {
	return &MysqlDB{db: nil}
}

func (mysql *MysqlDB) Init(url string, maxIdleConns int, maxOpenConns int) error {
	var err error
	mysql.db, err = sql.Open("mysql", url+"?parseTime=true")
	if err != nil {
		return err
	}
	mysql.db.SetMaxIdleConns(maxIdleConns)
	mysql.db.SetMaxOpenConns(maxOpenConns)
	return nil
}

func (mysql *MysqlDB) Close() {
	if mysql.db != nil {
		mysql.db.Close()
	}
}

func (mysql *MysqlDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return mysql.db.Query(query, args...)
}

func (mysql *MysqlDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return mysql.db.Exec(query, args...)
}

func (mysql *MysqlDB) BeginTx() (*sql.Tx, error) {
	return mysql.db.Begin()
}
