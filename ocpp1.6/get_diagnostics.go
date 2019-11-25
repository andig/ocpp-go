package ocpp16

import (
	"reflect"
)

// -------------------- Get Diagnostics (CS -> CP) --------------------

// The field definition of the GetDiagnostics request payload sent by the Central System to the Charge Point.
type GetDiagnosticsRequest struct {
	Location      string    `json:"location" validate:"required,uri"`
	Retries       int       `json:"retries" validate:"gte=0"`       //TODO: pointer?
	RetryInterval int       `json:"retryInterval" validate:"gte=0"` //TODO: pointer?
	StartTime     *DateTime `json:"startTime,omitempty"`
	EndTime       *DateTime `json:"endTime,omitempty"`
}

// This field definition of the GetDiagnostics confirmation payload, sent by the Charge Point to the Central System in response to a GetDiagnosticsRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type GetDiagnosticsConfirmation struct {
	FileName string `json:"fileName,omitempty" validate:"max=255"`
}

// Central System can request a Charge Point for diagnostic information.
// The Central System SHALL send a GetDiagnosticsRequest for getting diagnostic information of a Charge Point with a location where the
// Charge Point MUST upload its diagnostic data to and optionally a begin and end time for the requested diagnostic information.
// then Charge Point SHALL 48
// The Charge Point SHALL respond with a GetDiagnosticsConfirmation stating the name of the file containing the diagnostic information that will be uploaded.
// Charge Point SHALL upload a single file. Format of the diagnostics file is not prescribed.
// If no diagnostics file is available, then GetDiagnosticsConfirmation SHALL NOT contain a file name.
type GetDiagnosticsFeature struct{}

func (f GetDiagnosticsFeature) GetFeatureName() string {
	return GetDiagnosticsFeatureName
}

func (f GetDiagnosticsFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(GetDiagnosticsRequest{})
}

func (f GetDiagnosticsFeature) GetConfirmationType() reflect.Type {
	return reflect.TypeOf(GetDiagnosticsConfirmation{})
}

func (r GetDiagnosticsRequest) GetFeatureName() string {
	return GetDiagnosticsFeatureName
}

func (c GetDiagnosticsConfirmation) GetFeatureName() string {
	return GetDiagnosticsFeatureName
}

// Creates a new GetDiagnosticsRequest, which doesn't contain any required or optional fields.
func NewGetDiagnosticsRequest(location string) *GetDiagnosticsRequest {
	return &GetDiagnosticsRequest{Location: location}
}

// Creates a new GetDiagnosticsConfirmation, containing all required fields. There are no optional fields for this message.
func NewGetDiagnosticsConfirmation() *GetDiagnosticsConfirmation {
	return &GetDiagnosticsConfirmation{}
}