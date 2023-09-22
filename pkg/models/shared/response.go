// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smartcar/go-sdk-v2/pkg/utils"
)

type ResponseBodyType string

const (
	ResponseBodyTypeLocation              ResponseBodyType = "Location"
	ResponseBodyTypeOdometer              ResponseBodyType = "Odometer"
	ResponseBodyTypeFuelTank              ResponseBodyType = "FuelTank"
	ResponseBodyTypeTirePressure          ResponseBodyType = "TirePressure"
	ResponseBodyTypeEngineOil             ResponseBodyType = "EngineOil"
	ResponseBodyTypeChargeStatus          ResponseBodyType = "ChargeStatus"
	ResponseBodyTypeChargeLimit           ResponseBodyType = "ChargeLimit"
	ResponseBodyTypeChargeTime            ResponseBodyType = "ChargeTime"
	ResponseBodyTypeChargeVoltage         ResponseBodyType = "ChargeVoltage"
	ResponseBodyTypeBatteryLevel          ResponseBodyType = "BatteryLevel"
	ResponseBodyTypeBatteryCapacity       ResponseBodyType = "BatteryCapacity"
	ResponseBodyTypeCompatibilityResponse ResponseBodyType = "CompatibilityResponse"
	ResponseBodyTypeVinInfo               ResponseBodyType = "VinInfo"
	ResponseBodyTypeUserInfo              ResponseBodyType = "UserInfo"
	ResponseBodyTypeSuccessResponse       ResponseBodyType = "SuccessResponse"
	ResponseBodyTypeSecurityRead          ResponseBodyType = "SecurityRead"
)

type ResponseBody struct {
	Location              *Location
	Odometer              *Odometer
	FuelTank              *FuelTank
	TirePressure          *TirePressure
	EngineOil             *EngineOil
	ChargeStatus          *ChargeStatus
	ChargeLimit           *ChargeLimit
	ChargeTime            *ChargeTime
	ChargeVoltage         *ChargeVoltage
	BatteryLevel          *BatteryLevel
	BatteryCapacity       *BatteryCapacity
	CompatibilityResponse *CompatibilityResponse
	VinInfo               *VinInfo
	UserInfo              *UserInfo
	SuccessResponse       *SuccessResponse
	SecurityRead          *SecurityRead

	Type ResponseBodyType
}

func CreateResponseBodyLocation(location Location) ResponseBody {
	typ := ResponseBodyTypeLocation

	return ResponseBody{
		Location: &location,
		Type:     typ,
	}
}

func CreateResponseBodyOdometer(odometer Odometer) ResponseBody {
	typ := ResponseBodyTypeOdometer

	return ResponseBody{
		Odometer: &odometer,
		Type:     typ,
	}
}

func CreateResponseBodyFuelTank(fuelTank FuelTank) ResponseBody {
	typ := ResponseBodyTypeFuelTank

	return ResponseBody{
		FuelTank: &fuelTank,
		Type:     typ,
	}
}

func CreateResponseBodyTirePressure(tirePressure TirePressure) ResponseBody {
	typ := ResponseBodyTypeTirePressure

	return ResponseBody{
		TirePressure: &tirePressure,
		Type:         typ,
	}
}

func CreateResponseBodyEngineOil(engineOil EngineOil) ResponseBody {
	typ := ResponseBodyTypeEngineOil

	return ResponseBody{
		EngineOil: &engineOil,
		Type:      typ,
	}
}

func CreateResponseBodyChargeStatus(chargeStatus ChargeStatus) ResponseBody {
	typ := ResponseBodyTypeChargeStatus

	return ResponseBody{
		ChargeStatus: &chargeStatus,
		Type:         typ,
	}
}

func CreateResponseBodyChargeLimit(chargeLimit ChargeLimit) ResponseBody {
	typ := ResponseBodyTypeChargeLimit

	return ResponseBody{
		ChargeLimit: &chargeLimit,
		Type:        typ,
	}
}

func CreateResponseBodyChargeTime(chargeTime ChargeTime) ResponseBody {
	typ := ResponseBodyTypeChargeTime

	return ResponseBody{
		ChargeTime: &chargeTime,
		Type:       typ,
	}
}

func CreateResponseBodyChargeVoltage(chargeVoltage ChargeVoltage) ResponseBody {
	typ := ResponseBodyTypeChargeVoltage

	return ResponseBody{
		ChargeVoltage: &chargeVoltage,
		Type:          typ,
	}
}

