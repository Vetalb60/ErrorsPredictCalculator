// Package logger
//
//	________logger.go________
//
//	A package that provides information about the environment and system statuses.
//
// 	Copyright 2022 Alex Green. All rights reserved.
//
package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

const (
	_ROOT_DIR_NAME_ = "CourseWork"
)

// InitLogging
//	Logger initialization function.
//	In the parameters, pass the file to be recorded and the prefix that will be recorded in the logs.
func InitLogging(log_name string, prefix string) (*os.File, *log.Logger, error) {
	file, err := os.OpenFile(log_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	infoLog := log.New(file, "[ "+prefix+" ]: ", log.Ldate|log.Ltime)

	infoLog.SetFlags(log.LstdFlags | log.Lshortfile)

	return file, infoLog, nil
}

// GetMainPath
//	The function of getting the path of the root directory in the current runtime.
func GetMainPath() string {
	_, path_to_main, _, ok := runtime.Caller(0)

	if ok {
		for path_to_main[len(path_to_main)-10:] != _ROOT_DIR_NAME_ {
			path_to_main = path.Join(path.Dir(path_to_main), "") //	The the main function file directory
		}
	} else {
		path_to_main = "./"
	}

	return path_to_main
}

func RecordEvent(log *log.Logger, event string) {
	log.Printf(event)
	fmt.Println(event)
}
