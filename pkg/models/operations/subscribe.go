// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type SubscribeRequest struct {
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
	WebhookID   string              `pathParam:"style=simple,explode=false,name=webhookId"`
	WebhookInfo *shared.WebhookInfo `request:"mediaType=application/json"`
}

type SubscribeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Compatibility
	SuccessResponse *shared.SuccessResponse
}
