// Package mysql
//
//	________requests.go________
//
//	Contains functions for processing queries to the MYSQL database.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package mysql

import "database/sql"

func (req *requests) dbExecQuery(db *sql.DB, request_type string, args ...any) (int64, error) {
	result, err := db.Exec(request_type, args...)
	if err != nil {
		return BAD_ID, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (req *requests) dbGetQueryRow(db *sql.DB, request_type string, args ...any) *sql.Row {
	row := db.QueryRow(request_type, args...)

	return row
}

func (req *requests) dbGetQuery(db *sql.DB, request_type string, args ...any) (*sql.Rows, error) {
	row, err := db.Query(request_type, args...)
	if err != nil {
		return nil, err
	}

	return row, nil
}
