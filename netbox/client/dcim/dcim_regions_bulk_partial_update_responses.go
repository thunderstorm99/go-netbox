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

// DcimRegionsBulkPartialUpdateReader is a Reader for the DcimRegionsBulkPartialUpdate structure.
type DcimRegionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsBulkPartialUpdateOK creates a DcimRegionsBulkPartialUpdateOK with default headers values
func NewDcimRegionsBulkPartialUpdateOK() *DcimRegionsBulkPartialUpdateOK {
	return &DcimRegionsBulkPartialUpdateOK{}
}

/*
DcimRegionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRegionsBulkPartialUpdateOK dcim regions bulk partial update o k
*/
type DcimRegionsBulkPartialUpdateOK struct {
	Payload *models.Region
}

// IsSuccess returns true when this dcim regions bulk partial update o k response has a 2xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions bulk partial update o k response has a 3xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk partial update o k response has a 4xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions bulk partial update o k response has a 5xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk partial update o k response a status code equal to that given
func (o *DcimRegionsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim regions bulk partial update o k response
func (o *DcimRegionsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimRegionsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsBulkPartialUpdateDefault creates a DcimRegionsBulkPartialUpdateDefault with default headers values
func NewDcimRegionsBulkPartialUpdateDefault(code int) *DcimRegionsBulkPartialUpdateDefault {
	return &DcimRegionsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRegionsBulkPartialUpdateDefault dcim regions bulk partial update default
*/
type DcimRegionsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim regions bulk partial update default response has a 2xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions bulk partial update default response has a 3xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions bulk partial update default response has a 4xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions bulk partial update default response has a 5xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions bulk partial update default response a status code equal to that given
func (o *DcimRegionsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim regions bulk partial update default response
func (o *DcimRegionsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRegionsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcim_regions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcim_regions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
