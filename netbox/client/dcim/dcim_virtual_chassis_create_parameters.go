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

	"github.com/huazhihao/go-netbox/models"
)

// NewDcimVirtualChassisCreateParams creates a new DcimVirtualChassisCreateParams object
// with the default values initialized.
func NewDcimVirtualChassisCreateParams() *DcimVirtualChassisCreateParams {
	var ()
	return &DcimVirtualChassisCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisCreateParamsWithTimeout creates a new DcimVirtualChassisCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimVirtualChassisCreateParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisCreateParams {
	var ()
	return &DcimVirtualChassisCreateParams{

		timeout: timeout,
	}
}

// NewDcimVirtualChassisCreateParamsWithContext creates a new DcimVirtualChassisCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimVirtualChassisCreateParamsWithContext(ctx context.Context) *DcimVirtualChassisCreateParams {
	var ()
	return &DcimVirtualChassisCreateParams{

		Context: ctx,
	}
}

// NewDcimVirtualChassisCreateParamsWithHTTPClient creates a new DcimVirtualChassisCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimVirtualChassisCreateParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisCreateParams {
	var ()
	return &DcimVirtualChassisCreateParams{
		HTTPClient: client,
	}
}

/*DcimVirtualChassisCreateParams contains all the parameters to send to the API endpoint
for the dcim virtual chassis create operation typically these are written to a http.Request
*/
type DcimVirtualChassisCreateParams struct {

	/*Data*/
	Data *models.WritableVirtualChassis

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithContext(ctx context.Context) *DcimVirtualChassisCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) WithData(data *models.WritableVirtualChassis) *DcimVirtualChassisCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual chassis create params
func (o *DcimVirtualChassisCreateParams) SetData(data *models.WritableVirtualChassis) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
