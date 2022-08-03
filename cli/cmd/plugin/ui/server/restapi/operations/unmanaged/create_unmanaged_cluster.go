// Code generated by go-swagger; DO NOT EDIT.

package unmanaged

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateUnmanagedClusterHandlerFunc turns a function with the right signature into a create unmanaged cluster handler
type CreateUnmanagedClusterHandlerFunc func(CreateUnmanagedClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUnmanagedClusterHandlerFunc) Handle(params CreateUnmanagedClusterParams) middleware.Responder {
	return fn(params)
}

// CreateUnmanagedClusterHandler interface for that can handle valid create unmanaged cluster params
type CreateUnmanagedClusterHandler interface {
	Handle(CreateUnmanagedClusterParams) middleware.Responder
}

// NewCreateUnmanagedCluster creates a new http.Handler for the create unmanaged cluster operation
func NewCreateUnmanagedCluster(ctx *middleware.Context, handler CreateUnmanagedClusterHandler) *CreateUnmanagedCluster {
	return &CreateUnmanagedCluster{Context: ctx, Handler: handler}
}

/*CreateUnmanagedCluster swagger:route POST /api/unmanaged unmanaged createUnmanagedCluster

Create a new unmanaged cluster

*/
type CreateUnmanagedCluster struct {
	Context *middleware.Context
	Handler CreateUnmanagedClusterHandler
}

func (o *CreateUnmanagedCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateUnmanagedClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}