// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RegionPrices region prices
// swagger:model regionPrices
type RegionPrices struct {

	// cpu
	CPU *Price `json:"cpu,omitempty"`

	// disk
	Disk *Price `json:"disk,omitempty"`

	// memory
	Memory *Price `json:"memory,omitempty"`
}

// Validate validates this region prices
func (m *RegionPrices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionPrices) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {

		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *RegionPrices) validateDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {

		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *RegionPrices) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {

		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegionPrices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionPrices) UnmarshalBinary(b []byte) error {
	var res RegionPrices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
