// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3MACSettings v3 m a c settings
//
// swagger:model v3MACSettings
type V3MACSettings struct {

	// The ADR margin tells the network server how much margin it should add in ADR requests.
	// A bigger margin is less efficient, but gives a better chance of successful reception.
	// If unset, the default value from Network Server configuration will be used.
	AdrMargin float32 `json:"adr_margin,omitempty"`

	// Frequency of the class B beacon (Hz).
	// If unset, the default value from Network Server configuration will be used.
	BeaconFrequency string `json:"beacon_frequency,omitempty"`

	// Maximum delay for the device to answer a MAC request or a confirmed downlink frame.
	// If unset, the default value from Network Server configuration will be used.
	ClassbTimeout string `json:"class_b_timeout,omitempty"`

	// Maximum delay for the device to answer a MAC request or a confirmed downlink frame.
	// If unset, the default value from Network Server configuration will be used.
	ClasscTimeout string `json:"class_c_timeout,omitempty"`

	// The ADR ACK delay Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	DesiredAdrAckDelayExponent *V3ADRAckDelayExponentValue `json:"desired_adr_ack_delay_exponent,omitempty"`

	// The ADR ACK limit Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	DesiredAdrAckLimitExponent *V3ADRAckLimitExponentValue `json:"desired_adr_ack_limit_exponent,omitempty"`

	// The frequency of the class B beacon (Hz) Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration will be used.
	DesiredBeaconFrequency string `json:"desired_beacon_frequency,omitempty"`

	// The maximum uplink duty cycle (of all channels) Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration will be used.
	DesiredMaxDutyCycle *V3AggregatedDutyCycleValue `json:"desired_max_duty_cycle,omitempty"`

	// The data rate index of the class B ping slot Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration will be used.
	DesiredPingSlotDataRateIndex *V3DataRateIndexValue `json:"desired_ping_slot_data_rate_index,omitempty"`

	// The frequency of the class B ping slot (Hz) Network Server should configure device to use via MAC commands.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	DesiredPingSlotFrequency string `json:"desired_ping_slot_frequency,omitempty"`

	// The Rx1 data rate offset Network Server should configure device to use via MAC commands or Join-Accept.
	// If unset, the default value from Network Server configuration will be used.
	DesiredRx1DataRateOffset int64 `json:"desired_rx1_data_rate_offset,omitempty"`

	// The Rx1 delay Network Server should configure device to use via MAC commands or Join-Accept.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	DesiredRx1Delay *V3RxDelayValue `json:"desired_rx1_delay,omitempty"`

	// The Rx2 data rate index Network Server should configure device to use via MAC commands or Join-Accept.
	// If unset, the default value from frequency plan, Network Server configuration or regional parameters specification will be used.
	DesiredRx2DataRateIndex *V3DataRateIndexValue `json:"desired_rx2_data_rate_index,omitempty"`

	// The Rx2 frequency index Network Server should configure device to use via MAC commands.
	// If unset, the default value from frequency plan, Network Server configuration or regional parameters specification will be used.
	DesiredRx2Frequency string `json:"desired_rx2_frequency,omitempty"`

	// List of factory-preset frequencies.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	FactoryPresetFrequencies []string `json:"factory_preset_frequencies"`

	// Maximum uplink duty cycle (of all channels).
	MaxDutyCycle *V3AggregatedDutyCycleValue `json:"max_duty_cycle,omitempty"`

	// Data rate index of the class B ping slot.
	// If unset, the default value from Network Server configuration will be used.
	PingSlotDataRateIndex *V3DataRateIndexValue `json:"ping_slot_data_rate_index,omitempty"`

	// Frequency of the class B ping slot (Hz).
	// If unset, the default value from Network Server configuration will be used.
	PingSlotFrequency string `json:"ping_slot_frequency,omitempty"`

	// Periodicity of the class B ping slot.
	// If unset, the default value from Network Server configuration will be used.
	PingSlotPeriodicity *V3PingSlotPeriodValue `json:"ping_slot_periodicity,omitempty"`

	// Whether the device resets the frame counters (not LoRaWAN compliant).
	// If unset, the default value from Network Server configuration will be used.
	ResetsfCnt bool `json:"resets_f_cnt,omitempty"`

	// Rx1 data rate offset.
	// If unset, the default value from Network Server configuration will be used.
	Rx1DataRateOffset int64 `json:"rx1_data_rate_offset,omitempty"`

	// Class A Rx1 delay.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	Rx1Delay *V3RxDelayValue `json:"rx1_delay,omitempty"`

	// Data rate index for Rx2.
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	Rx2DataRateIndex *V3DataRateIndexValue `json:"rx2_data_rate_index,omitempty"`

	// Frequency for Rx2 (Hz).
	// If unset, the default value from Network Server configuration or regional parameters specification will be used.
	Rx2Frequency string `json:"rx2_frequency,omitempty"`

	// Number of uplink messages after which a DevStatusReq MACCommand shall be sent.
	// If unset, the default value from Network Server configuration will be used.
	StatusCountPeriodicity int64 `json:"status_count_periodicity,omitempty"`

	// The interval after which a DevStatusReq MACCommand shall be sent.
	// If unset, the default value from Network Server configuration will be used.
	StatusTimePeriodicity string `json:"status_time_periodicity,omitempty"`

	// Whether the device supports 32-bit frame counters.
	// If unset, the default value from Network Server configuration will be used.
	Supports32BitfCnt bool `json:"supports_32_bit_f_cnt,omitempty"`

	// Whether the Network Server should use ADR for the device.
	// If unset, the default value from Network Server configuration will be used.
	UseAdr bool `json:"use_adr,omitempty"`
}

