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

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// IpamServicesBulkPartialUpdateReader is a Reader for the IpamServicesBulkPartialUpdate structure.
type IpamServicesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesBulkPartialUpdateOK creates a IpamServicesBulkPartialUpdateOK with default headers values
func NewIpamServicesBulkPartialUpdateOK() *IpamServicesBulkPartialUpdateOK {
	return &IpamServicesBulkPartialUpdateOK{}
}

/*
IpamServicesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamServicesBulkPartialUpdateOK ipam services bulk partial update o k
*/
type IpamServicesBulkPartialUpdateOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services bulk partial update o k response has a 2xx status code
func (o *IpamServicesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services bulk partial update o k response has a 3xx status code
func (o *IpamServicesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services bulk partial update o k response has a 4xx status code
func (o *IpamServicesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services bulk partial update o k response has a 5xx status code
func (o *IpamServicesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services bulk partial update o k response a status code equal to that given
func (o *IpamServicesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam services bulk partial update o k response
func (o *IpamServicesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamServicesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/services/][%d] ipamServicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/services/][%d] ipamServicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesBulkPartialUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesBulkPartialUpdateDefault creates a IpamServicesBulkPartialUpdateDefault with default headers values
func NewIpamServicesBulkPartialUpdateDefault(code int) *IpamServicesBulkPartialUpdateDefault {
	return &IpamServicesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServicesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamServicesBulkPartialUpdateDefault ipam services bulk partial update default
*/
type IpamServicesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam services bulk partial update default response has a 2xx status code
func (o *IpamServicesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services bulk partial update default response has a 3xx status code
func (o *IpamServicesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services bulk partial update default response has a 4xx status code
func (o *IpamServicesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services bulk partial update default response has a 5xx status code
func (o *IpamServicesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services bulk partial update default response a status code equal to that given
func (o *IpamServicesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam services bulk partial update default response
func (o *IpamServicesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/services/][%d] ipam_services_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/services/][%d] ipam_services_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
