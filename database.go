package database

import "database/sql"

type Row map[string]interface{}

func (row Row) String(key string) string {
	value, _ := row[key].([]byte)
	return string(value)
}

func (row Row) Int64(key string) int64 {
	value, _ := row[key].(int64)
	return value
}

func (row Row) Bool(key string) bool {
	value, ok := row[key].(int64)
	if !ok {
		return false
	}

	return value == 1
}

func ToMap(row *sql.Rows) Row {
	cols, _ := row.Columns()

	m := make(map[string]interface{})

	for row.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := row.Scan(columnPointers...); err != nil {
			panic(err)
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
	}

	return Row(m)
}
