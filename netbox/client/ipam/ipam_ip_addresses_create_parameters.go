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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/huazhihao/go-netbox/netbox/models"
)

// NewIpamIPAddressesCreateParams creates a new IpamIPAddressesCreateParams object
// with the default values initialized.
func NewIpamIPAddressesCreateParams() *IpamIPAddressesCreateParams {
	var ()
	return &IpamIPAddressesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesCreateParamsWithTimeout creates a new IpamIPAddressesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamIPAddressesCreateParamsWithTimeout(timeout time.Duration) *IpamIPAddressesCreateParams {
	var ()
	return &IpamIPAddressesCreateParams{

		timeout: timeout,
	}
}

// NewIpamIPAddressesCreateParamsWithContext creates a new IpamIPAddressesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamIPAddressesCreateParamsWithContext(ctx context.Context) *IpamIPAddressesCreateParams {
	var ()
	return &IpamIPAddressesCreateParams{

		Context: ctx,
	}
}

// NewIpamIPAddressesCreateParamsWithHTTPClient creates a new IpamIPAddressesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamIPAddressesCreateParamsWithHTTPClient(client *http.Client) *IpamIPAddressesCreateParams {
	var ()
	return &IpamIPAddressesCreateParams{
		HTTPClient: client,
	}
}

/*IpamIPAddressesCreateParams contains all the parameters to send to the API endpoint
for the ipam ip addresses create operation typically these are written to a http.Request
*/
type IpamIPAddressesCreateParams struct {

	/*Data*/
	Data *models.WritableIPAddress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) WithTimeout(timeout time.Duration) *IpamIPAddressesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) WithContext(ctx context.Context) *IpamIPAddressesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) WithHTTPClient(client *http.Client) *IpamIPAddressesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) WithData(data *models.WritableIPAddress) *IpamIPAddressesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip addresses create params
func (o *IpamIPAddressesCreateParams) SetData(data *models.WritableIPAddress) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
