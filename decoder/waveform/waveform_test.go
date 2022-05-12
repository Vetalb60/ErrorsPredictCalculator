// Package mysql
//
//	________waveform_test.go________
//
//	Tests to check the correct functioning of the waveform package.
//	The essence of the tests is to check the method of decoding the test file
//	from the platform file system and from the mysql database.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package waveform

import (
	"CourseWork/api/databases/mysql"
	"CourseWork/logger"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestWaveform(t *testing.T) {
	decoder := new(WAVDecoder)

	main_path := logger.GetMainPath()

	file, log_, err := logger.InitLogging(main_path+"/logger/logs/waveform.log", "TestWaveform")

	if err != nil {
		t.Error(err.Error())
	}

	defer file.Close()

	type args struct {
		file_path string
		file_desc mysql.FileDescriptor
	}
	var tests = []struct {
		name           string
		args           args
		wantDescriptor WAVFormat
	}{
		{
			name: "success",
			args: args{
				file_path: "file.test",
			},
			wantDescriptor: WAVFormat{
				Meta_: MetaInfo{
					AudioFormat_:   1,
					NumChannels_:   1,
					SampleRate_:    8000,
					ByteRate_:      8000,
					BlockAlign_:    1,
					BitsPerSample_: 8,
					FileSize_:      8,
				},
			},
		},
		{
			name: "success",
			wantDescriptor: WAVFormat{
				Meta_: MetaInfo{
					AudioFormat_:   1,
					NumChannels_:   1,
					SampleRate_:    8000,
					ByteRate_:      8000,
					BlockAlign_:    1,
					BitsPerSample_: 8,
					FileSize_:      8,
				},
			},
		},
	}

	//	Tests to check the decoding of a file by its path.
	tt := tests[0]
	t.Run(tt.name, func(t *testing.T) {
		desc, err := decoder.DecodeFile(log_, tt.args.file_path)
		if err != nil {
			t.Error(err)
			log_.Printf(err.Error())
		}

		if !compareWAV(desc, &tt.wantDescriptor) {
			t.Error("The descriptor obtained from the Decode File(*log_.Logger, string) (*WAV Format, error) does not match the want descriptor")
			log_.Printf("The descriptor obtained from the Decode File(*log_.Logger, string) (*WAV Format, error) does not match the want descriptor")
		}
	})

	//	Tests to check the decoding of a file from a database.
	tt = tests[1]
	t.Run(tt.name, func(t *testing.T) {
		file_desc, db, err := addFileInDB(log_)

		if err != nil {
			t.Error(err)
			log_.Printf(err.Error())
		}

		desc, err := decoder.DecodeFileFromDB(log_, file_desc)
		if err != nil {
			t.Error(err)
			log_.Printf(err.Error())
		}

		if !compareWAV(desc, &tt.wantDescriptor) {
			t.Error("The descriptor obtained from the Decode File(*log_.Logger, string) (*WAV Format, error) does not match the want descriptor")
			log_.Printf("The descriptor obtained from the Decode File(*log_.Logger, string) (*WAV Format, error) does not match the want descriptor")
		}

		err = deleteFromDB(log_, db, file_desc.Meta_.Id_)

	})

}

func addFileInDB(log *log.Logger) (*mysql.FileDescriptor, *mysql.Mysql, error) {
	database := new(mysql.Mysql)

	connection := mysql.InitDBInfo{"root", "12345", "192.168.10.10", "3306", "mysql", "blobs", "files"}

	err := database.InitDB(log, connection)

	if err != nil {
		return nil, nil, err
	}

	id, err := database.AddFileInDB(log, "file.test")

	if err != nil {
		return nil, nil, err
	}

	if id == mysql.BAD_ID {
		return nil, nil, errors.New("invalid id")
	}

	file_desc, err := database.GetFileFromDB(log, id)

	if err != nil {
		return nil, nil, err
	}

	return file_desc, database, err
}

func deleteFromDB(log *log.Logger, db *mysql.Mysql, file_id int64) error {
	_, err := db.DeleteFileFromBD(log, file_id)

	if err != nil {
		return err
	}

	return nil
}

func compareWAV(desc_1 *WAVFormat, desc_2 *WAVFormat) bool {
	if desc_1.Meta_.ByteRate_ == desc_2.Meta_.ByteRate_ &&
		desc_1.Meta_.AudioFormat_ == desc_2.Meta_.AudioFormat_ &&
		desc_1.Meta_.FileSize_ == desc_2.Meta_.FileSize_ &&
		desc_1.Meta_.BitsPerSample_ == desc_2.Meta_.BitsPerSample_ &&
		desc_1.Meta_.BlockAlign_ == desc_2.Meta_.BlockAlign_ &&
		desc_1.Meta_.NumChannels_ == desc_2.Meta_.NumChannels_ &&
		desc_1.Meta_.SampleRate_ == desc_2.Meta_.SampleRate_ {
		return true
	}

	return false
}
