// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"CourseWork/gen/models"
)

// PostMethodIDOKCode is the HTTP code returned for type PostMethodIDOK
const PostMethodIDOKCode int = 200

/*PostMethodIDOK OK

swagger:response postMethodIdOK
*/
type PostMethodIDOK struct {

	/*
	  In: Body
	*/
	Payload models.Energy `json:"body,omitempty"`
}

// NewPostMethodIDOK creates PostMethodIDOK with default headers values
func NewPostMethodIDOK() *PostMethodIDOK {

	return &PostMethodIDOK{}
}

// WithPayload adds the payload to the post method Id o k response
func (o *PostMethodIDOK) WithPayload(payload models.Energy) *PostMethodIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method Id o k response
func (o *PostMethodIDOK) SetPayload(payload models.Energy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Energy{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostMethodIDBadRequestCode is the HTTP code returned for type PostMethodIDBadRequest
const PostMethodIDBadRequestCode int = 400

/*PostMethodIDBadRequest Bad request

swagger:response postMethodIdBadRequest
*/
type PostMethodIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMethodIDBadRequest creates PostMethodIDBadRequest with default headers values
func NewPostMethodIDBadRequest() *PostMethodIDBadRequest {

	return &PostMethodIDBadRequest{}
}

// WithPayload adds the payload to the post method Id bad request response
func (o *PostMethodIDBadRequest) WithPayload(payload *models.Error) *PostMethodIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method Id bad request response
func (o *PostMethodIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMethodIDInternalServerErrorCode is the HTTP code returned for type PostMethodIDInternalServerError
const PostMethodIDInternalServerErrorCode int = 500

/*PostMethodIDInternalServerError Internal server error

swagger:response postMethodIdInternalServerError
*/
type PostMethodIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMethodIDInternalServerError creates PostMethodIDInternalServerError with default headers values
func NewPostMethodIDInternalServerError() *PostMethodIDInternalServerError {

	return &PostMethodIDInternalServerError{}
}

// WithPayload adds the payload to the post method Id internal server error response
func (o *PostMethodIDInternalServerError) WithPayload(payload *models.Error) *PostMethodIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method Id internal server error response
func (o *PostMethodIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
