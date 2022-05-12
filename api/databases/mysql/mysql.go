// Package mysql
//
//	________mysql.go________
//
//	The mysql module is designed to interact with MYSQL database.
//	The module defines the functions of connection initialization, adding a file to a table,
//	deleting a file, and getting a description of tables.
//	Event logging is also organized into a separate file.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package mysql

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"
)

// Bad id return.
const (
	BAD_ID = -1
)

// Data base configuration.
type InitDBInfo struct {
	User_name_     string
	Password_      string
	Ip_address_    string
	Port_          string
	Database_name_ string
	Table_name_    string
}

// Time location format.
const (
	RFC822 = "02 Jan 06 15:04 MST"
)

// MYSQL struct with descriptors and requests inside.
type Mysql struct {
	desc descriptors
	req  requests
	db   *sql.DB
}

type descriptors struct{}
type requests struct{}

// InitDB
//	A function for connecting a client to a mysql database.
//	The parameter must pass user information.
//	The connection is also checked by a sequence of packets transmitted to the database server.
func (mysql *Mysql) InitDB(log *log.Logger, info InitDBInfo) error {
	log.Printf("Initialize database. \n\tUser:%s, \n\tIP:%s:%s, \n\tDatabase:%s, \n\tTable:%s",
		info.User_name_,
		info.Ip_address_,
		info.Port_,
		info.Database_name_,
		info.Table_name_)

	//	Creating a connection: "Username: Password @ tcp (IP: Port) / Database? Charset = utf8»
	path := strings.Join([]string{
		info.User_name_, ":",
		info.Password_, "@tcp(",
		info.Ip_address_, ":",
		info.Port_, ")/",
		info.Database_name_, "?charset=utf8"}, "")

	//	Initialize existing sql commands.
	initSqlCommands(info)

	//	Open the database, the first is the driver name, so import: _ "github.com/go-sql-driver/mysql ".
	db, err := sql.Open("mysql", path)
	if err != nil {
		return err
	}

	//	Check the connection using ping.
	if err := db.Ping(); err != nil {
		return err
	}

	mysql.db = db

	log.Printf("The connection to the database was successful.")

	return nil
}

// AddFileInDB Mysql
// 	The file of adding a new file to the database.
//	The full file name must be passed to parameters.
func (mysql *Mysql) AddFileInDB(log *log.Logger, file_path string) (int64, error) {
	var id int64

	//	Open the file by its path.
	file_desc, err := mysql.desc.openFile(log, file_path)

	log.Printf("Add file %s into database...", file_desc.Meta_.File_name_)

	//	We send a request to add a file to the database.
	id, err = mysql.req.dbExecQuery(mysql.db, _INSERT_FILE_IN_FILES_DB,
		file_desc.Meta_.File_name_,
		file_desc.Meta_.File_size_,
		file_desc.Meta_.Date_of_insert_,
		file_desc.Data_)
	if err != nil {
		log.Printf(err.Error())
		return BAD_ID, err
	}

	log.Printf("File %s was successful added.", file_desc.Meta_.File_name_)

	return id, nil
}

// GetFileFromDB Mysql
//	The function of getting a string with a description of the file from the database.
//	The file id must be passed in the parameters.
func (mysql *Mysql) GetFileFromDB(log *log.Logger, file_id int64) (*FileDescriptor, error) {
	log.Printf("Getting descriptor with id: %d.", file_id)

	//	Send a request to get a table row by file id.
	row := mysql.req.dbGetQueryRow(mysql.db, _SELECT_ROW_FROM_FILES_DB, file_id)

	desc := new(FileDescriptor)
	desc.Data_ = make([]byte, desc.Meta_.File_size_)

	err := row.Scan(&desc.Meta_.Id_,
		&desc.Meta_.File_name_,
		&desc.Meta_.File_size_,
		&desc.Meta_.Date_of_insert_,
		&desc.Data_)
	if err != nil {
		return nil, err
	}

	return desc, nil
}

// DeleteFileFromBD
//	The function of deleting a file from the database.
//	The file id must be passed in the parameters.
func (mysql *Mysql) DeleteFileFromBD(log *log.Logger, file_id int64) (int64, error) {
	log.Printf("Delete file with id: %d.", file_id)

	//	Send a deletion request by file id.
	id, err := mysql.req.dbExecQuery(mysql.db, _DELETE_ROW_FROM_FILES_DB, file_id)
	if err != nil {
		return BAD_ID, err
	}

	return id, nil
}

// GetFilesList
//	The function of getting all the list of all files from the database.
func (mysql *Mysql) GetFilesList(log *log.Logger) ([]FileDescriptor, error) {
	log.Printf("Getting database list...")

	//	Send a request to get all the elements of the tables.
	rows, err := mysql.req.dbGetQuery(mysql.db, _SELECT_ALL_ROWS_FROM_FILES_DB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filesMetaInfo := []FileDescriptor{}

	for rows.Next() {
		files := FileDescriptor{}
		err := rows.Scan(&files.Meta_.Id_,
			&files.Meta_.File_name_,
			&files.Meta_.File_size_,
			&files.Meta_.Date_of_insert_,
			&files.Data_)

		if err != nil {
			return nil, err
		}
		filesMetaInfo = append(filesMetaInfo, files)
	}

	return filesMetaInfo, err
}

// auth
//	Function for providing authentication and authorization in the database.
func Auth() (InitDBInfo, error) {
	var connection InitDBInfo
	if err := setEnvironmentVar(&connection); err != nil {
		return InitDBInfo{}, err
	}

	return connection, nil
}

func setEnvironmentVar(connection *InitDBInfo) error {
	var ok bool
	if connection.Ip_address_, ok = os.LookupEnv("MYSQL_POD_IP"); ok != true {
		return errors.New("error get environment MYSQL_POD_IP")
	}
	if connection.Password_, ok = os.LookupEnv("MYSQL_ROOT_PASSWORD"); ok != true {
		return errors.New("error get environment MYSQL_ROOT_PASSWORD")
	}
	if connection.Database_name_, ok = os.LookupEnv("MYSQL_DATABASE_NAME"); ok != true {
		return errors.New("error get environment MYSQL_DATABASE_NAME")
	}
	if connection.Port_, ok = os.LookupEnv("MYSQL_PORT"); ok != true {
		return errors.New("error get environment MYSQL_PORT")
	}
	if connection.User_name_, ok = os.LookupEnv("MYSQL_USER"); ok != true {
		return errors.New("error get environment MYSQL_USER")
	}
	if connection.Table_name_, ok = os.LookupEnv("MYSQL_TABLE_NAME"); ok != true {
		return errors.New("error get environment MYSQL_TABLE_NAME")
	}

	return nil
}
