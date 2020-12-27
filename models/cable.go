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

// Cable cable
//
// swagger:model Cable
type Cable struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Length
	// Maximum: 32767
	// Minimum: 0
	Length *int64 `json:"length,omitempty"`

	// length unit
	LengthUnit *CableLengthUnit `json:"length_unit,omitempty"`

	// status
	Status *CableStatus `json:"status,omitempty"`

	// Termination a
	// Read Only: true
	Terminationa string `json:"termination_a,omitempty"`

	// Termination a id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationaID *int64 `json:"termination_a_id"`

	// Termination a type
	// Required: true
	TerminationaType *string `json:"termination_a_type"`

	// Termination b
	// Read Only: true
	Terminationb string `json:"termination_b,omitempty"`

	// Termination b id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationbID *int64 `json:"termination_b_id"`

	// Termination b type
	// Required: true
	TerminationbType *string `json:"termination_b_type"`

	// Type
	// Enum: [1300 1500 1510 1600 1610 1700 1800 1810 3000 3010 3020 3030 3040 3500 3510 3520 3800 5000]
	Type *int64 `json:"type,omitempty"`
}

// Validate validates this cable
func (m *Cable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbType(formats); err != nil {
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

func (m *Cable) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", string(m.Color), 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", string(m.Color), `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 100); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLength(formats strfmt.Registry) error {

	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", int64(*m.Length), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", int64(*m.Length), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLengthUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	if m.LengthUnit != nil {
		if err := m.LengthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateTerminationaID(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_id", "body", m.TerminationaID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_a_id", "body", int64(*m.TerminationaID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_a_id", "body", int64(*m.TerminationaID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationaType(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_type", "body", m.TerminationaType); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationbID(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_id", "body", m.TerminationbID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_b_id", "body", int64(*m.TerminationbID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_b_id", "body", int64(*m.TerminationbID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationbType(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_type", "body", m.TerminationbType); err != nil {
		return err
	}

	return nil
}

var cableTypeTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1300,1500,1510,1600,1610,1700,1800,1810,3000,3010,3020,3030,3040,3500,3510,3520,3800,5000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableTypeTypePropEnum = append(cableTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *Cable) validateTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, cableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cable) UnmarshalBinary(b []byte) error {
	var res Cable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableLengthUnit Length unit
//
// swagger:model CableLengthUnit
type CableLengthUnit struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this cable length unit
func (m *CableLengthUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CableLengthUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *CableLengthUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableLengthUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableLengthUnit) UnmarshalBinary(b []byte) error {
	var res CableLengthUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableStatus Status
//
// swagger:model CableStatus
type CableStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this cable status
func (m *CableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CableStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *CableStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableStatus) UnmarshalBinary(b []byte) error {
	var res CableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
