// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllDataHandlerFunc turns a function with the right signature into a get all data handler
type GetAllDataHandlerFunc func(GetAllDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllDataHandlerFunc) Handle(params GetAllDataParams) middleware.Responder {
	return fn(params)
}

// GetAllDataHandler interface for that can handle valid get all data params
type GetAllDataHandler interface {
	Handle(GetAllDataParams) middleware.Responder
}

// NewGetAllData creates a new http.Handler for the get all data operation
func NewGetAllData(ctx *middleware.Context, handler GetAllDataHandler) *GetAllData {
	return &GetAllData{Context: ctx, Handler: handler}
}

/* GetAllData swagger:route GET /data data getAllData

Get all defi data

*/
type GetAllData struct {
	Context *middleware.Context
	Handler GetAllDataHandler
}

func (o *GetAllData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllDataParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
