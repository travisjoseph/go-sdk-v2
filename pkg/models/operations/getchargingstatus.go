// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetChargingStatusRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

func (o *GetChargingStatusRequest) GetVehicleID() string {
	if o == nil {
		return ""
	}
	return o.VehicleID
}

type GetChargingStatusResponse struct {
	// return EV Charge reading
	ChargeStatus *shared.ChargeStatus
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetChargingStatusResponse) GetChargeStatus() *shared.ChargeStatus {
	if o == nil {
		return nil
	}
	return o.ChargeStatus
}

func (o *GetChargingStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetChargingStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetChargingStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
