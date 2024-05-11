package Scanners

import "database/sql"

type RowScanner func(rows *sql.Rows) (interface{}, error)

func Query(db *sql.DB, query string, scanner RowScanner, args ...interface{}) ([]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var results []interface{}
	for rows.Next() {
		result, err := scanner(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
