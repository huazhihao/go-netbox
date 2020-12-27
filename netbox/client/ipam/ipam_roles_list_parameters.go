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
	"github.com/go-openapi/swag"
)

// NewIpamRolesListParams creates a new IpamRolesListParams object
// with the default values initialized.
func NewIpamRolesListParams() *IpamRolesListParams {
	var ()
	return &IpamRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesListParamsWithTimeout creates a new IpamRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamRolesListParamsWithTimeout(timeout time.Duration) *IpamRolesListParams {
	var ()
	return &IpamRolesListParams{

		timeout: timeout,
	}
}

// NewIpamRolesListParamsWithContext creates a new IpamRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamRolesListParamsWithContext(ctx context.Context) *IpamRolesListParams {
	var ()
	return &IpamRolesListParams{

		Context: ctx,
	}
}

// NewIpamRolesListParamsWithHTTPClient creates a new IpamRolesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamRolesListParamsWithHTTPClient(client *http.Client) *IpamRolesListParams {
	var ()
	return &IpamRolesListParams{
		HTTPClient: client,
	}
}

/*IpamRolesListParams contains all the parameters to send to the API endpoint
for the ipam roles list operation typically these are written to a http.Request
*/
type IpamRolesListParams struct {

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

// WithTimeout adds the timeout to the ipam roles list params
func (o *IpamRolesListParams) WithTimeout(timeout time.Duration) *IpamRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles list params
func (o *IpamRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles list params
func (o *IpamRolesListParams) WithContext(ctx context.Context) *IpamRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles list params
func (o *IpamRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles list params
func (o *IpamRolesListParams) WithHTTPClient(client *http.Client) *IpamRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles list params
func (o *IpamRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the ipam roles list params
func (o *IpamRolesListParams) WithLimit(limit *int64) *IpamRolesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam roles list params
func (o *IpamRolesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam roles list params
func (o *IpamRolesListParams) WithName(name *string) *IpamRolesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam roles list params
func (o *IpamRolesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam roles list params
func (o *IpamRolesListParams) WithOffset(offset *int64) *IpamRolesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam roles list params
func (o *IpamRolesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam roles list params
func (o *IpamRolesListParams) WithQ(q *string) *IpamRolesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam roles list params
func (o *IpamRolesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the ipam roles list params
func (o *IpamRolesListParams) WithSlug(slug *string) *IpamRolesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the ipam roles list params
func (o *IpamRolesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
