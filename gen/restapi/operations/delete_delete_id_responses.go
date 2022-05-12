// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"CourseWork/gen/models"
)

// DeleteDeleteIDOKCode is the HTTP code returned for type DeleteDeleteIDOK
const DeleteDeleteIDOKCode int = 200

/*DeleteDeleteIDOK successful response

swagger:response deleteDeleteIdOK
*/
type DeleteDeleteIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteKey `json:"body,omitempty"`
}

// NewDeleteDeleteIDOK creates DeleteDeleteIDOK with default headers values
func NewDeleteDeleteIDOK() *DeleteDeleteIDOK {

	return &DeleteDeleteIDOK{}
}

// WithPayload adds the payload to the delete delete Id o k response
func (o *DeleteDeleteIDOK) WithPayload(payload *models.DeleteKey) *DeleteDeleteIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete delete Id o k response
func (o *DeleteDeleteIDOK) SetPayload(payload *models.DeleteKey) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeleteIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDeleteIDBadRequestCode is the HTTP code returned for type DeleteDeleteIDBadRequest
const DeleteDeleteIDBadRequestCode int = 400

/*DeleteDeleteIDBadRequest Bad request

swagger:response deleteDeleteIdBadRequest
*/
type DeleteDeleteIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeleteIDBadRequest creates DeleteDeleteIDBadRequest with default headers values
func NewDeleteDeleteIDBadRequest() *DeleteDeleteIDBadRequest {

	return &DeleteDeleteIDBadRequest{}
}

// WithPayload adds the payload to the delete delete Id bad request response
func (o *DeleteDeleteIDBadRequest) WithPayload(payload *models.Error) *DeleteDeleteIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete delete Id bad request response
func (o *DeleteDeleteIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeleteIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDeleteIDInternalServerErrorCode is the HTTP code returned for type DeleteDeleteIDInternalServerError
const DeleteDeleteIDInternalServerErrorCode int = 500

/*DeleteDeleteIDInternalServerError Internal server error

swagger:response deleteDeleteIdInternalServerError
*/
type DeleteDeleteIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteDeleteIDInternalServerError creates DeleteDeleteIDInternalServerError with default headers values
func NewDeleteDeleteIDInternalServerError() *DeleteDeleteIDInternalServerError {

	return &DeleteDeleteIDInternalServerError{}
}

// WithPayload adds the payload to the delete delete Id internal server error response
func (o *DeleteDeleteIDInternalServerError) WithPayload(payload *models.Error) *DeleteDeleteIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete delete Id internal server error response
func (o *DeleteDeleteIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeleteIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}