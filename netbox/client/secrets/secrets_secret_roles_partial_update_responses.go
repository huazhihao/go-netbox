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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/huazhihao/go-netbox/netbox/models"
)

// SecretsSecretRolesPartialUpdateReader is a Reader for the SecretsSecretRolesPartialUpdate structure.
type SecretsSecretRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsSecretRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretRolesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretRolesPartialUpdateOK creates a SecretsSecretRolesPartialUpdateOK with default headers values
func NewSecretsSecretRolesPartialUpdateOK() *SecretsSecretRolesPartialUpdateOK {
	return &SecretsSecretRolesPartialUpdateOK{}
}

/*SecretsSecretRolesPartialUpdateOK handles this case with default header values.

SecretsSecretRolesPartialUpdateOK secrets secret roles partial update o k
*/
type SecretsSecretRolesPartialUpdateOK struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/secret-roles/{id}/][%d] secretsSecretRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretRolesPartialUpdateOK) GetPayload() *models.SecretRole {
	return o.Payload
}

func (o *SecretsSecretRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsSecretRolesPartialUpdateDefault creates a SecretsSecretRolesPartialUpdateDefault with default headers values
func NewSecretsSecretRolesPartialUpdateDefault(code int) *SecretsSecretRolesPartialUpdateDefault {
	return &SecretsSecretRolesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*SecretsSecretRolesPartialUpdateDefault handles this case with default header values.

SecretsSecretRolesPartialUpdateDefault secrets secret roles partial update default
*/
type SecretsSecretRolesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secret roles partial update default response
func (o *SecretsSecretRolesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretRolesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/secret-roles/{id}/][%d] secrets_secret-roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretRolesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretRolesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
