// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchClusterReader is a Reader for the PatchCluster structure.
type PatchClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchClusterOK creates a PatchClusterOK with default headers values
func NewPatchClusterOK() *PatchClusterOK {
	return &PatchClusterOK{}
}

/* PatchClusterOK describes a response with status code 200, with default header values.

Cluster
*/
type PatchClusterOK struct {
	Payload *models.Cluster
}

func (o *PatchClusterOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] patchClusterOK  %+v", 200, o.Payload)
}
func (o *PatchClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *PatchClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClusterUnauthorized creates a PatchClusterUnauthorized with default headers values
func NewPatchClusterUnauthorized() *PatchClusterUnauthorized {
	return &PatchClusterUnauthorized{}
}

/* PatchClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterUnauthorized struct {
}

func (o *PatchClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] patchClusterUnauthorized ", 401)
}

func (o *PatchClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterForbidden creates a PatchClusterForbidden with default headers values
func NewPatchClusterForbidden() *PatchClusterForbidden {
	return &PatchClusterForbidden{}
}

/* PatchClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchClusterForbidden struct {
}

func (o *PatchClusterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] patchClusterForbidden ", 403)
}

func (o *PatchClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterDefault creates a PatchClusterDefault with default headers values
func NewPatchClusterDefault(code int) *PatchClusterDefault {
	return &PatchClusterDefault{
		_statusCode: code,
	}
}

/* PatchClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch cluster default response
func (o *PatchClusterDefault) Code() int {
	return o._statusCode
}

func (o *PatchClusterDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] patchCluster default  %+v", o._statusCode, o.Payload)
}
func (o *PatchClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
