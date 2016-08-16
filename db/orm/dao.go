package orm

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// BaseDAO is default dao interface
type BaseDAO struct {
	tx    DBTX
	table *BaseTable
}

// NewBaseDAO create new base DAO
func NewBaseDAO(bt *BaseTable, bdtx DBTX) *BaseDAO {
	return &BaseDAO{
		tx:    bdtx,
		table: bt,
	}
}

// FindAll obtain all articles from database
func (dao *BaseDAO) FindAll() (*sqlx.Rows, error) {
	return dao.tx.Queryx(dao.table.selectSQL)
}

// FindByID obtain article of given ID from database
func (dao *BaseDAO) FindByID(id int64) *sqlx.Row {
	return dao.tx.QueryRowx(dao.table.selectByIDSQL, id)
}

// Insert store given articles to database
func (dao *BaseDAO) Insert(entity interface{}) (int64, error) {
	var (
		res sql.Result
		err error
	)
	if res, err = dao.tx.NamedExec(dao.table.insertSQL, entity); err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertWithID store given articles to database (It persist with id from entity)
func (dao *BaseDAO) InsertWithID(entity interface{}) error {
	if _, err := dao.tx.NamedExec(dao.table.insertWithIDSQL, entity); err != nil {
		return err
	}
	return nil
}

// Update data of article
func (dao *BaseDAO) Update(entity interface{}) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = dao.tx.NamedExec(dao.table.updateByIDSQL, entity); err != nil {
		return err
	}
	if count, err = res.RowsAffected(); err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds)", count)
	}
	return nil
}

// Delete remove specyfic record
func (dao *BaseDAO) Delete(id int64) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = dao.tx.NamedExec(dao.table.deleteByIDSQL, &IDContainer{id}); err != nil {
		return err
	}
	if count, err = res.RowsAffected(); err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

// CreateTable add new table to a database
func (dao *BaseDAO) CreateTable() error {
	dao.tx.MustExec(dao.table.createSQL)
	return nil
}
