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

// NewDcimSiteownersCreateParams creates a new DcimSiteownersCreateParams object
// with the default values initialized.
func NewDcimSiteownersCreateParams() *DcimSiteownersCreateParams {
	var ()
	return &DcimSiteownersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSiteownersCreateParamsWithTimeout creates a new DcimSiteownersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimSiteownersCreateParamsWithTimeout(timeout time.Duration) *DcimSiteownersCreateParams {
	var ()
	return &DcimSiteownersCreateParams{

		timeout: timeout,
	}
}

// NewDcimSiteownersCreateParamsWithContext creates a new DcimSiteownersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimSiteownersCreateParamsWithContext(ctx context.Context) *DcimSiteownersCreateParams {
	var ()
	return &DcimSiteownersCreateParams{

		Context: ctx,
	}
}

// NewDcimSiteownersCreateParamsWithHTTPClient creates a new DcimSiteownersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimSiteownersCreateParamsWithHTTPClient(client *http.Client) *DcimSiteownersCreateParams {
	var ()
	return &DcimSiteownersCreateParams{
		HTTPClient: client,
	}
}

/*DcimSiteownersCreateParams contains all the parameters to send to the API endpoint
for the dcim siteowners create operation typically these are written to a http.Request
*/
type DcimSiteownersCreateParams struct {

	/*Data*/
	Data *models.Siteowner

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) WithTimeout(timeout time.Duration) *DcimSiteownersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) WithContext(ctx context.Context) *DcimSiteownersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) WithHTTPClient(client *http.Client) *DcimSiteownersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) WithData(data *models.Siteowner) *DcimSiteownersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim siteowners create params
func (o *DcimSiteownersCreateParams) SetData(data *models.Siteowner) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSiteownersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
