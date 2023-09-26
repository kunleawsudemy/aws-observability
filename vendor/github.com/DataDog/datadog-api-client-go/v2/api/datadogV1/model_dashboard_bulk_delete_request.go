// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardBulkDeleteRequest Dashboard bulk delete request body.
type DashboardBulkDeleteRequest struct {
	// List of dashboard bulk action request data objects.
	Data []DashboardBulkActionData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardBulkDeleteRequest instantiates a new DashboardBulkDeleteRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardBulkDeleteRequest(data []DashboardBulkActionData) *DashboardBulkDeleteRequest {
	this := DashboardBulkDeleteRequest{}
	this.Data = data
	return &this
}

// NewDashboardBulkDeleteRequestWithDefaults instantiates a new DashboardBulkDeleteRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardBulkDeleteRequestWithDefaults() *DashboardBulkDeleteRequest {
	this := DashboardBulkDeleteRequest{}
	return &this
}

// GetData returns the Data field value.
func (o *DashboardBulkDeleteRequest) GetData() []DashboardBulkActionData {
	if o == nil {
		var ret []DashboardBulkActionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DashboardBulkDeleteRequest) GetDataOk() (*[]DashboardBulkActionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *DashboardBulkDeleteRequest) SetData(v []DashboardBulkActionData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardBulkDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]DashboardBulkActionData `json:"data"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
