// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EtcdBackupConfigCondition etcd backup config condition
//
// swagger:model EtcdBackupConfigCondition
type EtcdBackupConfigCondition struct {

	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`

	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// last heartbeat time
	// Format: date-time
	LastHeartbeatTime Time `json:"lastHeartbeatTime,omitempty"`

	// last transition time
	// Format: date-time
	LastTransitionTime Time `json:"lastTransitionTime,omitempty"`

	// status
	Status ConditionStatus `json:"status,omitempty"`

	// type
	Type EtcdBackupConfigConditionType `json:"type,omitempty"`
}

// Validate validates this etcd backup config condition
func (m *EtcdBackupConfigCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastHeartbeatTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EtcdBackupConfigCondition) validateLastHeartbeatTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastHeartbeatTime) { // not required
		return nil
	}

	if err := m.LastHeartbeatTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastHeartbeatTime")
		}
		return err
	}

	return nil
}

func (m *EtcdBackupConfigCondition) validateLastTransitionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *EtcdBackupConfigCondition) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *EtcdBackupConfigCondition) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EtcdBackupConfigCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EtcdBackupConfigCondition) UnmarshalBinary(b []byte) error {
	var res EtcdBackupConfigCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
