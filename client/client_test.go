// Package mysql
//
//	________client_test.go________
//
//	Tests to check the correct functioning swagger client.
//	The following checks are performed: the healthz of server,
//	uploading a file to the database, calculating errors,
//	checking the list of database files, deleting a file from the database.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package main

import (
	"exec/client/client/requests"
	"fmt"
	"strconv"
	"testing"
)

type returnArgs struct {
	OK     string
	Delete string
	Upload string
}

type args struct {
	ID int
}

func TestClient(t *testing.T) {
	req := new(requests.Requests)
	req.Init("127.0.0.1", "33333", "http")

	args_ := new(args)

	var tests = []struct {
		name_       string
		wantReturn_ returnArgs
	}{
		{
			name_: "GET, Healthz",
			wantReturn_: returnArgs{
				OK: "OK, server is available.",
			},
		},
		{
			name_: "Upload",
			wantReturn_: returnArgs{
				Upload: "File was successfully upload on the server!",
			},
		},
		{
			name_: "Errors",
		},
		{
			name_: "List",
		},
		{
			name_: "success",
		},
	}

	//	Healthz check.
	tt := tests[0]
	t.Run(tt.name_, func(t *testing.T) {
		res, err := req.Get()
		if err != nil {
			t.Error(err.Error())
		}
		if res != tt.wantReturn_.OK {
			t.Errorf("Error response.\nReturn: %s", res)
		}

		res, err = req.GetHealthz()
		if err != nil {
			t.Error(err.Error())
		}
		if res != tt.wantReturn_.OK {
			t.Errorf("Error response.\nReturn: %s", res)
		}
	})

	//	Upload file to db.
	tt = tests[1]
	t.Run(tt.name_, func(t *testing.T) {
		message, err := req.UploadFile("Jack_London.wav")
		if err != nil {
			t.Error(err.Error())
		}

		args_.ID, err = strconv.Atoi(message[47:])
		if err != nil {
			t.Error(err.Error())
		}
		tests[4].wantReturn_.Delete =
			fmt.Sprintf("File with ID:%d was deleted from database!", args_.ID)

		if message[:43] != tt.wantReturn_.Upload {
			t.Errorf("Error response.\nReturn: %s", message)
		}
	})

	//	Calculate errors check.
	tt = tests[2]
	t.Run(tt.name_, func(t *testing.T) {
		errs, err := req.GetErrorByMethod("covariation", int64(args_.ID))
		if err != nil {
			t.Error(err.Error())
		}
		if len(errs) == 0 {
			t.Error("Error in calculate errors of predict.")
		}

		errs, err = req.GetErrorByMethod("autocorrelation", int64(args_.ID))
		if err != nil {
			t.Error(err.Error())
		}
		if len(errs) == 0 {
			t.Error("Error in calculate errors of predict.")
		}

		errs, err = req.GetErrorByMethod("ladder", int64(args_.ID))
		if err != nil {
			t.Error(err.Error())
		}
		if len(errs) == 0 {
			t.Error("Error in calculate errors of predict.")
		}
	})

	//	Db list check.
	tt = tests[3]
	t.Run(tt.name_, func(t *testing.T) {
		info, err := req.GetFilesList()
		if err != nil {
			t.Error(err.Error())
		}

		for _, elem := range info {
			if elem.Id_ == int64(args_.ID) {
				return
			}
		}
		t.Error("No id upload file in list.")
	})

	//	Delete check.
	tt = tests[4]
	t.Run(tt.name_, func(t *testing.T) {
		message, err := req.DeleteFileById(int64(args_.ID))
		if err != nil {
			t.Error(err.Error())
		}
		if message != tt.wantReturn_.Delete {
			t.Errorf("File was not deleted.\nReturn: %s", message)
		}
	})
}
