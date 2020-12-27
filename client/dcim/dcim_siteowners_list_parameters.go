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

// NewDcimSiteownersListParams creates a new DcimSiteownersListParams object
// with the default values initialized.
func NewDcimSiteownersListParams() *DcimSiteownersListParams {
	var ()
	return &DcimSiteownersListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSiteownersListParamsWithTimeout creates a new DcimSiteownersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimSiteownersListParamsWithTimeout(timeout time.Duration) *DcimSiteownersListParams {
	var ()
	return &DcimSiteownersListParams{

		timeout: timeout,
	}
}

// NewDcimSiteownersListParamsWithContext creates a new DcimSiteownersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimSiteownersListParamsWithContext(ctx context.Context) *DcimSiteownersListParams {
	var ()
	return &DcimSiteownersListParams{

		Context: ctx,
	}
}

// NewDcimSiteownersListParamsWithHTTPClient creates a new DcimSiteownersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimSiteownersListParamsWithHTTPClient(client *http.Client) *DcimSiteownersListParams {
	var ()
	return &DcimSiteownersListParams{
		HTTPClient: client,
	}
}

/*DcimSiteownersListParams contains all the parameters to send to the API endpoint
for the dcim siteowners list operation typically these are written to a http.Request
*/
type DcimSiteownersListParams struct {

	/*Email*/
	Email *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithTimeout(timeout time.Duration) *DcimSiteownersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithContext(ctx context.Context) *DcimSiteownersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithHTTPClient(client *http.Client) *DcimSiteownersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithEmail(email *string) *DcimSiteownersListParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetEmail(email *string) {
	o.Email = email
}

// WithLimit adds the limit to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithLimit(limit *int64) *DcimSiteownersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithName(name *string) *DcimSiteownersListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithOffset(offset *int64) *DcimSiteownersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithQ(q *string) *DcimSiteownersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim siteowners list params
func (o *DcimSiteownersListParams) WithSlug(slug *string) *DcimSiteownersListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim siteowners list params
func (o *DcimSiteownersListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSiteownersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}