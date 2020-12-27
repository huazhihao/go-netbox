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

// DcimDevicesUpdateReader is a Reader for the DcimDevicesUpdate structure.
type DcimDevicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesUpdateOK creates a DcimDevicesUpdateOK with default headers values
func NewDcimDevicesUpdateOK() *DcimDevicesUpdateOK {
	return &DcimDevicesUpdateOK{}
}

/*DcimDevicesUpdateOK handles this case with default header values.

DcimDevicesUpdateOK dcim devices update o k
*/
type DcimDevicesUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcimDevicesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesUpdateDefault creates a DcimDevicesUpdateDefault with default headers values
func NewDcimDevicesUpdateDefault(code int) *DcimDevicesUpdateDefault {
	return &DcimDevicesUpdateDefault{
		_statusCode: code,
	}
}

/*DcimDevicesUpdateDefault handles this case with default header values.

DcimDevicesUpdateDefault dcim devices update default
*/
type DcimDevicesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices update default response
func (o *DcimDevicesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/devices/{id}/][%d] dcim_devices_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
