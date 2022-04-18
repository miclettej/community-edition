// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VsphereInfo vsphere info
// swagger:model vsphereInfo
type VsphereInfo struct {

	// has pacific
	HasPacific string `json:"hasPacific,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this vsphere info
func (m *VsphereInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsphereInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereInfo) UnmarshalBinary(b []byte) error {
	var res VsphereInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
