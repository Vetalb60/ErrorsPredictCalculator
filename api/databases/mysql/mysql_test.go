// Package mysql
//
//	________mysql_test.go________
//
//	Tests to check the correct functioning of the mysql package.
//	The essence of the tests is to connect to the database according to user parameters,
//	add a test file, get the necessary descriptions and delete the test file from the database.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package mysql

import (
	"CourseWork/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	mysql := new(Mysql)
	var id int64
	main_path := logger.GetMainPath()
	file, log, err := logger.InitLogging(main_path+"/logger/logs/database.log", "TestMysql")
	defer file.Close()
	if err != nil {
		t.Error(err.Error())
	}

	//-----------------------------------------------------------------------------------------

	/*
	 *	Test of add file in database.
	 */
	t.Run("success", func(t *testing.T) {
		connection := InitDBInfo{"root", "12345", "192.168.10.10", "3306", "mysql", "blobs", "files"}

		err = mysql.InitDB(log, connection)

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}

		//	The file "file.test" is being added, which is located in the folder with the test source.

		id, err = mysql.AddFileInDB(log, "file.test")

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}

		//	If the library function worked incorrectly,
		//	the AddFileInDB method will return BAD_ID.
		if id == BAD_ID {
			t.Error("Error:BAD_ID file is not added!")
			log.Printf("Error:BAD_ID file is not added!")
		}
	})
	//-----------------------------------------------------------------------------------------

	/*
	 *	A test to compare the descriptor obtained from the GetFileFromDB function.
	 */

	type args struct {
		id_ int64
	}
	var tests = []struct {
		name           string
		args           args
		wantDescriptor FileDescriptor
	}{
		{
			name: "success",
			args: args{
				id_: id,
			},
			wantDescriptor: FileDescriptor{
				Meta_: FilesMetaInfo{
					File_name_:      "file.test",
					File_size_:      34,
					Date_of_insert_: time.Now().Format(RFC822),
				},
				Data_: []byte("File attributes cannot be changed!"),
			},
		},
	}

	//	The name, size, modification date and binary data of the file are compared.
	tt := tests[0]

	t.Run(tt.name, func(t *testing.T) {
		gotDescriptor, err := mysql.GetFileFromDB(log, id)

		if gotDescriptor.Meta_.File_name_ != tt.wantDescriptor.Meta_.File_name_ ||
			gotDescriptor.Meta_.File_size_ != tt.wantDescriptor.Meta_.File_size_ ||
			gotDescriptor.Meta_.Date_of_insert_ != tt.wantDescriptor.Meta_.Date_of_insert_ ||
			string(gotDescriptor.Data_) != string(tt.wantDescriptor.Data_) {
			t.Error("Method GetFileFromBd(*sql.bd,int) got failed descriptor from database!Got:Name ", gotDescriptor.Meta_.File_name_,
				" | Size ", gotDescriptor.Meta_.File_size_, " | Date of insert ", gotDescriptor.Meta_.Date_of_insert_, " | \nWant:Name ",
				tt.wantDescriptor.Meta_.File_name_, " | Size ", tt.wantDescriptor.Meta_.File_size_, " | Date of insert ", tt.wantDescriptor.Meta_.Date_of_insert_, " | ")
		}

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}
	})

	//-----------------------------------------------------------------------------------------

	/*
	 *	The list of identifiers from descriptors is compared:
	 *  obtained from the functions GetFilesList(sql.DB) and from getFilesId(sql.DB).
	 */
	t.Run("success", func(t *testing.T) {
		list, err := mysql.GetFilesList(log)

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}

		test_list, err := getFilesID(mysql.db)

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}

		//	The 'id' field is scanned from the rows obtained by the getFilesId(sql.DB) method.
		for i := 0; test_list.Next() && i < len(list); i++ {
			var scan_id int64
			if err := test_list.Scan(&scan_id); err != nil {
				t.Error(err)
				log.Printf(err.Error())
			}

			if list[i].Meta_.Id_ != scan_id {
				t.Error("[ERROR]:The identifiers of the received descriptions do not match. " +
					"Check the get File sId(*sql.DB) method.")
				log.Printf("[ERROR]:The identifiers of the received descriptions do not match. " +
					"Check the get File sId(*sql.DB) method.")
			}
		}
	})
	//-----------------------------------------------------------------------------------------

	/*
	 *	The previously added test file "file.test" is removed from the mysql database.
	 *	A list of files from the database is requested and the deletion of the file by
	 *	the DeleteFileFromBD(sql.DB,int) function is checked for correctness.
	 */
	t.Run("success", func(t *testing.T) {
		_, err = mysql.DeleteFileFromBD(log, id)

		test_list, err := getFilesID(mysql.db)

		if err != nil {
			t.Error(err.Error())
			log.Printf(err.Error())
		}

		for i := 0; test_list.Next(); i++ {
			var scan_id int64
			if err := test_list.Scan(&scan_id); err != nil {
				t.Error(err)
				log.Printf(err.Error())
			}

			if id == scan_id {
				t.Error("[ERROR]:The identifiers of the received descriptions do not match. " +
					"Check the get File sId(*sql.DB) method.")
				log.Printf("[ERROR]:The identifiers of the received descriptions do not match. " +
					"Check the get File sId(*sql.DB) method.")
			}
		}
	})
	//-----------------------------------------------------------------------------------------
}

func getFilesID(db *sql.DB) (*sql.Rows, error) {
	list, err := db.Query("select id from blobs.files")

	if err != nil {
		return nil, err
	}

	return list, nil
}
