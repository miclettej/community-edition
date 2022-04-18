// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureAccountParams azure account params
// swagger:model AzureAccountParams
type AzureAccountParams struct {

	// azure cloud
	AzureCloud string `json:"azureCloud,omitempty"`

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// subscription Id
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this azure account params
func (m *AzureAccountParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureAccountParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureAccountParams) UnmarshalBinary(b []byte) error {
	var res AzureAccountParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
