// Package handlers
//
//	________default_handler.go________
//
//	Default handler. Returns the status of the server.
//	CRUD: GET
//	Path: /
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

func (handler *Handlers) GetHandler(log *log.Logger) operations.GetHandlerFunc {
	return func(params operations.GetParams) middleware.Responder {
		log.Printf("Getting local status from server...")
		ok_message := "OK, server is available."
		return operations.NewGetHealthzOK().WithPayload(&models.Health{Message: &ok_message})
	}
}
