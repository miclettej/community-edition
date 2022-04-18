// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// VerifyLdapCloseConnectionCreatedCode is the HTTP code returned for type VerifyLdapCloseConnectionCreated
const VerifyLdapCloseConnectionCreatedCode int = 201

/*VerifyLdapCloseConnectionCreated Verified LDAP credentials successfully

swagger:response verifyLdapCloseConnectionCreated
*/
type VerifyLdapCloseConnectionCreated struct {
}

// NewVerifyLdapCloseConnectionCreated creates VerifyLdapCloseConnectionCreated with default headers values
func NewVerifyLdapCloseConnectionCreated() *VerifyLdapCloseConnectionCreated {

	return &VerifyLdapCloseConnectionCreated{}
}

// WriteResponse to the client
func (o *VerifyLdapCloseConnectionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// VerifyLdapCloseConnectionBadRequestCode is the HTTP code returned for type VerifyLdapCloseConnectionBadRequest
const VerifyLdapCloseConnectionBadRequestCode int = 400

/*VerifyLdapCloseConnectionBadRequest Bad request

swagger:response verifyLdapCloseConnectionBadRequest
*/
type VerifyLdapCloseConnectionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyLdapCloseConnectionBadRequest creates VerifyLdapCloseConnectionBadRequest with default headers values
func NewVerifyLdapCloseConnectionBadRequest() *VerifyLdapCloseConnectionBadRequest {

	return &VerifyLdapCloseConnectionBadRequest{}
}

// WithPayload adds the payload to the verify ldap close connection bad request response
func (o *VerifyLdapCloseConnectionBadRequest) WithPayload(payload *models.Error) *VerifyLdapCloseConnectionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify ldap close connection bad request response
func (o *VerifyLdapCloseConnectionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyLdapCloseConnectionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyLdapCloseConnectionUnauthorizedCode is the HTTP code returned for type VerifyLdapCloseConnectionUnauthorized
const VerifyLdapCloseConnectionUnauthorizedCode int = 401

/*VerifyLdapCloseConnectionUnauthorized Incorrect credentials

swagger:response verifyLdapCloseConnectionUnauthorized
*/
type VerifyLdapCloseConnectionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyLdapCloseConnectionUnauthorized creates VerifyLdapCloseConnectionUnauthorized with default headers values
func NewVerifyLdapCloseConnectionUnauthorized() *VerifyLdapCloseConnectionUnauthorized {

	return &VerifyLdapCloseConnectionUnauthorized{}
}

// WithPayload adds the payload to the verify ldap close connection unauthorized response
func (o *VerifyLdapCloseConnectionUnauthorized) WithPayload(payload *models.Error) *VerifyLdapCloseConnectionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify ldap close connection unauthorized response
func (o *VerifyLdapCloseConnectionUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyLdapCloseConnectionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyLdapCloseConnectionInternalServerErrorCode is the HTTP code returned for type VerifyLdapCloseConnectionInternalServerError
const VerifyLdapCloseConnectionInternalServerErrorCode int = 500

/*VerifyLdapCloseConnectionInternalServerError Internal server error

swagger:response verifyLdapCloseConnectionInternalServerError
*/
type VerifyLdapCloseConnectionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerifyLdapCloseConnectionInternalServerError creates VerifyLdapCloseConnectionInternalServerError with default headers values
func NewVerifyLdapCloseConnectionInternalServerError() *VerifyLdapCloseConnectionInternalServerError {

	return &VerifyLdapCloseConnectionInternalServerError{}
}

// WithPayload adds the payload to the verify ldap close connection internal server error response
func (o *VerifyLdapCloseConnectionInternalServerError) WithPayload(payload *models.Error) *VerifyLdapCloseConnectionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify ldap close connection internal server error response
func (o *VerifyLdapCloseConnectionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyLdapCloseConnectionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
