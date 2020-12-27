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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/huazhihao/go-netbox/netbox/models"
)

// DcimSitesCreateReader is a Reader for the DcimSitesCreate structure.
type DcimSitesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimSitesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesCreateCreated creates a DcimSitesCreateCreated with default headers values
func NewDcimSitesCreateCreated() *DcimSitesCreateCreated {
	return &DcimSitesCreateCreated{}
}

/*DcimSitesCreateCreated handles this case with default header values.

DcimSitesCreateCreated dcim sites create created
*/
type DcimSitesCreateCreated struct {
	Payload *models.Site
}

func (o *DcimSitesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcimSitesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimSitesCreateCreated) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesCreateDefault creates a DcimSitesCreateDefault with default headers values
func NewDcimSitesCreateDefault(code int) *DcimSitesCreateDefault {
	return &DcimSitesCreateDefault{
		_statusCode: code,
	}
}

/*DcimSitesCreateDefault handles this case with default header values.

DcimSitesCreateDefault dcim sites create default
*/
type DcimSitesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites create default response
func (o *DcimSitesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/sites/][%d] dcim_sites_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
