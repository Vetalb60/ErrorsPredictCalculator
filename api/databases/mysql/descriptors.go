// Package mysql
//
//	________descriptors.go________
//
//	Contains functions for interacting with file descriptors.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package mysql

import (
	"log"
	"os"
	"time"
)

//	File Descriptor
type FileDescriptor struct {
	Meta_ FilesMetaInfo
	Data_ []byte
}

//	Meta information of file from database.
type FilesMetaInfo struct {
	Id_             int64
	File_name_      string
	File_size_      int64
	Date_of_insert_ string
}

//	Function of opening a file by path.
func (descr *descriptors) openFile(log *log.Logger, file_path string) (*FileDescriptor, error) {
	file_desc := new(FileDescriptor)

	file, err := os.Open(file_path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	err = descr.fillFileDescriptor(log, file, file_desc)
	if err != nil {
		return nil, err
	}

	return file_desc, nil
}

func (descr *descriptors) fillFileMetaInfo(file_name string, file_size int64, desc *FileDescriptor) {
	desc.Meta_.File_name_ = file_name
	desc.Meta_.File_size_ = file_size
	desc.Meta_.Date_of_insert_ = time.Now().Format(RFC822)
}

func (descr *descriptors) fillFileDescriptor(log *log.Logger, file_stream *os.File, desc *FileDescriptor) error {
	file_info, _ := file_stream.Stat()

	descr.fillFileMetaInfo(file_info.Name(), file_info.Size(), desc)
	log.Printf("Reading file %s...", desc.Meta_.File_name_)

	file_data := make([]byte, desc.Meta_.File_size_)

	number, err := file_stream.Read(file_data)
	log.Printf("%d bytes was read.\n", number)
	if err != nil {
		return err
	}

	desc.Data_ = file_data

	return nil
}
