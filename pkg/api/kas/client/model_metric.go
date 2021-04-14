/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// Metric struct for Metric
type Metric struct {
	Metric *map[string]string `json:"metric,omitempty"`
	Values *[]Values          `json:"values,omitempty"`
}

// NewMetric instantiates a new Metric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetric() *Metric {
	this := Metric{}
	return &this
}

// NewMetricWithDefaults instantiates a new Metric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricWithDefaults() *Metric {
	this := Metric{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *Metric) GetMetric() map[string]string {
	if o == nil || o.Metric == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetMetricOk() (*map[string]string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *Metric) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]string and assigns it to the Metric field.
func (o *Metric) SetMetric(v map[string]string) {
	o.Metric = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Metric) GetValues() []Values {
	if o == nil || o.Values == nil {
		var ret []Values
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metric) GetValuesOk() (*[]Values, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Metric) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []Values and assigns it to the Values field.
func (o *Metric) SetValues(v []Values) {
	o.Values = &v
}

func (o Metric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableMetric struct {
	value *Metric
	isSet bool
}

func (v NullableMetric) Get() *Metric {
	return v.value
}

func (v *NullableMetric) Set(val *Metric) {
	v.value = val
	v.isSet = true
}

func (v NullableMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetric(val *Metric) *NullableMetric {
	return &NullableMetric{value: val, isSet: true}
}

func (v NullableMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}