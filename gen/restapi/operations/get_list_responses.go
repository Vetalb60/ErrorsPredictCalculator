// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"CourseWork/gen/models"
)

// GetListOKCode is the HTTP code returned for type GetListOK
const GetListOKCode int = 200

/*GetListOK returns a list

swagger:response getListOK
*/
type GetListOK struct {

	/*
	  In: Body
	*/
	Payload models.Array `json:"body,omitempty"`
}

// NewGetListOK creates GetListOK with default headers values
func NewGetListOK() *GetListOK {

	return &GetListOK{}
}

// WithPayload adds the payload to the get list o k response
func (o *GetListOK) WithPayload(payload models.Array) *GetListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list o k response
func (o *GetListOK) SetPayload(payload models.Array) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Array{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetListBadRequestCode is the HTTP code returned for type GetListBadRequest
const GetListBadRequestCode int = 400

/*GetListBadRequest Bad request

swagger:response getListBadRequest
*/
type GetListBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetListBadRequest creates GetListBadRequest with default headers values
func NewGetListBadRequest() *GetListBadRequest {

	return &GetListBadRequest{}
}

// WithPayload adds the payload to the get list bad request response
func (o *GetListBadRequest) WithPayload(payload *models.Error) *GetListBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list bad request response
func (o *GetListBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetListInternalServerErrorCode is the HTTP code returned for type GetListInternalServerError
const GetListInternalServerErrorCode int = 500

/*GetListInternalServerError Internal server error

swagger:response getListInternalServerError
*/
type GetListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetListInternalServerError creates GetListInternalServerError with default headers values
func NewGetListInternalServerError() *GetListInternalServerError {

	return &GetListInternalServerError{}
}

// WithPayload adds the payload to the get list internal server error response
func (o *GetListInternalServerError) WithPayload(payload *models.Error) *GetListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get list internal server error response
func (o *GetListInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
