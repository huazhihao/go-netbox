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

// NewDcimPriorityReadParams creates a new DcimPriorityReadParams object
// with the default values initialized.
func NewDcimPriorityReadParams() *DcimPriorityReadParams {
	var ()
	return &DcimPriorityReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPriorityReadParamsWithTimeout creates a new DcimPriorityReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPriorityReadParamsWithTimeout(timeout time.Duration) *DcimPriorityReadParams {
	var ()
	return &DcimPriorityReadParams{

		timeout: timeout,
	}
}

// NewDcimPriorityReadParamsWithContext creates a new DcimPriorityReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPriorityReadParamsWithContext(ctx context.Context) *DcimPriorityReadParams {
	var ()
	return &DcimPriorityReadParams{

		Context: ctx,
	}
}

// NewDcimPriorityReadParamsWithHTTPClient creates a new DcimPriorityReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPriorityReadParamsWithHTTPClient(client *http.Client) *DcimPriorityReadParams {
	var ()
	return &DcimPriorityReadParams{
		HTTPClient: client,
	}
}

/*DcimPriorityReadParams contains all the parameters to send to the API endpoint
for the dcim priority read operation typically these are written to a http.Request
*/
type DcimPriorityReadParams struct {

	/*ID
	  A unique integer value identifying this priority.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim priority read params
func (o *DcimPriorityReadParams) WithTimeout(timeout time.Duration) *DcimPriorityReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim priority read params
func (o *DcimPriorityReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim priority read params
func (o *DcimPriorityReadParams) WithContext(ctx context.Context) *DcimPriorityReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim priority read params
func (o *DcimPriorityReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim priority read params
func (o *DcimPriorityReadParams) WithHTTPClient(client *http.Client) *DcimPriorityReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim priority read params
func (o *DcimPriorityReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim priority read params
func (o *DcimPriorityReadParams) WithID(id int64) *DcimPriorityReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim priority read params
func (o *DcimPriorityReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPriorityReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
