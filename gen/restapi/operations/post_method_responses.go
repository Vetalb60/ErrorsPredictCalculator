// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"CourseWork/gen/models"
)

// PostMethodOKCode is the HTTP code returned for type PostMethodOK
const PostMethodOKCode int = 200

/*PostMethodOK OK

swagger:response postMethodOK
*/
type PostMethodOK struct {

	/*
	  In: Body
	*/
	Payload *models.Energy `json:"body,omitempty"`
}

// NewPostMethodOK creates PostMethodOK with default headers values
func NewPostMethodOK() *PostMethodOK {

	return &PostMethodOK{}
}

// WithPayload adds the payload to the post method o k response
func (o *PostMethodOK) WithPayload(payload *models.Energy) *PostMethodOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method o k response
func (o *PostMethodOK) SetPayload(payload *models.Energy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMethodBadRequestCode is the HTTP code returned for type PostMethodBadRequest
const PostMethodBadRequestCode int = 400

/*PostMethodBadRequest Bad request

swagger:response postMethodBadRequest
*/
type PostMethodBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMethodBadRequest creates PostMethodBadRequest with default headers values
func NewPostMethodBadRequest() *PostMethodBadRequest {

	return &PostMethodBadRequest{}
}

// WithPayload adds the payload to the post method bad request response
func (o *PostMethodBadRequest) WithPayload(payload *models.Error) *PostMethodBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method bad request response
func (o *PostMethodBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMethodInternalServerErrorCode is the HTTP code returned for type PostMethodInternalServerError
const PostMethodInternalServerErrorCode int = 500

/*PostMethodInternalServerError Internal server error

swagger:response postMethodInternalServerError
*/
type PostMethodInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMethodInternalServerError creates PostMethodInternalServerError with default headers values
func NewPostMethodInternalServerError() *PostMethodInternalServerError {

	return &PostMethodInternalServerError{}
}

// WithPayload adds the payload to the post method internal server error response
func (o *PostMethodInternalServerError) WithPayload(payload *models.Error) *PostMethodInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post method internal server error response
func (o *PostMethodInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMethodInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}