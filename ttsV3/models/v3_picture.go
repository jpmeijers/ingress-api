// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3Picture v3 picture
//
// swagger:model v3Picture
type V3Picture struct {

	// Embedded picture.
	// Omitted if there are external URLs available (in sizes).
	Embedded *PictureEmbedded `json:"embedded,omitempty"`

	// URLs of the picture for different sizes, if available on a CDN.
	Sizes map[string]string `json:"sizes,omitempty"`
}

// Validate validates this v3 picture
func (m *V3Picture) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Picture) validateEmbedded(formats strfmt.Registry) error {

	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("embedded")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Picture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Picture) UnmarshalBinary(b []byte) error {
	var res V3Picture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}