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

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// DcimVirtualChassisCreateReader is a Reader for the DcimVirtualChassisCreate structure.
type DcimVirtualChassisCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimVirtualChassisCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisCreateCreated creates a DcimVirtualChassisCreateCreated with default headers values
func NewDcimVirtualChassisCreateCreated() *DcimVirtualChassisCreateCreated {
	return &DcimVirtualChassisCreateCreated{}
}

/*
DcimVirtualChassisCreateCreated describes a response with status code 201, with default header values.

DcimVirtualChassisCreateCreated dcim virtual chassis create created
*/
type DcimVirtualChassisCreateCreated struct {
	Payload *models.VirtualChassis
}

// IsSuccess returns true when this dcim virtual chassis create created response has a 2xx status code
func (o *DcimVirtualChassisCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis create created response has a 3xx status code
func (o *DcimVirtualChassisCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis create created response has a 4xx status code
func (o *DcimVirtualChassisCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis create created response has a 5xx status code
func (o *DcimVirtualChassisCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis create created response a status code equal to that given
func (o *DcimVirtualChassisCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim virtual chassis create created response
func (o *DcimVirtualChassisCreateCreated) Code() int {
	return 201
}

func (o *DcimVirtualChassisCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimVirtualChassisCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimVirtualChassisCreateCreated) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisCreateDefault creates a DcimVirtualChassisCreateDefault with default headers values
func NewDcimVirtualChassisCreateDefault(code int) *DcimVirtualChassisCreateDefault {
	return &DcimVirtualChassisCreateDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualChassisCreateDefault describes a response with status code -1, with default header values.

DcimVirtualChassisCreateDefault dcim virtual chassis create default
*/
type DcimVirtualChassisCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim virtual chassis create default response has a 2xx status code
func (o *DcimVirtualChassisCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual chassis create default response has a 3xx status code
func (o *DcimVirtualChassisCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual chassis create default response has a 4xx status code
func (o *DcimVirtualChassisCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual chassis create default response has a 5xx status code
func (o *DcimVirtualChassisCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual chassis create default response a status code equal to that given
func (o *DcimVirtualChassisCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim virtual chassis create default response
func (o *DcimVirtualChassisCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcim_virtual-chassis_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcim_virtual-chassis_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
