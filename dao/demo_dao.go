package dao

import (
	"database/sql"

	"github.com/my-gin-server/base/apperror"
	"github.com/my-gin-server/base/applog"
	"github.com/my-gin-server/base/db"
	"github.com/my-gin-server/model"
)

type DemoDao struct {
	db.DataAccessObject
}

func NewDemoDao() *DemoDao {
	return &DemoDao{DataAccessObject: db.MysqlAccessObj()}
}

func (dao *DemoDao) Create(demo *model.Demo) (int, error) {
	result, err := dao.DataAccessObject.Exec(
		"INSERT INTO demos (name, description, create_time) VALUES (?, ?, ?)",
		demo.Name, demo.Description, demo.CreateTime)
	if err != nil {
		applog.Error.Printf("Some errors should be hidden from apis: %s", err.Error())
		return 0, apperror.ErrDatabase
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func (dao *DemoDao) DeleteById(id int) error {
	result, err := dao.DataAccessObject.Exec("DELETE FROM demos WHERE id = ?", id)
	if err != nil {
		applog.Error.Print(err)
		return apperror.ErrDatabase
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return apperror.ErrNotFound
	}
	return nil
}

func (dat *DemoDao) DeleteWithTx(tx *sql.Tx, id int) error {
	result, err := tx.Exec("DELETE FROM demos WHERE id = ?", id)
	if err != nil {
		applog.Error.Print(err)
		return apperror.ErrDatabase
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return apperror.ErrNotFound
	}
	return nil
}

func (dao *DemoDao) List(offset int, limit int) ([]*model.Demo, int, error) {
	count, err := dao.DataAccessObject.Query("SELECT COUNT(*) FROM demos")
	if err != nil {
		applog.Error.Print(err)
		return nil, 0, apperror.ErrDatabase
	}
	defer count.Close()
	total := 0
	count.Next()
	count.Scan(&total)

	rows, err := dao.DataAccessObject.Query(
		"SELECT * FROM demos LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		applog.Error.Print(err)
		return nil, 0, apperror.ErrDatabase
	}
	defer rows.Close()

	var demos []*model.Demo
	for rows.Next() {
		var demo model.Demo
		err := rows.Scan(&demo.Id, &demo.Name, &demo.Description, &demo.CreateTime)
		if err != nil {
			applog.Error.Print(err)
			return nil, 0, apperror.ErrDatabase
		}
		demos = append(demos, &demo)
	}
	return demos, total, nil
}

func (dao *DemoDao) QueryById(id int) (*model.Demo, error) {
	rows, err := dao.DataAccessObject.Query(
		"SELECT * FROM demos WHERE id = ? LIMIT 1", id)
	if err != nil {
		applog.Error.Print(err)
		return nil, apperror.ErrDatabase
	}
	defer rows.Close()

	exist := rows.Next()
	if !exist {
		return nil, apperror.ErrNotFound
	}
	var demo model.Demo
	err = rows.Scan(&demo.Id, &demo.Name, &demo.Description, &demo.CreateTime)
	if err != nil {
		applog.Error.Print(err)
		return nil, apperror.ErrDatabase
	}
	return &demo, nil
}
