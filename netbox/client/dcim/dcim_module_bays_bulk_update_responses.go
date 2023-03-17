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

// DcimModuleBaysBulkUpdateReader is a Reader for the DcimModuleBaysBulkUpdate structure.
type DcimModuleBaysBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysBulkUpdateOK creates a DcimModuleBaysBulkUpdateOK with default headers values
func NewDcimModuleBaysBulkUpdateOK() *DcimModuleBaysBulkUpdateOK {
	return &DcimModuleBaysBulkUpdateOK{}
}

/*
DcimModuleBaysBulkUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysBulkUpdateOK dcim module bays bulk update o k
*/
type DcimModuleBaysBulkUpdateOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays bulk update o k response has a 2xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays bulk update o k response has a 3xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays bulk update o k response has a 4xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays bulk update o k response has a 5xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays bulk update o k response a status code equal to that given
func (o *DcimModuleBaysBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim module bays bulk update o k response
func (o *DcimModuleBaysBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimModuleBaysBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcimModuleBaysBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcimModuleBaysBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysBulkUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysBulkUpdateDefault creates a DcimModuleBaysBulkUpdateDefault with default headers values
func NewDcimModuleBaysBulkUpdateDefault(code int) *DcimModuleBaysBulkUpdateDefault {
	return &DcimModuleBaysBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysBulkUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBaysBulkUpdateDefault dcim module bays bulk update default
*/
type DcimModuleBaysBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim module bays bulk update default response has a 2xx status code
func (o *DcimModuleBaysBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays bulk update default response has a 3xx status code
func (o *DcimModuleBaysBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays bulk update default response has a 4xx status code
func (o *DcimModuleBaysBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays bulk update default response has a 5xx status code
func (o *DcimModuleBaysBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays bulk update default response a status code equal to that given
func (o *DcimModuleBaysBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim module bays bulk update default response
func (o *DcimModuleBaysBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleBaysBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcim_module-bays_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcim_module-bays_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
