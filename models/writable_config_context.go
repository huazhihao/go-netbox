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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConfigContext writable config context
//
// swagger:model WritableConfigContext
type WritableConfigContext struct {

	// Data
	// Required: true
	Data *string `json:"data"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is active
	IsActive bool `json:"is_active,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// platforms
	// Unique: true
	Platforms []int64 `json:"platforms"`

	// regions
	// Unique: true
	Regions []int64 `json:"regions"`

	// roles
	// Unique: true
	Roles []int64 `json:"roles"`

	// sites
	// Unique: true
	Sites []int64 `json:"sites"`

	// tenant groups
	// Unique: true
	TenantGroups []int64 `json:"tenant_groups"`

	// tenants
	// Unique: true
	Tenants []int64 `json:"tenants"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this writable config context
func (m *WritableConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConfigContext) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validatePlatforms(formats strfmt.Registry) error {

	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	if err := validate.UniqueItems("platforms", "body", m.Platforms); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateSites(formats strfmt.Registry) error {

	if swag.IsZero(m.Sites) { // not required
		return nil
	}

	if err := validate.UniqueItems("sites", "body", m.Sites); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateTenantGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.TenantGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenant_groups", "body", m.TenantGroups); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateTenants(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenants", "body", m.Tenants); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", int64(*m.Weight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", int64(*m.Weight), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
