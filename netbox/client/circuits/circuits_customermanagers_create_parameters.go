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

package circuits

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

	"github.com/huazhihao/go-netbox/models"
)

// NewCircuitsCustomermanagersCreateParams creates a new CircuitsCustomermanagersCreateParams object
// with the default values initialized.
func NewCircuitsCustomermanagersCreateParams() *CircuitsCustomermanagersCreateParams {
	var ()
	return &CircuitsCustomermanagersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCustomermanagersCreateParamsWithTimeout creates a new CircuitsCustomermanagersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCustomermanagersCreateParamsWithTimeout(timeout time.Duration) *CircuitsCustomermanagersCreateParams {
	var ()
	return &CircuitsCustomermanagersCreateParams{

		timeout: timeout,
	}
}

// NewCircuitsCustomermanagersCreateParamsWithContext creates a new CircuitsCustomermanagersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCustomermanagersCreateParamsWithContext(ctx context.Context) *CircuitsCustomermanagersCreateParams {
	var ()
	return &CircuitsCustomermanagersCreateParams{

		Context: ctx,
	}
}

// NewCircuitsCustomermanagersCreateParamsWithHTTPClient creates a new CircuitsCustomermanagersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCustomermanagersCreateParamsWithHTTPClient(client *http.Client) *CircuitsCustomermanagersCreateParams {
	var ()
	return &CircuitsCustomermanagersCreateParams{
		HTTPClient: client,
	}
}

/*CircuitsCustomermanagersCreateParams contains all the parameters to send to the API endpoint
for the circuits customermanagers create operation typically these are written to a http.Request
*/
type CircuitsCustomermanagersCreateParams struct {

	/*Data*/
	Data *models.CustomerManager

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) WithTimeout(timeout time.Duration) *CircuitsCustomermanagersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) WithContext(ctx context.Context) *CircuitsCustomermanagersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) WithHTTPClient(client *http.Client) *CircuitsCustomermanagersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) WithData(data *models.CustomerManager) *CircuitsCustomermanagersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits customermanagers create params
func (o *CircuitsCustomermanagersCreateParams) SetData(data *models.CustomerManager) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCustomermanagersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
