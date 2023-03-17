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

// DcimFrontPortsReadReader is a Reader for the DcimFrontPortsRead structure.
type DcimFrontPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsReadOK creates a DcimFrontPortsReadOK with default headers values
func NewDcimFrontPortsReadOK() *DcimFrontPortsReadOK {
	return &DcimFrontPortsReadOK{}
}

/*
DcimFrontPortsReadOK describes a response with status code 200, with default header values.

DcimFrontPortsReadOK dcim front ports read o k
*/
type DcimFrontPortsReadOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports read o k response has a 2xx status code
func (o *DcimFrontPortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports read o k response has a 3xx status code
func (o *DcimFrontPortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports read o k response has a 4xx status code
func (o *DcimFrontPortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports read o k response has a 5xx status code
func (o *DcimFrontPortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports read o k response a status code equal to that given
func (o *DcimFrontPortsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front ports read o k response
func (o *DcimFrontPortsReadOK) Code() int {
	return 200
}

func (o *DcimFrontPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcimFrontPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcimFrontPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsReadOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsReadDefault creates a DcimFrontPortsReadDefault with default headers values
func NewDcimFrontPortsReadDefault(code int) *DcimFrontPortsReadDefault {
	return &DcimFrontPortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortsReadDefault describes a response with status code -1, with default header values.

DcimFrontPortsReadDefault dcim front ports read default
*/
type DcimFrontPortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim front ports read default response has a 2xx status code
func (o *DcimFrontPortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front ports read default response has a 3xx status code
func (o *DcimFrontPortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front ports read default response has a 4xx status code
func (o *DcimFrontPortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front ports read default response has a 5xx status code
func (o *DcimFrontPortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front ports read default response a status code equal to that given
func (o *DcimFrontPortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim front ports read default response
func (o *DcimFrontPortsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcim_front-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcim_front-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
