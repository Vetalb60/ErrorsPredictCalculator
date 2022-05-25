// Package init
//
//	________main.go________
//
//	Generate swagger code. Edit by alex.green. DO NOT EDIT.
//
package main

import (
	"CourseWork/api/databases/mysql"
	"CourseWork/gen/restapi/handlers"
	"CourseWork/logger"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"CourseWork/gen/restapi"
	"CourseWork/gen/restapi/operations"
)

func main() {

	handlers_ := new(handlers.Handlers)
	mysql_ := new(mysql.Mysql)

	log_file, log_main, err := logger.InitLogging("logger/logs/info.log", "swagger-server.init")
	defer log_file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := mysql.Auth()
	if err != nil {
		log.Fatalln(err)
	}

	if err = mysql_.InitDB(log_main, conn); err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewPredictionAlgorithmsServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Example service"
	parser.LongDescription = "Example service"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	api.GetHandler = handlers_.GetHandler(log_main)
	api.GetHealthzHandler = handlers_.GetHealthzHandler(log_main)
	api.PostUploadHandler = handlers_.PostUploadHandler(log_main, mysql_)
	api.DeleteDeleteIDHandler = handlers_.DeleteFilesIDHandler(log_main, mysql_)
	api.GetListHandler = handlers_.GetListHandler(log_main, mysql_)
	api.PostMethodIDHandler = handlers_.PostMethodHandler(log_main, mysql_)

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
