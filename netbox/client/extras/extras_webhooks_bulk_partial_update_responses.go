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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// ExtrasWebhooksBulkPartialUpdateReader is a Reader for the ExtrasWebhooksBulkPartialUpdate structure.
type ExtrasWebhooksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksBulkPartialUpdateOK creates a ExtrasWebhooksBulkPartialUpdateOK with default headers values
func NewExtrasWebhooksBulkPartialUpdateOK() *ExtrasWebhooksBulkPartialUpdateOK {
	return &ExtrasWebhooksBulkPartialUpdateOK{}
}

/*
ExtrasWebhooksBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkPartialUpdateOK extras webhooks bulk partial update o k
*/
type ExtrasWebhooksBulkPartialUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks bulk partial update o k response has a 2xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks bulk partial update o k response has a 3xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks bulk partial update o k response has a 4xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks bulk partial update o k response has a 5xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks bulk partial update o k response a status code equal to that given
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks bulk partial update o k response
func (o *ExtrasWebhooksBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extrasWebhooksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extrasWebhooksBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksBulkPartialUpdateDefault creates a ExtrasWebhooksBulkPartialUpdateDefault with default headers values
func NewExtrasWebhooksBulkPartialUpdateDefault(code int) *ExtrasWebhooksBulkPartialUpdateDefault {
	return &ExtrasWebhooksBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasWebhooksBulkPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasWebhooksBulkPartialUpdateDefault extras webhooks bulk partial update default
*/
type ExtrasWebhooksBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras webhooks bulk partial update default response has a 2xx status code
func (o *ExtrasWebhooksBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras webhooks bulk partial update default response has a 3xx status code
func (o *ExtrasWebhooksBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras webhooks bulk partial update default response has a 4xx status code
func (o *ExtrasWebhooksBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras webhooks bulk partial update default response has a 5xx status code
func (o *ExtrasWebhooksBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras webhooks bulk partial update default response a status code equal to that given
func (o *ExtrasWebhooksBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras webhooks bulk partial update default response
func (o *ExtrasWebhooksBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extras_webhooks_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extras_webhooks_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
