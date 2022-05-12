// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMethodIDHandlerFunc turns a function with the right signature into a post method ID handler
type PostMethodIDHandlerFunc func(PostMethodIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMethodIDHandlerFunc) Handle(params PostMethodIDParams) middleware.Responder {
	return fn(params)
}

// PostMethodIDHandler interface for that can handle valid post method ID params
type PostMethodIDHandler interface {
	Handle(PostMethodIDParams) middleware.Responder
}

// NewPostMethodID creates a new http.Handler for the post method ID operation
func NewPostMethodID(ctx *middleware.Context, handler PostMethodIDHandler) *PostMethodID {
	return &PostMethodID{Context: ctx, Handler: handler}
}

/* PostMethodID swagger:route POST /method/{id} postMethodId

Calculate energy of error selected method of predict.

*/
type PostMethodID struct {
	Context *middleware.Context
	Handler PostMethodIDHandler
}

func (o *PostMethodID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostMethodIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
