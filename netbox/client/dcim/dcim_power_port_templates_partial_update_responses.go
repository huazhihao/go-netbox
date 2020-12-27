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

// DcimPowerPortTemplatesPartialUpdateReader is a Reader for the DcimPowerPortTemplatesPartialUpdate structure.
type DcimPowerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesPartialUpdateOK creates a DcimPowerPortTemplatesPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesPartialUpdateOK() *DcimPowerPortTemplatesPartialUpdateOK {
	return &DcimPowerPortTemplatesPartialUpdateOK{}
}

/*DcimPowerPortTemplatesPartialUpdateOK handles this case with default header values.

DcimPowerPortTemplatesPartialUpdateOK dcim power port templates partial update o k
*/
type DcimPowerPortTemplatesPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesPartialUpdateDefault creates a DcimPowerPortTemplatesPartialUpdateDefault with default headers values
func NewDcimPowerPortTemplatesPartialUpdateDefault(code int) *DcimPowerPortTemplatesPartialUpdateDefault {
	return &DcimPowerPortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimPowerPortTemplatesPartialUpdateDefault handles this case with default header values.

DcimPowerPortTemplatesPartialUpdateDefault dcim power port templates partial update default
*/
type DcimPowerPortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power port templates partial update default response
func (o *DcimPowerPortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