func CreateResponseBodyBatteryLevel(batteryLevel BatteryLevel) ResponseBody {
	typ := ResponseBodyTypeBatteryLevel

	return ResponseBody{
		BatteryLevel: &batteryLevel,
		Type:         typ,
	}
}

func CreateResponseBodyBatteryCapacity(batteryCapacity BatteryCapacity) ResponseBody {
	typ := ResponseBodyTypeBatteryCapacity

	return ResponseBody{
		BatteryCapacity: &batteryCapacity,
		Type:            typ,
	}
}

func CreateResponseBodyCompatibilityResponse(compatibilityResponse CompatibilityResponse) ResponseBody {
	typ := ResponseBodyTypeCompatibilityResponse

	return ResponseBody{
		CompatibilityResponse: &compatibilityResponse,
		Type:                  typ,
	}
}

func CreateResponseBodyVinInfo(vinInfo VinInfo) ResponseBody {
	typ := ResponseBodyTypeVinInfo

	return ResponseBody{
		VinInfo: &vinInfo,
		Type:    typ,
	}
}

func CreateResponseBodyUserInfo(userInfo UserInfo) ResponseBody {
	typ := ResponseBodyTypeUserInfo

	return ResponseBody{
		UserInfo: &userInfo,
		Type:     typ,
	}
}

func CreateResponseBodySuccessResponse(successResponse SuccessResponse) ResponseBody {
	typ := ResponseBodyTypeSuccessResponse

	return ResponseBody{
		SuccessResponse: &successResponse,
		Type:            typ,
	}
}

func CreateResponseBodySecurityRead(securityRead SecurityRead) ResponseBody {
	typ := ResponseBodyTypeSecurityRead

	return ResponseBody{
		SecurityRead: &securityRead,
		Type:         typ,
	}
}

