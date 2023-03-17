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

// DcimInventoryItemTemplatesBulkUpdateReader is a Reader for the DcimInventoryItemTemplatesBulkUpdate structure.
type DcimInventoryItemTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesBulkUpdateOK creates a DcimInventoryItemTemplatesBulkUpdateOK with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateOK() *DcimInventoryItemTemplatesBulkUpdateOK {
	return &DcimInventoryItemTemplatesBulkUpdateOK{}
}

/*
DcimInventoryItemTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesBulkUpdateOK dcim inventory item templates bulk update o k
*/
type DcimInventoryItemTemplatesBulkUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates bulk update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates bulk update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates bulk update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim inventory item templates bulk update o k response
func (o *DcimInventoryItemTemplatesBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkUpdateDefault creates a DcimInventoryItemTemplatesBulkUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateDefault(code int) *DcimInventoryItemTemplatesBulkUpdateDefault {
	return &DcimInventoryItemTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemTemplatesBulkUpdateDefault dcim inventory item templates bulk update default
*/
type DcimInventoryItemTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item templates bulk update default response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates bulk update default response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates bulk update default response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates bulk update default response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates bulk update default response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim inventory item templates bulk update default response
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
