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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// CircuitsProviderNetworksBulkUpdateReader is a Reader for the CircuitsProviderNetworksBulkUpdate structure.
type CircuitsProviderNetworksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProviderNetworksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProviderNetworksBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProviderNetworksBulkUpdateOK creates a CircuitsProviderNetworksBulkUpdateOK with default headers values
func NewCircuitsProviderNetworksBulkUpdateOK() *CircuitsProviderNetworksBulkUpdateOK {
	return &CircuitsProviderNetworksBulkUpdateOK{}
}

/*
CircuitsProviderNetworksBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsProviderNetworksBulkUpdateOK circuits provider networks bulk update o k
*/
type CircuitsProviderNetworksBulkUpdateOK struct {
	Payload *models.ProviderNetwork
}

// IsSuccess returns true when this circuits provider networks bulk update o k response has a 2xx status code
func (o *CircuitsProviderNetworksBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits provider networks bulk update o k response has a 3xx status code
func (o *CircuitsProviderNetworksBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks bulk update o k response has a 4xx status code
func (o *CircuitsProviderNetworksBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits provider networks bulk update o k response has a 5xx status code
func (o *CircuitsProviderNetworksBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks bulk update o k response a status code equal to that given
func (o *CircuitsProviderNetworksBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits provider networks bulk update o k response
func (o *CircuitsProviderNetworksBulkUpdateOK) Code() int {
	return 200
}

func (o *CircuitsProviderNetworksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/provider-networks/][%d] circuitsProviderNetworksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /circuits/provider-networks/][%d] circuitsProviderNetworksBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksBulkUpdateOK) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksBulkUpdateDefault creates a CircuitsProviderNetworksBulkUpdateDefault with default headers values
func NewCircuitsProviderNetworksBulkUpdateDefault(code int) *CircuitsProviderNetworksBulkUpdateDefault {
	return &CircuitsProviderNetworksBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProviderNetworksBulkUpdateDefault describes a response with status code -1, with default header values.

CircuitsProviderNetworksBulkUpdateDefault circuits provider networks bulk update default
*/
type CircuitsProviderNetworksBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this circuits provider networks bulk update default response has a 2xx status code
func (o *CircuitsProviderNetworksBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits provider networks bulk update default response has a 3xx status code
func (o *CircuitsProviderNetworksBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits provider networks bulk update default response has a 4xx status code
func (o *CircuitsProviderNetworksBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits provider networks bulk update default response has a 5xx status code
func (o *CircuitsProviderNetworksBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits provider networks bulk update default response a status code equal to that given
func (o *CircuitsProviderNetworksBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the circuits provider networks bulk update default response
func (o *CircuitsProviderNetworksBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProviderNetworksBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/provider-networks/][%d] circuits_provider-networks_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /circuits/provider-networks/][%d] circuits_provider-networks_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
