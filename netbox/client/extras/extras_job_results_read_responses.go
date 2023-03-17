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

// ExtrasJobResultsReadReader is a Reader for the ExtrasJobResultsRead structure.
type ExtrasJobResultsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobResultsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJobResultsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJobResultsReadOK creates a ExtrasJobResultsReadOK with default headers values
func NewExtrasJobResultsReadOK() *ExtrasJobResultsReadOK {
	return &ExtrasJobResultsReadOK{}
}

/*
ExtrasJobResultsReadOK describes a response with status code 200, with default header values.

ExtrasJobResultsReadOK extras job results read o k
*/
type ExtrasJobResultsReadOK struct {
	Payload *models.JobResult
}

// IsSuccess returns true when this extras job results read o k response has a 2xx status code
func (o *ExtrasJobResultsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras job results read o k response has a 3xx status code
func (o *ExtrasJobResultsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras job results read o k response has a 4xx status code
func (o *ExtrasJobResultsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras job results read o k response has a 5xx status code
func (o *ExtrasJobResultsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras job results read o k response a status code equal to that given
func (o *ExtrasJobResultsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras job results read o k response
func (o *ExtrasJobResultsReadOK) Code() int {
	return 200
}

func (o *ExtrasJobResultsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/job-results/{id}/][%d] extrasJobResultsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasJobResultsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/job-results/{id}/][%d] extrasJobResultsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasJobResultsReadOK) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *ExtrasJobResultsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJobResultsReadDefault creates a ExtrasJobResultsReadDefault with default headers values
func NewExtrasJobResultsReadDefault(code int) *ExtrasJobResultsReadDefault {
	return &ExtrasJobResultsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasJobResultsReadDefault describes a response with status code -1, with default header values.

ExtrasJobResultsReadDefault extras job results read default
*/
type ExtrasJobResultsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras job results read default response has a 2xx status code
func (o *ExtrasJobResultsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras job results read default response has a 3xx status code
func (o *ExtrasJobResultsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras job results read default response has a 4xx status code
func (o *ExtrasJobResultsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras job results read default response has a 5xx status code
func (o *ExtrasJobResultsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras job results read default response a status code equal to that given
func (o *ExtrasJobResultsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras job results read default response
func (o *ExtrasJobResultsReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJobResultsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/job-results/{id}/][%d] extras_job-results_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJobResultsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/job-results/{id}/][%d] extras_job-results_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJobResultsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJobResultsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
