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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCircuit writable circuit
//
// swagger:model WritableCircuit
type WritableCircuit struct {

	// Action
	Action string `json:"action,omitempty"`

	// Bandwidth history
	BandwidthHistory string `json:"bandwidth_history,omitempty"`

	// Circuit ID
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Cid *string `json:"cid"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Commit rate (Mbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	CommitRate *int64 `json:"commit_rate,omitempty"`

	// Contract end date
	// Format: date
	ContractEndDate *strfmt.Date `json:"contract_end_date,omitempty"`

	// Contract start date
	// Format: date
	ContractStartDate *strfmt.Date `json:"contract_start_date,omitempty"`

	// Cost center
	// Max Length: 100
	CostCenter *string `json:"cost_center,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Customermanager
	Customermanager *int64 `json:"customermanager,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Date installed
	// Format: date
	InstallDate *strfmt.Date `json:"install_date,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Librenms hostid
	// Maximum: 2.147483647e+09
	// Minimum: -2.147483648e+09
	LibrenmsHostid *int64 `json:"librenms_hostid,omitempty"`

	// Librenms itemid
	// Maximum: 2.147483647e+09
	// Minimum: -2.147483648e+09
	LibrenmsItemid *int64 `json:"librenms_itemid,omitempty"`

	// Line name
	// Max Length: 50
	LineName *string `json:"line_name,omitempty"`

	// Netflow hostid
	// Maximum: 2.147483647e+09
	// Minimum: -2.147483648e+09
	NetflowHostid *int64 `json:"netflow_hostid,omitempty"`

	// Netflow itemid
	// Maximum: 2.147483647e+09
	// Minimum: -2.147483648e+09
	NetflowItemid *int64 `json:"netflow_itemid,omitempty"`

	// Payment circle
	// Enum: [1 2 3 4 0]
	PaymentCircle int64 `json:"payment_circle,omitempty"`

	// Provider
	// Required: true
	Provider *int64 `json:"provider"`

	// Qos1 commit rate (Mbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Qos1Bandwidth *int64 `json:"qos1_bandwidth,omitempty"`

	// Qos2 commit rate (Mbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Qos2Bandwidth *int64 `json:"qos2_bandwidth,omitempty"`

	// Qos3 commit rate (Mbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Qos3Bandwidth *int64 `json:"qos3_bandwidth,omitempty"`

	// Redundancy type
	// Enum: [0 1 2 3]
	RedundancyType int64 `json:"redundancy_type,omitempty"`

	// Status
	// Enum: [2 3 1 4 0 5]
	Status int64 `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Termination a
	// Read Only: true
	Terminationa string `json:"termination_a,omitempty"`

	// Termination z
	// Read Only: true
	Terminationz string `json:"termination_z,omitempty"`

	// Type
	// Required: true
	Type *int64 `json:"type"`

	// Vendor availabe ips
	// Max Length: 100
	VendorAvailabeIps *string `json:"vendor_availabe_ips,omitempty"`

	// Vendor gw
	// Max Length: 50
	VendorGw *string `json:"vendor_gw,omitempty"`

	// Vendor interface type
	// Enum: [0 1 2]
	VendorInterfaceType int64 `json:"vendor_interface_type,omitempty"`
}

// Validate validates this writable circuit
func (m *WritableCircuit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostCenter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibrenmsHostid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibrenmsItemid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetflowHostid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetflowItemid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentCircle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos1Bandwidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos2Bandwidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos3Bandwidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedundancyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorAvailabeIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorGw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorInterfaceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCircuit) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	if err := validate.MinLength("cid", "body", string(*m.Cid), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("cid", "body", string(*m.Cid), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateCommitRate(formats strfmt.Registry) error {

	if swag.IsZero(m.CommitRate) { // not required
		return nil
	}

	if err := validate.MinimumInt("commit_rate", "body", int64(*m.CommitRate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("commit_rate", "body", int64(*m.CommitRate), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateContractEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ContractEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("contract_end_date", "body", "date", m.ContractEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateContractStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ContractStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("contract_start_date", "body", "date", m.ContractStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateCostCenter(formats strfmt.Registry) error {

	if swag.IsZero(m.CostCenter) { // not required
		return nil
	}

	if err := validate.MaxLength("cost_center", "body", string(*m.CostCenter), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateInstallDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallDate) { // not required
		return nil
	}

	if err := validate.FormatOf("install_date", "body", "date", m.InstallDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateLibrenmsHostid(formats strfmt.Registry) error {

	if swag.IsZero(m.LibrenmsHostid) { // not required
		return nil
	}

	if err := validate.MinimumInt("librenms_hostid", "body", int64(*m.LibrenmsHostid), -2.147483648e+09, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("librenms_hostid", "body", int64(*m.LibrenmsHostid), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateLibrenmsItemid(formats strfmt.Registry) error {

	if swag.IsZero(m.LibrenmsItemid) { // not required
		return nil
	}

	if err := validate.MinimumInt("librenms_itemid", "body", int64(*m.LibrenmsItemid), -2.147483648e+09, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("librenms_itemid", "body", int64(*m.LibrenmsItemid), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateLineName(formats strfmt.Registry) error {

	if swag.IsZero(m.LineName) { // not required
		return nil
	}

	if err := validate.MaxLength("line_name", "body", string(*m.LineName), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateNetflowHostid(formats strfmt.Registry) error {

	if swag.IsZero(m.NetflowHostid) { // not required
		return nil
	}

	if err := validate.MinimumInt("netflow_hostid", "body", int64(*m.NetflowHostid), -2.147483648e+09, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("netflow_hostid", "body", int64(*m.NetflowHostid), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateNetflowItemid(formats strfmt.Registry) error {

	if swag.IsZero(m.NetflowItemid) { // not required
		return nil
	}

	if err := validate.MinimumInt("netflow_itemid", "body", int64(*m.NetflowItemid), -2.147483648e+09, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("netflow_itemid", "body", int64(*m.NetflowItemid), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

var writableCircuitTypePaymentCirclePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4,0]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTypePaymentCirclePropEnum = append(writableCircuitTypePaymentCirclePropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuit) validatePaymentCircleEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableCircuitTypePaymentCirclePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuit) validatePaymentCircle(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentCircle) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentCircleEnum("payment_circle", "body", m.PaymentCircle); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateQos1Bandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos1Bandwidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("qos1_bandwidth", "body", int64(*m.Qos1Bandwidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("qos1_bandwidth", "body", int64(*m.Qos1Bandwidth), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateQos2Bandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos2Bandwidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("qos2_bandwidth", "body", int64(*m.Qos2Bandwidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("qos2_bandwidth", "body", int64(*m.Qos2Bandwidth), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateQos3Bandwidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos3Bandwidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("qos3_bandwidth", "body", int64(*m.Qos3Bandwidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("qos3_bandwidth", "body", int64(*m.Qos3Bandwidth), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

var writableCircuitTypeRedundancyTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTypeRedundancyTypePropEnum = append(writableCircuitTypeRedundancyTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuit) validateRedundancyTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableCircuitTypeRedundancyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuit) validateRedundancyType(formats strfmt.Registry) error {

	if swag.IsZero(m.RedundancyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedundancyTypeEnum("redundancy_type", "body", m.RedundancyType); err != nil {
		return err
	}

	return nil
}

var writableCircuitTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[2,3,1,4,0,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTypeStatusPropEnum = append(writableCircuitTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuit) validateStatusEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableCircuitTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuit) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *WritableCircuit) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateVendorAvailabeIps(formats strfmt.Registry) error {

	if swag.IsZero(m.VendorAvailabeIps) { // not required
		return nil
	}

	if err := validate.MaxLength("vendor_availabe_ips", "body", string(*m.VendorAvailabeIps), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateVendorGw(formats strfmt.Registry) error {

	if swag.IsZero(m.VendorGw) { // not required
		return nil
	}

	if err := validate.MaxLength("vendor_gw", "body", string(*m.VendorGw), 50); err != nil {
		return err
	}

	return nil
}

var writableCircuitTypeVendorInterfaceTypePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTypeVendorInterfaceTypePropEnum = append(writableCircuitTypeVendorInterfaceTypePropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuit) validateVendorInterfaceTypeEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, writableCircuitTypeVendorInterfaceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuit) validateVendorInterfaceType(formats strfmt.Registry) error {

	if swag.IsZero(m.VendorInterfaceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVendorInterfaceTypeEnum("vendor_interface_type", "body", m.VendorInterfaceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCircuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCircuit) UnmarshalBinary(b []byte) error {
	var res WritableCircuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}