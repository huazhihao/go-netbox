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

// DeviceWithConfigContextStatus Status
//
// swagger:model deviceWithConfigContextStatus
type DeviceWithConfigContextStatus struct {

	// label
	// Required: true
	// Enum: [Offline Active Planned Staged Failed Inventory Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [offline active planned staged failed inventory decommissioning]
	Value *string `json:"value"`
}

// Validate validates this device with config context status
func (m *DeviceWithConfigContextStatus) Validate(formats strfmt.Registry) error {
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

var deviceWithConfigContextStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Active","Planned","Staged","Failed","Inventory","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextStatusTypeLabelPropEnum = append(deviceWithConfigContextStatusTypeLabelPropEnum, v)
	}
}

const (

	// DeviceWithConfigContextStatusLabelOffline captures enum value "Offline"
	DeviceWithConfigContextStatusLabelOffline string = "Offline"

	// DeviceWithConfigContextStatusLabelActive captures enum value "Active"
	DeviceWithConfigContextStatusLabelActive string = "Active"

	// DeviceWithConfigContextStatusLabelPlanned captures enum value "Planned"
	DeviceWithConfigContextStatusLabelPlanned string = "Planned"

	// DeviceWithConfigContextStatusLabelStaged captures enum value "Staged"
	DeviceWithConfigContextStatusLabelStaged string = "Staged"

	// DeviceWithConfigContextStatusLabelFailed captures enum value "Failed"
	DeviceWithConfigContextStatusLabelFailed string = "Failed"

	// DeviceWithConfigContextStatusLabelInventory captures enum value "Inventory"
	DeviceWithConfigContextStatusLabelInventory string = "Inventory"

	// DeviceWithConfigContextStatusLabelDecommissioning captures enum value "Decommissioning"
	DeviceWithConfigContextStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *DeviceWithConfigContextStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceWithConfigContextStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var deviceWithConfigContextStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","inventory","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceWithConfigContextStatusTypeValuePropEnum = append(deviceWithConfigContextStatusTypeValuePropEnum, v)
	}
}

const (

	// DeviceWithConfigContextStatusValueOffline captures enum value "offline"
	DeviceWithConfigContextStatusValueOffline string = "offline"

	// DeviceWithConfigContextStatusValueActive captures enum value "active"
	DeviceWithConfigContextStatusValueActive string = "active"

	// DeviceWithConfigContextStatusValuePlanned captures enum value "planned"
	DeviceWithConfigContextStatusValuePlanned string = "planned"

	// DeviceWithConfigContextStatusValueStaged captures enum value "staged"
	DeviceWithConfigContextStatusValueStaged string = "staged"

	// DeviceWithConfigContextStatusValueFailed captures enum value "failed"
	DeviceWithConfigContextStatusValueFailed string = "failed"

	// DeviceWithConfigContextStatusValueInventory captures enum value "inventory"
	DeviceWithConfigContextStatusValueInventory string = "inventory"

	// DeviceWithConfigContextStatusValueDecommissioning captures enum value "decommissioning"
	DeviceWithConfigContextStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *DeviceWithConfigContextStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deviceWithConfigContextStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeviceWithConfigContextStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceWithConfigContextStatus) UnmarshalBinary(b []byte) error {
	var res DeviceWithConfigContextStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}