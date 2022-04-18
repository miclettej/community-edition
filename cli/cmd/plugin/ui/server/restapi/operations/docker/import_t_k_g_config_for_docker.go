// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ImportTKGConfigForDockerHandlerFunc turns a function with the right signature into a import t k g config for docker handler
type ImportTKGConfigForDockerHandlerFunc func(ImportTKGConfigForDockerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ImportTKGConfigForDockerHandlerFunc) Handle(params ImportTKGConfigForDockerParams) middleware.Responder {
	return fn(params)
}

// ImportTKGConfigForDockerHandler interface for that can handle valid import t k g config for docker params
type ImportTKGConfigForDockerHandler interface {
	Handle(ImportTKGConfigForDockerParams) middleware.Responder
}

// NewImportTKGConfigForDocker creates a new http.Handler for the import t k g config for docker operation
func NewImportTKGConfigForDocker(ctx *middleware.Context, handler ImportTKGConfigForDockerHandler) *ImportTKGConfigForDocker {
	return &ImportTKGConfigForDocker{Context: ctx, Handler: handler}
}

/*ImportTKGConfigForDocker swagger:route POST /api/provider/docker/config/import docker importTKGConfigForDocker

Generate TKG configuration object for docker

*/
type ImportTKGConfigForDocker struct {
	Context *middleware.Context
	Handler ImportTKGConfigForDockerHandler
}

func (o *ImportTKGConfigForDocker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewImportTKGConfigForDockerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
