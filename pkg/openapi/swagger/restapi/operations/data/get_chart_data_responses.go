// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/IoTube-analytics/go-iotube-analytics/pkg/openapi/swagger/models"
)

// GetChartDataOKCode is the HTTP code returned for type GetChartDataOK
const GetChartDataOKCode int = 200

/*GetChartDataOK successful operation

swagger:response getChartDataOK
*/
type GetChartDataOK struct {

	/*
	  In: Body
	*/
	Payload models.ChartData `json:"body,omitempty"`
}

// NewGetChartDataOK creates GetChartDataOK with default headers values
func NewGetChartDataOK() *GetChartDataOK {

	return &GetChartDataOK{}
}

// WithPayload adds the payload to the get chart data o k response
func (o *GetChartDataOK) WithPayload(payload models.ChartData) *GetChartDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart data o k response
func (o *GetChartDataOK) SetPayload(payload models.ChartData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetChartDataNotFoundCode is the HTTP code returned for type GetChartDataNotFound
const GetChartDataNotFoundCode int = 404

/*GetChartDataNotFound not found

swagger:response getChartDataNotFound
*/
type GetChartDataNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetChartDataNotFound creates GetChartDataNotFound with default headers values
func NewGetChartDataNotFound() *GetChartDataNotFound {

	return &GetChartDataNotFound{}
}

// WithPayload adds the payload to the get chart data not found response
func (o *GetChartDataNotFound) WithPayload(payload *models.APIResponse) *GetChartDataNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart data not found response
func (o *GetChartDataNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartDataNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
