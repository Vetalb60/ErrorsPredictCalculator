// Package handlers
//
//	________get_list_handler.go________
//
//	Retrieves the list of current files from the database.
//	CRUD: GET
//	Path: /files
//
//	Swagger handlers. Edit by alex.green
//
package handlers

import (
	"CourseWork/api/databases/mysql"
	"CourseWork/gen/models"
	"CourseWork/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (handler *Handlers) GetListHandler(log *log.Logger, mysql *mysql.Mysql) operations.GetListHandlerFunc {
	return func(params operations.GetListParams) middleware.Responder {
		list, err := mysql.GetFilesList(log)
		if err != nil {
			err_str := err.Error()
			log.Printf(err.Error())
			return operations.NewPostMethodInternalServerError().WithPayload(&models.Error{
				nil,
				&err_str,
			})
		}

		retMessage := make([]*models.ArrayItems0, len(list))

		for index, elem := range list {
			retMessage[index] = &models.ArrayItems0{
				ID:         elem.Meta_.Id_,
				FileSize:   elem.Meta_.File_size_,
				InsertDate: elem.Meta_.Date_of_insert_,
				Name:       elem.Meta_.File_name_,
			}
		}

		return operations.NewGetListOK().WithPayload(models.Array{
			retMessage,
		})
	}
}
