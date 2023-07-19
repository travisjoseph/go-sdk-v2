// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetFuelTankRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetFuelTankRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetFuelTankResponse struct {
	ContentType string
	// return fuel tank reading
	FuelTank    *shared.FuelTank
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetFuelTankResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFuelTankResponse) GetFuelTank() *shared.FuelTank {
	if o == nil {
		return nil
	}
	return o.FuelTank
}

func (o *GetFuelTankResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFuelTankResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
