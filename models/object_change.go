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

// ObjectChange object change
//
// swagger:model ObjectChange
type ObjectChange struct {

	// Action
	// Required: true
	// Enum: [1 2 3]
	Action *int64 `json:"action"`

	// Changed object
	// Read Only: true
	ChangedObject string `json:"changed_object,omitempty"`

	// Content type
	// Read Only: true
	ContentType string `json:"content_type,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Object data
	// Read Only: true
	ObjectData string `json:"object_data,omitempty"`

	// Request id
	// Read Only: true
	// Format: uuid
	RequestID strfmt.UUID `json:"request_id,omitempty"`

	// Time
	// Read Only: true
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// user
	User *NestedUser `json:"user,omitempty"`

	// User name
	// Read Only: true
	// Min Length: 1
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this object change
func (m *ObjectChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectChangeTypeActionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectChangeTypeActionPropEnum = append(objectChangeTypeActionPropEnum, v)
	}
}

// prop value enum
func (m *ObjectChange) validateActionEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, objectChangeTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectChange) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("request_id", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectChange) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectChange) validateUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.UserName) { // not required
		return nil
	}

	if err := validate.MinLength("user_name", "body", string(m.UserName), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectChange) UnmarshalBinary(b []byte) error {
	var res ObjectChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}