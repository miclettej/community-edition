// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetAzureInstanceTypesOKCode is the HTTP code returned for type GetAzureInstanceTypesOK
const GetAzureInstanceTypesOKCode int = 200

/*GetAzureInstanceTypesOK Successful retrieval of Azure instance Types

swagger:response getAzureInstanceTypesOK
*/
type GetAzureInstanceTypesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AzureInstanceType `json:"body,omitempty"`
}

// NewGetAzureInstanceTypesOK creates GetAzureInstanceTypesOK with default headers values
func NewGetAzureInstanceTypesOK() *GetAzureInstanceTypesOK {

	return &GetAzureInstanceTypesOK{}
}

// WithPayload adds the payload to the get azure instance types o k response
func (o *GetAzureInstanceTypesOK) WithPayload(payload []*models.AzureInstanceType) *GetAzureInstanceTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure instance types o k response
func (o *GetAzureInstanceTypesOK) SetPayload(payload []*models.AzureInstanceType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureInstanceTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AzureInstanceType, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAzureInstanceTypesBadRequestCode is the HTTP code returned for type GetAzureInstanceTypesBadRequest
const GetAzureInstanceTypesBadRequestCode int = 400

/*GetAzureInstanceTypesBadRequest Bad Request

swagger:response getAzureInstanceTypesBadRequest
*/
type GetAzureInstanceTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureInstanceTypesBadRequest creates GetAzureInstanceTypesBadRequest with default headers values
func NewGetAzureInstanceTypesBadRequest() *GetAzureInstanceTypesBadRequest {

	return &GetAzureInstanceTypesBadRequest{}
}

// WithPayload adds the payload to the get azure instance types bad request response
func (o *GetAzureInstanceTypesBadRequest) WithPayload(payload *models.Error) *GetAzureInstanceTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure instance types bad request response
func (o *GetAzureInstanceTypesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureInstanceTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureInstanceTypesUnauthorizedCode is the HTTP code returned for type GetAzureInstanceTypesUnauthorized
const GetAzureInstanceTypesUnauthorizedCode int = 401

/*GetAzureInstanceTypesUnauthorized Incorrect credentials

swagger:response getAzureInstanceTypesUnauthorized
*/
type GetAzureInstanceTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureInstanceTypesUnauthorized creates GetAzureInstanceTypesUnauthorized with default headers values
func NewGetAzureInstanceTypesUnauthorized() *GetAzureInstanceTypesUnauthorized {

	return &GetAzureInstanceTypesUnauthorized{}
}

// WithPayload adds the payload to the get azure instance types unauthorized response
func (o *GetAzureInstanceTypesUnauthorized) WithPayload(payload *models.Error) *GetAzureInstanceTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure instance types unauthorized response
func (o *GetAzureInstanceTypesUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureInstanceTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureInstanceTypesInternalServerErrorCode is the HTTP code returned for type GetAzureInstanceTypesInternalServerError
const GetAzureInstanceTypesInternalServerErrorCode int = 500

/*GetAzureInstanceTypesInternalServerError Internal server error

swagger:response getAzureInstanceTypesInternalServerError
*/
type GetAzureInstanceTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureInstanceTypesInternalServerError creates GetAzureInstanceTypesInternalServerError with default headers values
func NewGetAzureInstanceTypesInternalServerError() *GetAzureInstanceTypesInternalServerError {

	return &GetAzureInstanceTypesInternalServerError{}
}

// WithPayload adds the payload to the get azure instance types internal server error response
func (o *GetAzureInstanceTypesInternalServerError) WithPayload(payload *models.Error) *GetAzureInstanceTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure instance types internal server error response
func (o *GetAzureInstanceTypesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureInstanceTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
