// Package handlers
//
//	________post_upload_handler.go________
//
//	Upload file at database by the file name passed in the parameters.
//	CRUD: POST
//	Path: /upload
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
	"io"
	"log"
	"net/http"
	"os"
)

func (handler *Handlers) PostUploadHandler(log *log.Logger, mysql_ *mysql.Mysql) operations.PostUploadHandlerFunc {
	return func(params operations.PostUploadParams) middleware.Responder {
		id, err := uploadFile(log, params.HTTPRequest, mysql_)
		if err != nil {
			err_str := err.Error()
			return operations.NewPostUploadOK().WithPayload(&models.PostUpload{
				Message: &err_str,
			})
		}

		message := fmt.Sprintf("File was successfully upload on the server! ID:%d", id)
		return operations.NewPostUploadOK().WithPayload(&models.PostUpload{
			Message: &message,
		})
	}
}

func uploadFile(log *log.Logger, r *http.Request, mysql_ *mysql.Mysql) (int64, error) {
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	r.Close = true
	defer file.Close()
	if err != nil {
		log.Printf("Error: Retrieving the File.")
		return -1, err
	}

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create("gen/upload/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		log.Printf("Error: %s", err)
		return -1, err
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err = io.Copy(dst, file); err != nil {
		log.Printf("Error: %s", err)
		return -1, err
	}

	id, err := mysql_.AddFileInDB(log, "gen/upload/"+handler.Filename)
	if err != nil {
		log.Printf("Error: %s", err)
		return -1, err
	}

	log.Printf("File was successfuly upload on the server with id: %d!", id)

	return id, nil
}
