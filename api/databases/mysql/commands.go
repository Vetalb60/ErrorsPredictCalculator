// Package mysql
//
//	________commands.go________
//
//	Sets commands for sending queries to the database.
//
// 	Copyright 2022 Alex Green. All rights reserved.
//
package mysql

var _INSERT_FILE_IN_FILES_DB string
var _SELECT_ROW_FROM_FILES_DB string
var _DELETE_ROW_FROM_FILES_DB string
var _SELECT_ALL_ROWS_FROM_FILES_DB string

// initSqlCommands
// Default string Mysql commands
func initSqlCommands(info InitDBInfo) {
	_INSERT_FILE_IN_FILES_DB = "insert into " + info.Database_name_ + "." + info.Table_name_ +
		"(file_name, file_size, date_of_insert, bin_data)" +
		" values (?, ?, ?, ?)"
	_SELECT_ROW_FROM_FILES_DB = "select * from " + info.Database_name_ + "." + info.Table_name_ + " where id = ?"
	_DELETE_ROW_FROM_FILES_DB = "delete from " + info.Database_name_ + "." + info.Table_name_ + " where id = ?"
	_SELECT_ALL_ROWS_FROM_FILES_DB = "select * from " + info.Database_name_ + "." + info.Table_name_
}
