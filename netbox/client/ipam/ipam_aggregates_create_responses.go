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

// IpamAggregatesCreateReader is a Reader for the IpamAggregatesCreate structure.
type IpamAggregatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamAggregatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAggregatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesCreateCreated creates a IpamAggregatesCreateCreated with default headers values
func NewIpamAggregatesCreateCreated() *IpamAggregatesCreateCreated {
	return &IpamAggregatesCreateCreated{}
}

/*
IpamAggregatesCreateCreated describes a response with status code 201, with default header values.

IpamAggregatesCreateCreated ipam aggregates create created
*/
type IpamAggregatesCreateCreated struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates create created response has a 2xx status code
func (o *IpamAggregatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates create created response has a 3xx status code
func (o *IpamAggregatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates create created response has a 4xx status code
func (o *IpamAggregatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates create created response has a 5xx status code
func (o *IpamAggregatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates create created response a status code equal to that given
func (o *IpamAggregatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam aggregates create created response
func (o *IpamAggregatesCreateCreated) Code() int {
	return 201
}

func (o *IpamAggregatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/aggregates/][%d] ipamAggregatesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamAggregatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/aggregates/][%d] ipamAggregatesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamAggregatesCreateCreated) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesCreateDefault creates a IpamAggregatesCreateDefault with default headers values
func NewIpamAggregatesCreateDefault(code int) *IpamAggregatesCreateDefault {
	return &IpamAggregatesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesCreateDefault describes a response with status code -1, with default header values.

IpamAggregatesCreateDefault ipam aggregates create default
*/
type IpamAggregatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam aggregates create default response has a 2xx status code
func (o *IpamAggregatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates create default response has a 3xx status code
func (o *IpamAggregatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates create default response has a 4xx status code
func (o *IpamAggregatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates create default response has a 5xx status code
func (o *IpamAggregatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates create default response a status code equal to that given
func (o *IpamAggregatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam aggregates create default response
func (o *IpamAggregatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamAggregatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/aggregates/][%d] ipam_aggregates_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/aggregates/][%d] ipam_aggregates_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