// Validate validates this v3 m a c settings
func (m *V3MACSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredAdrAckDelayExponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredAdrAckLimitExponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredMaxDutyCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredPingSlotDataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredRx1Delay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredRx2DataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxDutyCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePingSlotDataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePingSlotPeriodicity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRx1Delay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRx2DataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3MACSettings) validateDesiredAdrAckDelayExponent(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredAdrAckDelayExponent) { // not required
		return nil
	}

	if m.DesiredAdrAckDelayExponent != nil {
		if err := m.DesiredAdrAckDelayExponent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_adr_ack_delay_exponent")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateDesiredAdrAckLimitExponent(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredAdrAckLimitExponent) { // not required
		return nil
	}

	if m.DesiredAdrAckLimitExponent != nil {
		if err := m.DesiredAdrAckLimitExponent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_adr_ack_limit_exponent")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateDesiredMaxDutyCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredMaxDutyCycle) { // not required
		return nil
	}

	if m.DesiredMaxDutyCycle != nil {
		if err := m.DesiredMaxDutyCycle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_max_duty_cycle")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateDesiredPingSlotDataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredPingSlotDataRateIndex) { // not required
		return nil
	}

	if m.DesiredPingSlotDataRateIndex != nil {
		if err := m.DesiredPingSlotDataRateIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_ping_slot_data_rate_index")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateDesiredRx1Delay(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredRx1Delay) { // not required
		return nil
	}

	if m.DesiredRx1Delay != nil {
		if err := m.DesiredRx1Delay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_rx1_delay")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateDesiredRx2DataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.DesiredRx2DataRateIndex) { // not required
		return nil
	}

	if m.DesiredRx2DataRateIndex != nil {
		if err := m.DesiredRx2DataRateIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desired_rx2_data_rate_index")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateMaxDutyCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxDutyCycle) { // not required
		return nil
	}

	if m.MaxDutyCycle != nil {
		if err := m.MaxDutyCycle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_duty_cycle")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validatePingSlotDataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.PingSlotDataRateIndex) { // not required
		return nil
	}

	if m.PingSlotDataRateIndex != nil {
		if err := m.PingSlotDataRateIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ping_slot_data_rate_index")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validatePingSlotPeriodicity(formats strfmt.Registry) error {

	if swag.IsZero(m.PingSlotPeriodicity) { // not required
		return nil
	}

	if m.PingSlotPeriodicity != nil {
		if err := m.PingSlotPeriodicity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ping_slot_periodicity")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateRx1Delay(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx1Delay) { // not required
		return nil
	}

	if m.Rx1Delay != nil {
		if err := m.Rx1Delay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rx1_delay")
			}
			return err
		}
	}

	return nil
}

func (m *V3MACSettings) validateRx2DataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx2DataRateIndex) { // not required
		return nil
	}

	if m.Rx2DataRateIndex != nil {
		if err := m.Rx2DataRateIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rx2_data_rate_index")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3MACSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3MACSettings) UnmarshalBinary(b []byte) error {
	var res V3MACSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}