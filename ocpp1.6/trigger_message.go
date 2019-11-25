package ocpp16

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Trigger Message (CS -> CP) --------------------

// Status reported in TriggerMessageConfirmation.
type TriggerMessageStatus string

// Type of request to be triggered in a TriggerMessageRequest
type MessageTrigger string

const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)

func isValidTriggerMessageStatus(fl validator.FieldLevel) bool {
	status := TriggerMessageStatus(fl.Field().String())
	switch status {
	case TriggerMessageStatusAccepted, TriggerMessageStatusRejected, TriggerMessageStatusNotImplemented:
		return true
	default:
		return false
	}
}

func isValidMessageTrigger(fl validator.FieldLevel) bool {
	trigger := MessageTrigger(fl.Field().String())
	switch trigger {
	case BootNotificationFeatureName, DiagnosticsStatusNotificationFeatureName, FirmwareStatusNotificationFeatureName, HeartbeatFeatureName, MeterValuesFeatureName, StatusNotificationFeatureName:
		return true
	default:
		return false
	}
}

// The field definition of the TriggerMessage request payload sent by the Central System to the Charge Point.
type TriggerMessageRequest struct {
	RequestedMessage MessageTrigger `json:"requestedMessage" validate:"required,messageTrigger"`
	ConnectorId      int            `json:"connectorId,omitempty" validate:"omitempty,gt=0"`
}

// This field definition of the TriggerMessage confirmation payload, sent by the Charge Point to the Central System in response to a TriggerMessageRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type TriggerMessageConfirmation struct {
	Status TriggerMessageStatus `json:"status" validate:"required,triggerMessageStatus"`
}

// The TriggerMessageRequest makes it possible for the Central System, to request the Charge Point, to send Charge Point-initiated messages.
// In the request the Central System indicates which message it wishes to receive.
// The Charge Point SHALL first send the TriggerMessage response, before sending the requested message.
// In the TriggerMessageConfirmation the Charge Point SHALL indicate whether it will send it or not, by returning ACCEPTED or REJECTED.
// It is up to the Charge Point if it accepts or rejects the request to send.
// If the requested message is unknown or not implemented the Charge Point SHALL return NOT_IMPLEMENTED.
type TriggerMessageFeature struct{}

func (f TriggerMessageFeature) GetFeatureName() string {
	return TriggerMessageFeatureName
}

func (f TriggerMessageFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(TriggerMessageRequest{})
}

func (f TriggerMessageFeature) GetConfirmationType() reflect.Type {
	return reflect.TypeOf(TriggerMessageConfirmation{})
}

func (r TriggerMessageRequest) GetFeatureName() string {
	return TriggerMessageFeatureName
}

func (c TriggerMessageConfirmation) GetFeatureName() string {
	return TriggerMessageFeatureName
}

// Creates a new TriggerMessageRequest, containing all required fields. Optional fields may be set afterwards.
func NewTriggerMessageRequest(requestedMessage MessageTrigger) *TriggerMessageRequest {
	return &TriggerMessageRequest{RequestedMessage: requestedMessage}
}

// Creates a new TriggerMessageConfirmation, containing all required fields. There are no optional fields for this message.
func NewTriggerMessageConfirmation(status TriggerMessageStatus) *TriggerMessageConfirmation {
	return &TriggerMessageConfirmation{Status: status}
}

func init() {
	_ = Validate.RegisterValidation("triggerMessageStatus", isValidTriggerMessageStatus)
	_ = Validate.RegisterValidation("messageTrigger", isValidMessageTrigger)
}