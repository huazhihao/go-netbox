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

package dcim

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
	"github.com/go-openapi/swag"
)

// NewDcimDeviceDepartmentsReadParams creates a new DcimDeviceDepartmentsReadParams object
// with the default values initialized.
func NewDcimDeviceDepartmentsReadParams() *DcimDeviceDepartmentsReadParams {
	var ()
	return &DcimDeviceDepartmentsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceDepartmentsReadParamsWithTimeout creates a new DcimDeviceDepartmentsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceDepartmentsReadParamsWithTimeout(timeout time.Duration) *DcimDeviceDepartmentsReadParams {
	var ()
	return &DcimDeviceDepartmentsReadParams{

		timeout: timeout,
	}
}

// NewDcimDeviceDepartmentsReadParamsWithContext creates a new DcimDeviceDepartmentsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceDepartmentsReadParamsWithContext(ctx context.Context) *DcimDeviceDepartmentsReadParams {
	var ()
	return &DcimDeviceDepartmentsReadParams{

		Context: ctx,
	}
}

// NewDcimDeviceDepartmentsReadParamsWithHTTPClient creates a new DcimDeviceDepartmentsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceDepartmentsReadParamsWithHTTPClient(client *http.Client) *DcimDeviceDepartmentsReadParams {
	var ()
	return &DcimDeviceDepartmentsReadParams{
		HTTPClient: client,
	}
}

/*DcimDeviceDepartmentsReadParams contains all the parameters to send to the API endpoint
for the dcim device departments read operation typically these are written to a http.Request
*/
type DcimDeviceDepartmentsReadParams struct {

	/*ID
	  A unique integer value identifying this device department.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) WithTimeout(timeout time.Duration) *DcimDeviceDepartmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) WithContext(ctx context.Context) *DcimDeviceDepartmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) WithHTTPClient(client *http.Client) *DcimDeviceDepartmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) WithID(id int64) *DcimDeviceDepartmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device departments read params
func (o *DcimDeviceDepartmentsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceDepartmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}