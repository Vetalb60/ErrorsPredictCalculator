// Package args
//
//	________args_handler.go________
//
//	Switch is a request handler from the user.
//
//	Copyright 2022 Alex Green. All rights reserved.
//
package args

import (
	"exec/client/client/requests"
	"exec/client/logger"
	_ "exec/client/logger"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

//	Request references.
const (
	GET_    = "get"
	HEALTH_ = "healthz"
	ERRORS_ = "errors"
	UPLOAD_ = "upload"
	DELETE_ = "delete"
	LIST_   = "list"
)

const (
	AUTOCORRELATION_METHOD_ = "autocorrelation" // method define
	COVARIATION_METHOD_     = "covariation"     // method define
	LADDER_METHOD_          = "ladder"          // method define
)

//	User arguments.
type Args struct {
	Ip_server    *string
	Port         *string
	Handle_type  *string
	Method_type  *string
	File_upload  *string
	Implement_id *int
}

// ArgumentsScheduler
//	Switch handler.
func ArgumentsScheduler(log *log.Logger, args Args, req *requests.Requests) int {
	logger.RecordEvent(log, "Try to send request on the server...")
	switch *args.Handle_type {
	case GET_:
		logger.RecordEvent(log, "Send GET request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		res, err := req.Get()
		if err != nil {
			logger.RecordEvent(log, err.Error())
			break
		}
		logger.RecordEvent(log, res)

		break
	case HEALTH_:
		logger.RecordEvent(log, "Send GET request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		res, err := req.GetHealthz()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		logger.RecordEvent(log, "OK, the response was received.")
		logger.RecordEvent(log, res)
		break
	case ERRORS_:
		if int64(*args.Implement_id) == -1 {
			logger.RecordEvent(log, "Non-file id. Input id file in database.")
			break
		}
		logger.RecordEvent(log, "Send GET request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		switch *args.Method_type {
		case COVARIATION_METHOD_:
			errs, err := req.GetErrorByMethod("covariation", int64(*args.Implement_id))
			if err != nil {
				logger.RecordEvent(log, err.Error())
				break
			}
			logger.RecordEvent(log, "OK, the response was received.")
			fmt.Printf("Covariation errors:\n")
			for index, elem := range errs {
				fmt.Printf(" | %d \t| %.10f \t|\n", index+2, elem)
			}
			break
		case AUTOCORRELATION_METHOD_:
			errs, err := req.GetErrorByMethod("autocorrelation", int64(*args.Implement_id))
			if err != nil {
				logger.RecordEvent(log, err.Error())
				break
			}
			logger.RecordEvent(log, "OK, the response was received.")
			fmt.Printf("Autocorrealtion errors:\n")
			for index, elem := range errs {
				fmt.Printf(" | %d \t| %.10f \t|\n", index+2, elem)
			}
			break
		case LADDER_METHOD_:
			errs, err := req.GetErrorByMethod("ladder", int64(*args.Implement_id))
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			logger.RecordEvent(log, "OK, the response was received.")
			fmt.Printf("Ladder errors:\n")
			for index, elem := range errs {
				fmt.Printf(" | %d \t| %.10f \t|\n", index+2, elem)
			}
			break
		default:
			logger.RecordEvent(log, "[WARNING]:Enter method is not implement. Try input other.")
		}
		break
	case UPLOAD_:
		logger.RecordEvent(log, "Send POST request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		message, err := req.UploadFile(*args.File_upload)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		logger.RecordEvent(log, "OK, the response was received.")
		logger.RecordEvent(log, message)
		break
	case DELETE_:
		logger.RecordEvent(log, "Send DELETE request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		if int64(*args.Implement_id) == -1 {
			logger.RecordEvent(log, "Non-file id. Input id file in database.")
			break
		}
		message, err := req.DeleteFileById(int64(*args.Implement_id))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		logger.RecordEvent(log, "OK, the response was received.")
		logger.RecordEvent(log, message)
		break
	case LIST_:
		logger.RecordEvent(log, "Send GET request...")
		logger.RecordEvent(log, "Getting response by the calculating server...")
		info, err := req.GetFilesList()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		logger.RecordEvent(log, "OK, the response was received.")
		w := tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', tabwriter.AlignRight)
		for _, elem := range info {
			fmt.Fprintf(w, "File ID:\t %d \t|\t Name:\t %s \t|\t Size:\t %d \t|\t Insert Date: %s  |\n",
				elem.Id_, elem.File_name_, elem.File_size_, elem.Date_of_insert_)
		}
		w.Flush()
		break
	default:
		logger.RecordEvent(log, "[Warning]:No method to use.")
		logger.RecordEvent(log, "Try to input other method.")
	}

	return 0
}
