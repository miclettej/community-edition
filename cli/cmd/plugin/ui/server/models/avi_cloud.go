// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AviCloud avi cloud
// swagger:model AviCloud
type AviCloud struct {

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this avi cloud
func (m *AviCloud) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AviCloud) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AviCloud) UnmarshalBinary(b []byte) error {
	var res AviCloud
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
