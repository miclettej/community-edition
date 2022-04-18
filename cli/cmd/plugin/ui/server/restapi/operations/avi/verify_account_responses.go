// Code generated by go-swagger; DO NOT EDIT.

package avi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// VerifyAccountCreatedCode is the HTTP code returned for type VerifyAccountCreated
const VerifyAccountCreatedCode int = 201

/*VerifyAccountCreated Verified AVI credentials successfully

swagger:response verifyAccountCreated
*/
type VerifyAccountCreated struct {
}

// NewVerifyAccountCreated creates VerifyAccountCreated with default headers values
func NewVerifyAccountCreated() *VerifyAccountCreated {

	return &VerifyAccountCreated{}
}

// WriteResponse to the client
func (o *VerifyAccountCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// VerifyAccountBadRequestCode is the HTTP code returned for type VerifyAccountBadRequest
const VerifyAccountBadRequestCode int = 400

/*VerifyAccountBadRequest Bad request

swagger:response verifyAccountBadRequest
*/
type VerifyAccountBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyAccountBadRequest creates VerifyAccountBadRequest with default headers values
func NewVerifyAccountBadRequest() *VerifyAccountBadRequest {

	return &VerifyAccountBadRequest{}
}

// WithPayload adds the payload to the verify account bad request response
func (o *VerifyAccountBadRequest) WithPayload(payload *models.Error) *VerifyAccountBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify account bad request response
func (o *VerifyAccountBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyAccountUnauthorizedCode is the HTTP code returned for type VerifyAccountUnauthorized
const VerifyAccountUnauthorizedCode int = 401

/*VerifyAccountUnauthorized Incorrect credentials

swagger:response verifyAccountUnauthorized
*/
type VerifyAccountUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyAccountUnauthorized creates VerifyAccountUnauthorized with default headers values
func NewVerifyAccountUnauthorized() *VerifyAccountUnauthorized {

	return &VerifyAccountUnauthorized{}
}

// WithPayload adds the payload to the verify account unauthorized response
func (o *VerifyAccountUnauthorized) WithPayload(payload *models.Error) *VerifyAccountUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify account unauthorized response
func (o *VerifyAccountUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyAccountInternalServerErrorCode is the HTTP code returned for type VerifyAccountInternalServerError
const VerifyAccountInternalServerErrorCode int = 500

/*VerifyAccountInternalServerError Internal server error

swagger:response verifyAccountInternalServerError
*/
type VerifyAccountInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyAccountInternalServerError creates VerifyAccountInternalServerError with default headers values
func NewVerifyAccountInternalServerError() *VerifyAccountInternalServerError {

	return &VerifyAccountInternalServerError{}
}

// WithPayload adds the payload to the verify account internal server error response
func (o *VerifyAccountInternalServerError) WithPayload(payload *models.Error) *VerifyAccountInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify account internal server error response
func (o *VerifyAccountInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyAccountInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
