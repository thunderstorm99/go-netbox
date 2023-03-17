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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// TenancyContactRolesReadReader is a Reader for the TenancyContactRolesRead structure.
type TenancyContactRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesReadOK creates a TenancyContactRolesReadOK with default headers values
func NewTenancyContactRolesReadOK() *TenancyContactRolesReadOK {
	return &TenancyContactRolesReadOK{}
}

/*
TenancyContactRolesReadOK describes a response with status code 200, with default header values.

TenancyContactRolesReadOK tenancy contact roles read o k
*/
type TenancyContactRolesReadOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles read o k response has a 2xx status code
func (o *TenancyContactRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles read o k response has a 3xx status code
func (o *TenancyContactRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles read o k response has a 4xx status code
func (o *TenancyContactRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles read o k response has a 5xx status code
func (o *TenancyContactRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles read o k response a status code equal to that given
func (o *TenancyContactRolesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact roles read o k response
func (o *TenancyContactRolesReadOK) Code() int {
	return 200
}

func (o *TenancyContactRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/contact-roles/{id}/][%d] tenancyContactRolesReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesReadOK) String() string {
	return fmt.Sprintf("[GET /tenancy/contact-roles/{id}/][%d] tenancyContactRolesReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesReadOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesReadDefault creates a TenancyContactRolesReadDefault with default headers values
func NewTenancyContactRolesReadDefault(code int) *TenancyContactRolesReadDefault {
	return &TenancyContactRolesReadDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesReadDefault describes a response with status code -1, with default header values.

TenancyContactRolesReadDefault tenancy contact roles read default
*/
type TenancyContactRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact roles read default response has a 2xx status code
func (o *TenancyContactRolesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles read default response has a 3xx status code
func (o *TenancyContactRolesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles read default response has a 4xx status code
func (o *TenancyContactRolesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles read default response has a 5xx status code
func (o *TenancyContactRolesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles read default response a status code equal to that given
func (o *TenancyContactRolesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact roles read default response
func (o *TenancyContactRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesReadDefault) String() string {
	return fmt.Sprintf("[GET /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
