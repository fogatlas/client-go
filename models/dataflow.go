// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Dataflow dataflow
// swagger:model Dataflow
type Dataflow struct {

	// Required amount of bandwidth available on the link - bit/s
	BandwidthRequired int64 `json:"bandwidth_required,omitempty"`

	// destination id
	DestinationID string `json:"destination_id,omitempty"`

	// Required link latency - milliseconds
	LatencyRequired int64 `json:"latency_required,omitempty"`

	// source id
	SourceID string `json:"source_id,omitempty"`
}

// Validate validates this dataflow
func (m *Dataflow) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Dataflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dataflow) UnmarshalBinary(b []byte) error {
	var res Dataflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
