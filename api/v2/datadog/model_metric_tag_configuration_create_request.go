/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MetricTagConfigurationCreateRequest Request object that includes the metric that you would like to configure tags for.
type MetricTagConfigurationCreateRequest struct {
	Data MetricTagConfigurationCreateData `json:"data"`
}

// NewMetricTagConfigurationCreateRequest instantiates a new MetricTagConfigurationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTagConfigurationCreateRequest(data MetricTagConfigurationCreateData) *MetricTagConfigurationCreateRequest {
	this := MetricTagConfigurationCreateRequest{}
	this.Data = data
	return &this
}

// NewMetricTagConfigurationCreateRequestWithDefaults instantiates a new MetricTagConfigurationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTagConfigurationCreateRequestWithDefaults() *MetricTagConfigurationCreateRequest {
	this := MetricTagConfigurationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *MetricTagConfigurationCreateRequest) GetData() MetricTagConfigurationCreateData {
	if o == nil {
		var ret MetricTagConfigurationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateRequest) GetDataOk() (*MetricTagConfigurationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MetricTagConfigurationCreateRequest) SetData(v MetricTagConfigurationCreateData) {
	o.Data = v
}

func (o MetricTagConfigurationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMetricTagConfigurationCreateRequest struct {
	value *MetricTagConfigurationCreateRequest
	isSet bool
}

func (v NullableMetricTagConfigurationCreateRequest) Get() *MetricTagConfigurationCreateRequest {
	return v.value
}

func (v *NullableMetricTagConfigurationCreateRequest) Set(val *MetricTagConfigurationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTagConfigurationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTagConfigurationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTagConfigurationCreateRequest(val *MetricTagConfigurationCreateRequest) *NullableMetricTagConfigurationCreateRequest {
	return &NullableMetricTagConfigurationCreateRequest{value: val, isSet: true}
}

func (v NullableMetricTagConfigurationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTagConfigurationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}