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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/huazhihao/go-netbox/netbox/models"
)

// IpamVrfsUpdateReader is a Reader for the IpamVrfsUpdate structure.
type IpamVrfsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsUpdateOK creates a IpamVrfsUpdateOK with default headers values
func NewIpamVrfsUpdateOK() *IpamVrfsUpdateOK {
	return &IpamVrfsUpdateOK{}
}

/*IpamVrfsUpdateOK handles this case with default header values.

IpamVrfsUpdateOK ipam vrfs update o k
*/
type IpamVrfsUpdateOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipamVrfsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsUpdateDefault creates a IpamVrfsUpdateDefault with default headers values
func NewIpamVrfsUpdateDefault(code int) *IpamVrfsUpdateDefault {
	return &IpamVrfsUpdateDefault{
		_statusCode: code,
	}
}

/*IpamVrfsUpdateDefault handles this case with default header values.

IpamVrfsUpdateDefault ipam vrfs update default
*/
type IpamVrfsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs update default response
func (o *IpamVrfsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipam_vrfs_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
