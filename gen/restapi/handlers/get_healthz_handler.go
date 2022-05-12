// Package handlers
//
//	________get_healthz_handler.go________
//
//	Returns the status of the server.
//	CRUD: GET
//	Path: /healthz
//
//	Swagger handlers. Edit by alex.green
//
package handlers

import (
	"CourseWork/gen/models"
	"CourseWork/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (handler *Handlers) GetHealthzHandler(log *log.Logger) operations.GetHealthzHandlerFunc {
	return func(params operations.GetHealthzParams) middleware.Responder {
		log.Printf("Getting local status from server...")
		ok_message := "OK, server is available."
		return operations.NewGetHealthzOK().WithPayload(&models.Health{Message: &ok_message})
	}
}
