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

// DcimModulesUpdateReader is a Reader for the DcimModulesUpdate structure.
type DcimModulesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModulesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModulesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModulesUpdateOK creates a DcimModulesUpdateOK with default headers values
func NewDcimModulesUpdateOK() *DcimModulesUpdateOK {
	return &DcimModulesUpdateOK{}
}

/*
DcimModulesUpdateOK describes a response with status code 200, with default header values.

DcimModulesUpdateOK dcim modules update o k
*/
type DcimModulesUpdateOK struct {
	Payload *models.Module
}

// IsSuccess returns true when this dcim modules update o k response has a 2xx status code
func (o *DcimModulesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim modules update o k response has a 3xx status code
func (o *DcimModulesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules update o k response has a 4xx status code
func (o *DcimModulesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim modules update o k response has a 5xx status code
func (o *DcimModulesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules update o k response a status code equal to that given
func (o *DcimModulesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim modules update o k response
func (o *DcimModulesUpdateOK) Code() int {
	return 200
}

func (o *DcimModulesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/modules/{id}/][%d] dcimModulesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModulesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/modules/{id}/][%d] dcimModulesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModulesUpdateOK) GetPayload() *models.Module {
	return o.Payload
}

func (o *DcimModulesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Module)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModulesUpdateDefault creates a DcimModulesUpdateDefault with default headers values
func NewDcimModulesUpdateDefault(code int) *DcimModulesUpdateDefault {
	return &DcimModulesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModulesUpdateDefault describes a response with status code -1, with default header values.

DcimModulesUpdateDefault dcim modules update default
*/
type DcimModulesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim modules update default response has a 2xx status code
func (o *DcimModulesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim modules update default response has a 3xx status code
func (o *DcimModulesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim modules update default response has a 4xx status code
func (o *DcimModulesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim modules update default response has a 5xx status code
func (o *DcimModulesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim modules update default response a status code equal to that given
func (o *DcimModulesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim modules update default response
func (o *DcimModulesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModulesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/modules/{id}/][%d] dcim_modules_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/modules/{id}/][%d] dcim_modules_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
