// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedCircuitTermination Termination a
//
// swagger:model NestedCircuitTermination
type NestedCircuitTermination struct {

	// circuit
	// Required: true
	Circuit *NestedCircuit `json:"circuit"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Termination
	// Required: true
	// Enum: [A Z]
	TermSide *string `json:"term_side"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested circuit termination
func (m *NestedCircuitTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedCircuitTermination) validateCircuit(formats strfmt.Registry) error {

	if err := validate.Required("circuit", "body", m.Circuit); err != nil {
		return err
	}

	if m.Circuit != nil {
		if err := m.Circuit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("circuit")
			}
			return err
		}
	}

	return nil
}

var nestedCircuitTerminationTypeTermSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","Z"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedCircuitTerminationTypeTermSidePropEnum = append(nestedCircuitTerminationTypeTermSidePropEnum, v)
	}
}

const (

	// NestedCircuitTerminationTermSideA captures enum value "A"
	NestedCircuitTerminationTermSideA string = "A"

	// NestedCircuitTerminationTermSideZ captures enum value "Z"
	NestedCircuitTerminationTermSideZ string = "Z"
)

// prop value enum
func (m *NestedCircuitTermination) validateTermSideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nestedCircuitTerminationTypeTermSidePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedCircuitTermination) validateTermSide(formats strfmt.Registry) error {

	if err := validate.Required("term_side", "body", m.TermSide); err != nil {
		return err
	}

	// value enum
	if err := m.validateTermSideEnum("term_side", "body", *m.TermSide); err != nil {
		return err
	}

	return nil
}

func (m *NestedCircuitTermination) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedCircuitTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedCircuitTermination) UnmarshalBinary(b []byte) error {
	var res NestedCircuitTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
