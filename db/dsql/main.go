package dsql

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/types"
)

type DSQL struct{}

func NewDSQL() db.DSQL {
	return db.DSQL(DSQL{})
}

// NewSelectSQL create new select query
func (dsql DSQL) NewSelectSQL(table string, fields []string) (string, error) {
	sql := "SELECT id"
	for _, row := range fields {
		sql += ", " + row
	}
	return sql + " FROM " + table, nil
}

// NewSelectWhereSQL create new select query
func (dsql DSQL) NewSelectWhereSQL(table string, fields []string, where string) (string, error) {
	selectSQL, err := dsql.NewSelectSQL(table, fields)
	if err != nil {
		return "", err
	}
	return selectSQL + " WHERE " + where, nil
}

// NewInsertSQL create new insert sql
func (dsql DSQL) NewInsertSQL(table string, fields []string) (string, error) {
	sqlUpdate := "INSERT INTO " + table + " ("
	sqlValues := "VALUES ("
	for i, row := range fields {
		if i == 0 {
			sqlUpdate += "" + row
			sqlValues += ":" + row
		} else {
			sqlUpdate += ", " + row
			sqlValues += ", :" + row
		}
	}
	return sqlUpdate + ") " + sqlValues + ")", nil
}

// NewUpdateSQL create new update sql
func (dsql DSQL) NewUpdateSQL(table string, fields []string) (string, error) {
	sqlUpdate := "UPDATE " + table + " SET "
	for i, row := range fields {
		if i == 0 {
			sqlUpdate += row + " = :" + row
		} else {
			sqlUpdate += ", " + row + " = :" + row
		}
	}
	return sqlUpdate, nil
}

// NewUpdateWhereSQL create new select query
func (dsql DSQL) NewUpdateWhereSQL(table string, fields []string, where string) (string, error) {
	sql, err := dsql.NewUpdateSQL(table, fields)
	if err != nil {
		return "", err
	}
	return sql + " WHERE " + where, nil
}

// NewDeleteSQL create new delete query
func (dsql DSQL) NewDeleteSQL(table string) (string, error) {
	return "DELETE FROM " + table, nil
}

// NewDeleteWhereSQL create new delete query with where restrict
func (dsql DSQL) NewDeleteWhereSQL(table string, where string) (string, error) {
	deleteSQL, err := dsql.NewDeleteSQL(table)
	if err != nil {
		return "", err
	}
	return deleteSQL + " WHERE " + where, nil
}

// NewSQLType create a sql type description
func (dsql DSQL) NewSQLType(t types.CustomType) (string, error) {
	s := t.SQLType()
	if t.HasAttribute(types.Unique) {
		s += " UNIQUE"
	}
	if t.HasAttribute(types.Primary) {
		s += " PRIMARY KEY"
	}
	if t.HasAttribute(types.NotNull) {
		s += " NOT NULL"
	}
	if t.HasAttribute(types.Required) {
		s += " NOT NULL"
	}
	return s, nil
}

// NewCreateSQL create new create table query
func (dsql DSQL) NewCreateSQL(table string, types map[string]types.CustomType) (string, error) {
	var i = 1
	sql := "CREATE TABLE " + table + " (\n"
	sql += "id INTEGER PRIMARY KEY AUTOINCREMENT"
	for name, customType := range types {
		if i > 0 {
			sql += ",\n"
		}
		typeStr, err := dsql.NewSQLType(customType)
		if err != nil {
			return "", err
		}
		sql += name + " " + typeStr
		i++
	}
	return sql + ")", nil
}
