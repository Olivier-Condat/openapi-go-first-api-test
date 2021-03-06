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

// Sensordata struct for Sensordata
type Sensordata struct {
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
}

// NewSensordata instantiates a new Sensordata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensordata(temperature float32, humidity float32) *Sensordata {
	this := Sensordata{}
	this.Temperature = temperature
	this.Humidity = humidity
	return &this
}

// NewSensordataWithDefaults instantiates a new Sensordata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensordataWithDefaults() *Sensordata {
	this := Sensordata{}
	return &this
}

// GetTemperature returns the Temperature field value
func (o *Sensordata) GetTemperature() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value
// and a boolean to check if the value has been set.
func (o *Sensordata) GetTemperatureOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Temperature, true
}

// SetTemperature sets field value
func (o *Sensordata) SetTemperature(v float32) {
	o.Temperature = v
}

// GetHumidity returns the Humidity field value
func (o *Sensordata) GetHumidity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value
// and a boolean to check if the value has been set.
func (o *Sensordata) GetHumidityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Humidity, true
}

// SetHumidity sets field value
func (o *Sensordata) SetHumidity(v float32) {
	o.Humidity = v
}

func (o Sensordata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["temperature"] = o.Temperature
	}
	if true {
		toSerialize["humidity"] = o.Humidity
	}
	return json.Marshal(toSerialize)
}

type NullableSensordata struct {
	value *Sensordata
	isSet bool
}

func (v NullableSensordata) Get() *Sensordata {
	return v.value
}

func (v *NullableSensordata) Set(val *Sensordata) {
	v.value = val
	v.isSet = true
}

func (v NullableSensordata) IsSet() bool {
	return v.isSet
}

func (v *NullableSensordata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensordata(val *Sensordata) *NullableSensordata {
	return &NullableSensordata{value: val, isSet: true}
}

func (v NullableSensordata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensordata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


