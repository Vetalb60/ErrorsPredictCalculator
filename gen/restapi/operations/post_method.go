// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMethodHandlerFunc turns a function with the right signature into a post method handler
type PostMethodHandlerFunc func(PostMethodParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMethodHandlerFunc) Handle(params PostMethodParams) middleware.Responder {
	return fn(params)
}

// PostMethodHandler interface for that can handle valid post method params
type PostMethodHandler interface {
	Handle(PostMethodParams) middleware.Responder
}

// NewPostMethod creates a new http.Handler for the post method operation
func NewPostMethod(ctx *middleware.Context, handler PostMethodHandler) *PostMethod {
	return &PostMethod{Context: ctx, Handler: handler}
}

/* PostMethod swagger:route POST /method postMethod

Creates a new user.

*/
type PostMethod struct {
	Context *middleware.Context
	Handler PostMethodHandler
}

func (o *PostMethod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostMethodParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
