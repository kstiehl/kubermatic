// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListAWSZonesNoCredentialsReader is a Reader for the ListAWSZonesNoCredentials structure.
type ListAWSZonesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSZonesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAWSZonesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAWSZonesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSZonesNoCredentialsOK creates a ListAWSZonesNoCredentialsOK with default headers values
func NewListAWSZonesNoCredentialsOK() *ListAWSZonesNoCredentialsOK {
	return &ListAWSZonesNoCredentialsOK{}
}

/*ListAWSZonesNoCredentialsOK handles this case with default header values.

AWSZoneList
*/
type ListAWSZonesNoCredentialsOK struct {
	Payload models.AWSZoneList
}

func (o *ListAWSZonesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/zones][%d] listAWSZonesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListAWSZonesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSZonesNoCredentialsDefault creates a ListAWSZonesNoCredentialsDefault with default headers values
func NewListAWSZonesNoCredentialsDefault(code int) *ListAWSZonesNoCredentialsDefault {
	return &ListAWSZonesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListAWSZonesNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListAWSZonesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s zones no credentials default response
func (o *ListAWSZonesNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListAWSZonesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/zones][%d] listAWSZonesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSZonesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
