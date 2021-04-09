// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MLASettings m l a settings
//
// swagger:model MLASettings
type MLASettings struct {

	// LoggingEnabled is the flag for enabling logging in user cluster.
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`

	// MonitoringEnabled is the flag for enabling monitoring in user cluster.
	MonitoringEnabled bool `json:"monitoringEnabled,omitempty"`
}

// Validate validates this m l a settings
func (m *MLASettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MLASettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MLASettings) UnmarshalBinary(b []byte) error {
	var res MLASettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
