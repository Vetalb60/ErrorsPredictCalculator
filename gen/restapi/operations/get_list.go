// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetListHandlerFunc turns a function with the right signature into a get list handler
type GetListHandlerFunc func(GetListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetListHandlerFunc) Handle(params GetListParams) middleware.Responder {
	return fn(params)
}

// GetListHandler interface for that can handle valid get list params
type GetListHandler interface {
	Handle(GetListParams) middleware.Responder
}

// NewGetList creates a new http.Handler for the get list operation
func NewGetList(ctx *middleware.Context, handler GetListHandler) *GetList {
	return &GetList{Context: ctx, Handler: handler}
}

/* GetList swagger:route GET /files getList

Get files list from database.

*/
type GetList struct {
	Context *middleware.Context
	Handler GetListHandler
}

func (o *GetList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
