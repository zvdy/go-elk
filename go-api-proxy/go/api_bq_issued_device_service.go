// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Leasing
 *
 * The leasing products enables customers to finance equipment purchases using the leased item as collateral when necessary
 *
 * API version: 12.0.0
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// BQIssuedDeviceAPIService is a service that implements the logic for the BQIssuedDeviceAPIServicer
// This service should implement the business logic for every endpoint for the BQIssuedDeviceAPI API.
// Include any external packages or services that will be required by this service.
type BQIssuedDeviceAPIService struct {
}

// NewBQIssuedDeviceAPIService creates a default api service
func NewBQIssuedDeviceAPIService() *BQIssuedDeviceAPIService {
	return &BQIssuedDeviceAPIService{}
}

// InitiateIssuedDevice - InBQ Instantiate a new Issued Device
func (s *BQIssuedDeviceAPIService) InitiateIssuedDevice(ctx context.Context, leasingid string, issuedDevice IssuedDevice) (ImplResponse, error) {
	// TODO - update InitiateIssuedDevice with the required logic for this service method.
	// Add api_bq_issued_device_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IssuedDevice{}) or use other options such as http.Ok ...
	// return Response(200, IssuedDevice{}), nil

	// TODO: Uncomment the next line to return response Response(400, HttpError{}) or use other options such as http.Ok ...
	// return Response(400, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(401, HttpError{}) or use other options such as http.Ok ...
	// return Response(401, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(403, HttpError{}) or use other options such as http.Ok ...
	// return Response(403, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(404, HttpError{}) or use other options such as http.Ok ...
	// return Response(404, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(429, HttpError{}) or use other options such as http.Ok ...
	// return Response(429, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(500, HttpError{}) or use other options such as http.Ok ...
	// return Response(500, HttpError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("InitiateIssuedDevice method not implemented")
}

// RetrieveIssuedDevice - ReBQ Retrieve details about any aspect of Issued Device
func (s *BQIssuedDeviceAPIService) RetrieveIssuedDevice(ctx context.Context, leasingid string, issueddeviceid string) (ImplResponse, error) {
	// TODO - update RetrieveIssuedDevice with the required logic for this service method.
	// Add api_bq_issued_device_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IssuedDevice{}) or use other options such as http.Ok ...
	// return Response(200, IssuedDevice{}), nil

	// TODO: Uncomment the next line to return response Response(400, HttpError{}) or use other options such as http.Ok ...
	// return Response(400, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(401, HttpError{}) or use other options such as http.Ok ...
	// return Response(401, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(403, HttpError{}) or use other options such as http.Ok ...
	// return Response(403, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(404, HttpError{}) or use other options such as http.Ok ...
	// return Response(404, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(429, HttpError{}) or use other options such as http.Ok ...
	// return Response(429, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(500, HttpError{}) or use other options such as http.Ok ...
	// return Response(500, HttpError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RetrieveIssuedDevice method not implemented")
}

// UpdateIssuedDevice - UpBQ Update details relating to Issued Device
func (s *BQIssuedDeviceAPIService) UpdateIssuedDevice(ctx context.Context, leasingid string, issueddeviceid string, issuedDevice IssuedDevice) (ImplResponse, error) {
	// TODO - update UpdateIssuedDevice with the required logic for this service method.
	// Add api_bq_issued_device_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IssuedDevice{}) or use other options such as http.Ok ...
	// return Response(200, IssuedDevice{}), nil

	// TODO: Uncomment the next line to return response Response(400, HttpError{}) or use other options such as http.Ok ...
	// return Response(400, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(401, HttpError{}) or use other options such as http.Ok ...
	// return Response(401, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(403, HttpError{}) or use other options such as http.Ok ...
	// return Response(403, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(404, HttpError{}) or use other options such as http.Ok ...
	// return Response(404, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(429, HttpError{}) or use other options such as http.Ok ...
	// return Response(429, HttpError{}), nil

	// TODO: Uncomment the next line to return response Response(500, HttpError{}) or use other options such as http.Ok ...
	// return Response(500, HttpError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateIssuedDevice method not implemented")
}
