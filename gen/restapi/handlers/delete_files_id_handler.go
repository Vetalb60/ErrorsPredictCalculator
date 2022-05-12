// Package handlers
//
//	________delete_files_id_handler.go________
//
//	Deletes files from the database by the id passed in the parameters.
//	CRUD: DELETE
//	Path: /delete/{id}
//
//	Swagger handlers. Edit by alex.green
//
package handlers

import (
	"CourseWork/api/databases/mysql"
	"CourseWork/gen/models"
	"CourseWork/gen/restapi/operations"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (handlers *Handlers) DeleteFilesIDHandler(log *log.Logger, mysql *mysql.Mysql) operations.DeleteDeleteIDHandlerFunc {
	return func(params operations.DeleteDeleteIDParams) middleware.Responder {
		_, err := mysql.GetFileFromDB(log, params.ID)
		if err != nil {
			err_code := int64(404)
			err_str := fmt.Sprintf("Code %d%s", err_code, ". File is not found.")
			log.Printf(err.Error())
			return operations.NewGetHealthzOK().WithPayload(&models.Health{
				Message: &err_str,
			})
		}

		_, err = mysql.DeleteFileFromBD(log, params.ID)
		if err != nil {
			err_str := err.Error()
			log.Printf(err.Error())
			return operations.NewPostMethodInternalServerError().WithPayload(&models.Error{
				nil,
				&err_str,
			})
		}

		message := fmt.Sprintf("File with ID:%d was deleted from database!", params.ID)

		return operations.NewDeleteDeleteIDOK().WithPayload(&models.DeleteKey{
			Message: &message,
		})
	}
}
