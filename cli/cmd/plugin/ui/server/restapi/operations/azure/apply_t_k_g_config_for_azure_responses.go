// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// ApplyTKGConfigForAzureOKCode is the HTTP code returned for type ApplyTKGConfigForAzureOK
const ApplyTKGConfigForAzureOKCode int = 200

/*ApplyTKGConfigForAzureOK Apply change to TKG configuration successfully

swagger:response applyTKGConfigForAzureOK
*/
type ApplyTKGConfigForAzureOK struct {

	/*
	  In: Body
	*/
	Payload *models.ConfigFileInfo `json:"body,omitempty"`
}

// NewApplyTKGConfigForAzureOK creates ApplyTKGConfigForAzureOK with default headers values
func NewApplyTKGConfigForAzureOK() *ApplyTKGConfigForAzureOK {

	return &ApplyTKGConfigForAzureOK{}
}

// WithPayload adds the payload to the apply t k g config for azure o k response
func (o *ApplyTKGConfigForAzureOK) WithPayload(payload *models.ConfigFileInfo) *ApplyTKGConfigForAzureOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for azure o k response
func (o *ApplyTKGConfigForAzureOK) SetPayload(payload *models.ConfigFileInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForAzureOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApplyTKGConfigForAzureBadRequestCode is the HTTP code returned for type ApplyTKGConfigForAzureBadRequest
const ApplyTKGConfigForAzureBadRequestCode int = 400

/*ApplyTKGConfigForAzureBadRequest Bad request

swagger:response applyTKGConfigForAzureBadRequest
*/
type ApplyTKGConfigForAzureBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewApplyTKGConfigForAzureBadRequest creates ApplyTKGConfigForAzureBadRequest with default headers values
func NewApplyTKGConfigForAzureBadRequest() *ApplyTKGConfigForAzureBadRequest {

	return &ApplyTKGConfigForAzureBadRequest{}
}

// WithPayload adds the payload to the apply t k g config for azure bad request response
func (o *ApplyTKGConfigForAzureBadRequest) WithPayload(payload *models.Error) *ApplyTKGConfigForAzureBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for azure bad request response
func (o *ApplyTKGConfigForAzureBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForAzureBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApplyTKGConfigForAzureUnauthorizedCode is the HTTP code returned for type ApplyTKGConfigForAzureUnauthorized
const ApplyTKGConfigForAzureUnauthorizedCode int = 401

/*ApplyTKGConfigForAzureUnauthorized Incorrect credentials

swagger:response applyTKGConfigForAzureUnauthorized
*/
type ApplyTKGConfigForAzureUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewApplyTKGConfigForAzureUnauthorized creates ApplyTKGConfigForAzureUnauthorized with default headers values
func NewApplyTKGConfigForAzureUnauthorized() *ApplyTKGConfigForAzureUnauthorized {

	return &ApplyTKGConfigForAzureUnauthorized{}
}

// WithPayload adds the payload to the apply t k g config for azure unauthorized response
func (o *ApplyTKGConfigForAzureUnauthorized) WithPayload(payload *models.Error) *ApplyTKGConfigForAzureUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for azure unauthorized response
func (o *ApplyTKGConfigForAzureUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForAzureUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ApplyTKGConfigForAzureInternalServerErrorCode is the HTTP code returned for type ApplyTKGConfigForAzureInternalServerError
const ApplyTKGConfigForAzureInternalServerErrorCode int = 500

/*ApplyTKGConfigForAzureInternalServerError Internal server error

swagger:response applyTKGConfigForAzureInternalServerError
*/
type ApplyTKGConfigForAzureInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewApplyTKGConfigForAzureInternalServerError creates ApplyTKGConfigForAzureInternalServerError with default headers values
func NewApplyTKGConfigForAzureInternalServerError() *ApplyTKGConfigForAzureInternalServerError {

	return &ApplyTKGConfigForAzureInternalServerError{}
}

// WithPayload adds the payload to the apply t k g config for azure internal server error response
func (o *ApplyTKGConfigForAzureInternalServerError) WithPayload(payload *models.Error) *ApplyTKGConfigForAzureInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the apply t k g config for azure internal server error response
func (o *ApplyTKGConfigForAzureInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ApplyTKGConfigForAzureInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
