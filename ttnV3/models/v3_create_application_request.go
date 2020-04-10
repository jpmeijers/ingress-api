// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3CreateApplicationRequest v3 create application request
//
// swagger:model v3CreateApplicationRequest
type V3CreateApplicationRequest struct {

	// application
	Application *V3Application `json:"application,omitempty"`

	// Collaborator to grant all rights on the newly created application.
	Collaborator *V3OrganizationOrUserIdentifiers `json:"collaborator,omitempty"`
}

// Validate validates this v3 create application request
func (m *V3CreateApplicationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollaborator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3CreateApplicationRequest) validateApplication(formats strfmt.Registry) error {

	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *V3CreateApplicationRequest) validateCollaborator(formats strfmt.Registry) error {

	if swag.IsZero(m.Collaborator) { // not required
		return nil
	}

	if m.Collaborator != nil {
		if err := m.Collaborator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collaborator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3CreateApplicationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3CreateApplicationRequest) UnmarshalBinary(b []byte) error {
	var res V3CreateApplicationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}