func (u *ResponseBody) UnmarshalJSON(data []byte) error {

	odometer := new(Odometer)
	if err := utils.UnmarshalJSON(data, &odometer, "", true, true); err == nil {
		u.Odometer = odometer
		u.Type = ResponseBodyTypeOdometer
		return nil
	}

	engineOil := new(EngineOil)
	if err := utils.UnmarshalJSON(data, &engineOil, "", true, true); err == nil {
		u.EngineOil = engineOil
		u.Type = ResponseBodyTypeEngineOil
		return nil
	}

	chargeLimit := new(ChargeLimit)
	if err := utils.UnmarshalJSON(data, &chargeLimit, "", true, true); err == nil {
		u.ChargeLimit = chargeLimit
		u.Type = ResponseBodyTypeChargeLimit
		return nil
	}

	chargeTime := new(ChargeTime)
	if err := utils.UnmarshalJSON(data, &chargeTime, "", true, true); err == nil {
		u.ChargeTime = chargeTime
		u.Type = ResponseBodyTypeChargeTime
		return nil
	}

	chargeVoltage := new(ChargeVoltage)
	if err := utils.UnmarshalJSON(data, &chargeVoltage, "", true, true); err == nil {
		u.ChargeVoltage = chargeVoltage
		u.Type = ResponseBodyTypeChargeVoltage
		return nil
	}

	batteryCapacity := new(BatteryCapacity)
	if err := utils.UnmarshalJSON(data, &batteryCapacity, "", true, true); err == nil {
		u.BatteryCapacity = batteryCapacity
		u.Type = ResponseBodyTypeBatteryCapacity
		return nil
	}

	vinInfo := new(VinInfo)
	if err := utils.UnmarshalJSON(data, &vinInfo, "", true, true); err == nil {
		u.VinInfo = vinInfo
		u.Type = ResponseBodyTypeVinInfo
		return nil
	}

	userInfo := new(UserInfo)
	if err := utils.UnmarshalJSON(data, &userInfo, "", true, true); err == nil {
		u.UserInfo = userInfo
		u.Type = ResponseBodyTypeUserInfo
		return nil
	}

	location := new(Location)
	if err := utils.UnmarshalJSON(data, &location, "", true, true); err == nil {
		u.Location = location
		u.Type = ResponseBodyTypeLocation
		return nil
	}

	chargeStatus := new(ChargeStatus)
	if err := utils.UnmarshalJSON(data, &chargeStatus, "", true, true); err == nil {
		u.ChargeStatus = chargeStatus
		u.Type = ResponseBodyTypeChargeStatus
		return nil
	}

	batteryLevel := new(BatteryLevel)
	if err := utils.UnmarshalJSON(data, &batteryLevel, "", true, true); err == nil {
		u.BatteryLevel = batteryLevel
		u.Type = ResponseBodyTypeBatteryLevel
		return nil
	}

	successResponse := new(SuccessResponse)
	if err := utils.UnmarshalJSON(data, &successResponse, "", true, true); err == nil {
		u.SuccessResponse = successResponse
		u.Type = ResponseBodyTypeSuccessResponse
		return nil
	}

	fuelTank := new(FuelTank)
	if err := utils.UnmarshalJSON(data, &fuelTank, "", true, true); err == nil {
		u.FuelTank = fuelTank
		u.Type = ResponseBodyTypeFuelTank
		return nil
	}

	compatibilityResponse := new(CompatibilityResponse)
	if err := utils.UnmarshalJSON(data, &compatibilityResponse, "", true, true); err == nil {
		u.CompatibilityResponse = compatibilityResponse
		u.Type = ResponseBodyTypeCompatibilityResponse
		return nil
	}

	tirePressure := new(TirePressure)
	if err := utils.UnmarshalJSON(data, &tirePressure, "", true, true); err == nil {
		u.TirePressure = tirePressure
		u.Type = ResponseBodyTypeTirePressure
		return nil
	}

	securityRead := new(SecurityRead)
	if err := utils.UnmarshalJSON(data, &securityRead, "", true, true); err == nil {
		u.SecurityRead = securityRead
		u.Type = ResponseBodyTypeSecurityRead
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ResponseBody) MarshalJSON() ([]byte, error) {
	if u.Location != nil {
		return utils.MarshalJSON(u.Location, "", true)
	}

	if u.Odometer != nil {
		return utils.MarshalJSON(u.Odometer, "", true)
	}

	if u.FuelTank != nil {
		return utils.MarshalJSON(u.FuelTank, "", true)
	}

	if u.TirePressure != nil {
		return utils.MarshalJSON(u.TirePressure, "", true)
	}

	if u.EngineOil != nil {
		return utils.MarshalJSON(u.EngineOil, "", true)
	}

	if u.ChargeStatus != nil {
		return utils.MarshalJSON(u.ChargeStatus, "", true)
	}

	if u.ChargeLimit != nil {
		return utils.MarshalJSON(u.ChargeLimit, "", true)
	}

	if u.ChargeTime != nil {
		return utils.MarshalJSON(u.ChargeTime, "", true)
	}

	if u.ChargeVoltage != nil {
		return utils.MarshalJSON(u.ChargeVoltage, "", true)
	}

	if u.BatteryLevel != nil {
		return utils.MarshalJSON(u.BatteryLevel, "", true)
	}

	if u.BatteryCapacity != nil {
		return utils.MarshalJSON(u.BatteryCapacity, "", true)
	}

	if u.CompatibilityResponse != nil {
		return utils.MarshalJSON(u.CompatibilityResponse, "", true)
	}

	if u.VinInfo != nil {
		return utils.MarshalJSON(u.VinInfo, "", true)
	}

	if u.UserInfo != nil {
		return utils.MarshalJSON(u.UserInfo, "", true)
	}

	if u.SuccessResponse != nil {
		return utils.MarshalJSON(u.SuccessResponse, "", true)
	}

	if u.SecurityRead != nil {
		return utils.MarshalJSON(u.SecurityRead, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ResponseCode string

const (
	ResponseCodeTwoHundred  ResponseCode = "200"
	ResponseCodeFiveHundred ResponseCode = "500"
)

func (e ResponseCode) ToPointer() *ResponseCode {
	return &e
}

func (e *ResponseCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "200":
		fallthrough
	case "500":
		*e = ResponseCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseCode: %v", v)
	}
}

type Response struct {
	Body    *ResponseBody `json:"body,omitempty"`
	Code    *ResponseCode `json:"code,omitempty"`
	Headers []Header      `json:"headers,omitempty"`
	Path    *string       `json:"path,omitempty"`
}

func (o *Response) GetBody() *ResponseBody {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Response) GetCode() *ResponseCode {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Response) GetHeaders() []Header {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Response) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}
