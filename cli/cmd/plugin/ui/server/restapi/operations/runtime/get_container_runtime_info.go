// Code generated by go-swagger; DO NOT EDIT.

package runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetContainerRuntimeInfoHandlerFunc turns a function with the right signature into a get container runtime info handler
type GetContainerRuntimeInfoHandlerFunc func(GetContainerRuntimeInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContainerRuntimeInfoHandlerFunc) Handle(params GetContainerRuntimeInfoParams) middleware.Responder {
	return fn(params)
}

// GetContainerRuntimeInfoHandler interface for that can handle valid get container runtime info params
type GetContainerRuntimeInfoHandler interface {
	Handle(GetContainerRuntimeInfoParams) middleware.Responder
}

// NewGetContainerRuntimeInfo creates a new http.Handler for the get container runtime info operation
func NewGetContainerRuntimeInfo(ctx *middleware.Context, handler GetContainerRuntimeInfoHandler) *GetContainerRuntimeInfo {
	return &GetContainerRuntimeInfo{Context: ctx, Handler: handler}
}

/* GetContainerRuntimeInfo swagger:route GET /api/containerruntime runtime getContainerRuntimeInfo

Get container runtime information

*/
type GetContainerRuntimeInfo struct {
	Context *middleware.Context
	Handler GetContainerRuntimeInfoHandler
}

func (o *GetContainerRuntimeInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetContainerRuntimeInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
