/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/subhradip-bose/openfga-go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
	"time"
)

// Tuple struct for Tuple
type Tuple struct {
	Key       TupleKey  `json:"key"yaml:"key"`
	Timestamp time.Time `json:"timestamp"yaml:"timestamp"`
}

// NewTuple instantiates a new Tuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTuple(key TupleKey, timestamp time.Time) *Tuple {
	this := Tuple{}
	this.Key = key
	this.Timestamp = timestamp
	return &this
}

// NewTupleWithDefaults instantiates a new Tuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleWithDefaults() *Tuple {
	this := Tuple{}
	return &this
}

// GetKey returns the Key field value
func (o *Tuple) GetKey() TupleKey {
	if o == nil {
		var ret TupleKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Tuple) GetKeyOk() (*TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Tuple) SetKey(v TupleKey) {
	o.Key = v
}

// GetTimestamp returns the Timestamp field value
func (o *Tuple) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Tuple) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Tuple) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o Tuple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["timestamp"] = o.Timestamp
	return json.Marshal(toSerialize)
}

type NullableTuple struct {
	value *Tuple
	isSet bool
}

func (v NullableTuple) Get() *Tuple {
	return v.value
}

func (v *NullableTuple) Set(val *Tuple) {
	v.value = val
	v.isSet = true
}

func (v NullableTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTuple(val *Tuple) *NullableTuple {
	return &NullableTuple{value: val, isSet: true}
}

func (v NullableTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
