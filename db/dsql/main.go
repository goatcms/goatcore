package dsql

import "github.com/goatcms/goat-core/types"

// NewSelectSQL create new select query
func NewSelectSQL(table string, fields []string) string {
	sql := "SELECT id"
	for _, row := range fields {
		sql += ", " + row
	}
	return sql + " FROM " + table
}

// NewSelectWhereSQL create new select query
func NewSelectWhereSQL(table string, fields []string, where string) string {
	selectSQL := NewSelectSQL(table, fields)
	return selectSQL + " WHERE " + where
}

// NewInsertSQL create new insert sql
func NewInsertSQL(table string, fields []string) string {
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
	return sqlUpdate + ") " + sqlValues + ")"
}

// NewUpdateSQL create new update sql
func NewUpdateSQL(table string, fields []string) string {
	sqlUpdate := "UPDATE " + table + " SET "
	for i, row := range fields {
		if i == 0 {
			sqlUpdate += row + " = :" + row
		} else {
			sqlUpdate += ", " + row + " = :" + row
		}
	}
	return sqlUpdate
}

// NewUpdateWhereSQL create new select query
func NewUpdateWhereSQL(table string, fields []string, where string) string {
	sql := NewUpdateSQL(table, fields)
	return sql + " WHERE " + where
}

// NewDeleteSQL create new delete query
func NewDeleteSQL(table string) string {
	return "DELETE FROM " + table
}

// NewDeleteWhereSQL create new delete query with where restrict
func NewDeleteWhereSQL(table string, where string) string {
	return NewDeleteSQL(table) + " WHERE " + where
}

// NewSQLType create a sql type description
func NewSQLType(t types.CustomType) string {
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
	return s
}

// NewCreateSQL create new create table query
func NewCreateSQL(table string, types map[string]types.CustomType) string {
	var i = 0
	sql := "CREATE TABLE " + table + " (\n"
	sql += "id INTEGER PRIMARY KEY AUTOINCREMENT,\n"
	for name, customType := range types {
		if i > 0 {
			sql += ",\n"
		}
		sql += name + " " + NewSQLType(customType)
		i++
	}
	return sql + ")"
}
