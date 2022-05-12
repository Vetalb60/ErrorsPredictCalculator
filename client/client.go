// Package main
//
//	________client.go________
//
//	Swagger client. Edit by alex.green. DO NOT EDIT.
//
package main

import (
	schd "exec/client/args"
	"exec/client/client/requests"
	"exec/client/logger"
	"flag"
	"fmt"
)

func main() {
	args := new(schd.Args)

	args.Ip_server = flag.String("ip", "127.0.0.1", "Ip address of server.")
	args.Port = flag.String("port", "33333", "Used port.")
	args.Handle_type = flag.String("req", schd.GET_, "Type of request. "+
		"\nValid types: \n\t--get\n\t--healthz\n\t--errors\n\t--upload\n\t--delete\n\t--list\n")
	args.Method_type = flag.String("method", schd.LADDER_METHOD_, "Method of linear predict."+
		"\nValid types: \n\t--covariation\n\t--autocorrelation\n\t--ladder\n")
	args.File_upload = flag.String("file", "", "Upload file in database. Default is empty string.")
	args.Implement_id = flag.Int("id", -1, "Id of implemented file. Default is non-file.")

	flag.Parse()

	req := new(requests.Requests)
	req.Init(*args.Ip_server, *args.Port, "http")

	log_file, log_main, err := logger.InitLogging("./client/logger/logs/info.log", "swagger-client.main")
	defer log_file.Close()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	schd.ArgumentsScheduler(log_main, *args, req)
}
