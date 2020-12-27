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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/huazhihao/go-netbox/netbox/models"
)

// ExtrasTagsCreateReader is a Reader for the ExtrasTagsCreate structure.
type ExtrasTagsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasTagsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasTagsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsCreateCreated creates a ExtrasTagsCreateCreated with default headers values
func NewExtrasTagsCreateCreated() *ExtrasTagsCreateCreated {
	return &ExtrasTagsCreateCreated{}
}

/*ExtrasTagsCreateCreated handles this case with default header values.

ExtrasTagsCreateCreated extras tags create created
*/
type ExtrasTagsCreateCreated struct {
	Payload *models.Tag
}

func (o *ExtrasTagsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/tags/][%d] extrasTagsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasTagsCreateCreated) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsCreateDefault creates a ExtrasTagsCreateDefault with default headers values
func NewExtrasTagsCreateDefault(code int) *ExtrasTagsCreateDefault {
	return &ExtrasTagsCreateDefault{
		_statusCode: code,
	}
}

/*ExtrasTagsCreateDefault handles this case with default header values.

ExtrasTagsCreateDefault extras tags create default
*/
type ExtrasTagsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags create default response
func (o *ExtrasTagsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasTagsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/tags/][%d] extras_tags_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
