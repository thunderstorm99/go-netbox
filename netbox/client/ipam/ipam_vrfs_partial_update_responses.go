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

// IpamVrfsPartialUpdateReader is a Reader for the IpamVrfsPartialUpdate structure.
type IpamVrfsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsPartialUpdateOK creates a IpamVrfsPartialUpdateOK with default headers values
func NewIpamVrfsPartialUpdateOK() *IpamVrfsPartialUpdateOK {
	return &IpamVrfsPartialUpdateOK{}
}

/*
IpamVrfsPartialUpdateOK describes a response with status code 200, with default header values.

IpamVrfsPartialUpdateOK ipam vrfs partial update o k
*/
type IpamVrfsPartialUpdateOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs partial update o k response has a 2xx status code
func (o *IpamVrfsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs partial update o k response has a 3xx status code
func (o *IpamVrfsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs partial update o k response has a 4xx status code
func (o *IpamVrfsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs partial update o k response has a 5xx status code
func (o *IpamVrfsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs partial update o k response a status code equal to that given
func (o *IpamVrfsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vrfs partial update o k response
func (o *IpamVrfsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamVrfsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsPartialUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsPartialUpdateDefault creates a IpamVrfsPartialUpdateDefault with default headers values
func NewIpamVrfsPartialUpdateDefault(code int) *IpamVrfsPartialUpdateDefault {
	return &IpamVrfsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamVrfsPartialUpdateDefault ipam vrfs partial update default
*/
type IpamVrfsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vrfs partial update default response has a 2xx status code
func (o *IpamVrfsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs partial update default response has a 3xx status code
func (o *IpamVrfsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs partial update default response has a 4xx status code
func (o *IpamVrfsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs partial update default response has a 5xx status code
func (o *IpamVrfsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs partial update default response a status code equal to that given
func (o *IpamVrfsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vrfs partial update default response
func (o *IpamVrfsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipam_vrfs_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipam_vrfs_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
