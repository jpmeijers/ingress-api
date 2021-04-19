// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3GetCollaboratorResponse v3 get collaborator response
//
// swagger:model v3GetCollaboratorResponse
type V3GetCollaboratorResponse struct {

	// ids
	Ids *V3OrganizationOrUserIdentifiers `json:"ids,omitempty"`

	// rights
	Rights []V3Right `json:"rights"`
}

// Validate validates this v3 get collaborator response
func (m *V3GetCollaboratorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3GetCollaboratorResponse) validateIds(formats strfmt.Registry) error {

	if swag.IsZero(m.Ids) { // not required
		return nil
	}

	if m.Ids != nil {
		if err := m.Ids.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3GetCollaboratorResponse) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {

		if err := m.Rights[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3GetCollaboratorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3GetCollaboratorResponse) UnmarshalBinary(b []byte) error {
	var res V3GetCollaboratorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}