/*
 * test API
 *
 * retrieve temperature & humidity and compute tests
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Sensorefficiency struct for Sensorefficiency
type Sensorefficiency struct {
	Value string `json:"value"`
}

// NewSensorefficiency instantiates a new Sensorefficiency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorefficiency(value string) *Sensorefficiency {
	this := Sensorefficiency{}
	this.Value = value
	return &this
}

// NewSensorefficiencyWithDefaults instantiates a new Sensorefficiency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorefficiencyWithDefaults() *Sensorefficiency {
	this := Sensorefficiency{}
	return &this
}

// GetValue returns the Value field value
func (o *Sensorefficiency) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Sensorefficiency) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Sensorefficiency) SetValue(v string) {
	o.Value = v
}

func (o Sensorefficiency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSensorefficiency struct {
	value *Sensorefficiency
	isSet bool
}

func (v NullableSensorefficiency) Get() *Sensorefficiency {
	return v.value
}

func (v *NullableSensorefficiency) Set(val *Sensorefficiency) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorefficiency) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorefficiency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorefficiency(val *Sensorefficiency) *NullableSensorefficiency {
	return &NullableSensorefficiency{value: val, isSet: true}
}

func (v NullableSensorefficiency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorefficiency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


