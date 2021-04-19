// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConcentratorConfigLoRaStandardChannel concentrator config lo ra standard channel
//
// swagger:model ConcentratorConfigLoRaStandardChannel
type ConcentratorConfigLoRaStandardChannel struct {

	// Bandwidth (Hz).
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// Frequency (Hz).
	Frequency string `json:"frequency,omitempty"`

	// radio
	Radio int64 `json:"radio,omitempty"`

	// spreading factor
	SpreadingFactor int64 `json:"spreading_factor,omitempty"`
}

// Validate validates this concentrator config lo ra standard channel
func (m *ConcentratorConfigLoRaStandardChannel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConcentratorConfigLoRaStandardChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConcentratorConfigLoRaStandardChannel) UnmarshalBinary(b []byte) error {
	var res ConcentratorConfigLoRaStandardChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}