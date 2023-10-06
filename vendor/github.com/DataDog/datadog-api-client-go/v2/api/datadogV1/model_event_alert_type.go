// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"
)

// EventAlertType If an alert event is enabled, set its type.
// For example, `error`, `warning`, `info`, `success`, `user_update`,
// `recommendation`, and `snapshot`.
type EventAlertType string

// List of EventAlertType.
const (
	EVENTALERTTYPE_ERROR          EventAlertType = "error"
	EVENTALERTTYPE_WARNING        EventAlertType = "warning"
	EVENTALERTTYPE_INFO           EventAlertType = "info"
	EVENTALERTTYPE_SUCCESS        EventAlertType = "success"
	EVENTALERTTYPE_USER_UPDATE    EventAlertType = "user_update"
	EVENTALERTTYPE_RECOMMENDATION EventAlertType = "recommendation"
	EVENTALERTTYPE_SNAPSHOT       EventAlertType = "snapshot"
)

var allowedEventAlertTypeEnumValues = []EventAlertType{
	EVENTALERTTYPE_ERROR,
	EVENTALERTTYPE_WARNING,
	EVENTALERTTYPE_INFO,
	EVENTALERTTYPE_SUCCESS,
	EVENTALERTTYPE_USER_UPDATE,
	EVENTALERTTYPE_RECOMMENDATION,
	EVENTALERTTYPE_SNAPSHOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventAlertType) GetAllowedValues() []EventAlertType {
	return allowedEventAlertTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventAlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventAlertType(value)
	return nil
}

// NewEventAlertTypeFromValue returns a pointer to a valid EventAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventAlertTypeFromValue(v string) (*EventAlertType, error) {
	ev := EventAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventAlertType: valid values are %v", v, allowedEventAlertTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventAlertType) IsValid() bool {
	for _, existing := range allowedEventAlertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventAlertType value.
func (v EventAlertType) Ptr() *EventAlertType {
	return &v
}