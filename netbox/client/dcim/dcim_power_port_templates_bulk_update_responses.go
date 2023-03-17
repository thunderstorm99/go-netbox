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

// DcimPowerPortTemplatesBulkUpdateReader is a Reader for the DcimPowerPortTemplatesBulkUpdate structure.
type DcimPowerPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesBulkUpdateOK creates a DcimPowerPortTemplatesBulkUpdateOK with default headers values
func NewDcimPowerPortTemplatesBulkUpdateOK() *DcimPowerPortTemplatesBulkUpdateOK {
	return &DcimPowerPortTemplatesBulkUpdateOK{}
}

/*
DcimPowerPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesBulkUpdateOK dcim power port templates bulk update o k
*/
type DcimPowerPortTemplatesBulkUpdateOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates bulk update o k response has a 2xx status code
func (o *DcimPowerPortTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates bulk update o k response has a 3xx status code
func (o *DcimPowerPortTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates bulk update o k response has a 4xx status code
func (o *DcimPowerPortTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates bulk update o k response has a 5xx status code
func (o *DcimPowerPortTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates bulk update o k response a status code equal to that given
func (o *DcimPowerPortTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates bulk update o k response
func (o *DcimPowerPortTemplatesBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesBulkUpdateDefault creates a DcimPowerPortTemplatesBulkUpdateDefault with default headers values
func NewDcimPowerPortTemplatesBulkUpdateDefault(code int) *DcimPowerPortTemplatesBulkUpdateDefault {
	return &DcimPowerPortTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortTemplatesBulkUpdateDefault dcim power port templates bulk update default
*/
type DcimPowerPortTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates bulk update default response has a 2xx status code
func (o *DcimPowerPortTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates bulk update default response has a 3xx status code
func (o *DcimPowerPortTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates bulk update default response has a 4xx status code
func (o *DcimPowerPortTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates bulk update default response has a 5xx status code
func (o *DcimPowerPortTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates bulk update default response a status code equal to that given
func (o *DcimPowerPortTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power port templates bulk update default response
func (o *DcimPowerPortTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/][%d] dcim_power-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/][%d] dcim_power-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
