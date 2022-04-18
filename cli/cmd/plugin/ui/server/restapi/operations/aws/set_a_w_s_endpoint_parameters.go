// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// NewSetAWSEndpointParams creates a new SetAWSEndpointParams object
// no default values defined in spec.
func NewSetAWSEndpointParams() SetAWSEndpointParams {

	return SetAWSEndpointParams{}
}

// SetAWSEndpointParams contains all the bound params for the set a w s endpoint operation
// typically these are obtained from a http.Request
//
// swagger:parameters setAWSEndpoint
type SetAWSEndpointParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*AWS account parameters
	  In: body
	*/
	AccountParams *models.AWSAccountParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetAWSEndpointParams() beforehand.
func (o *SetAWSEndpointParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AWSAccountParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("accountParams", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.AccountParams = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
