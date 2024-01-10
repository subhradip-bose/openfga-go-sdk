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
)

// Assertion struct for Assertion
type Assertion struct {
	TupleKey    AssertionTupleKey `json:"tuple_key"yaml:"tuple_key"`
	Expectation bool              `json:"expectation"yaml:"expectation"`
}

// NewAssertion instantiates a new Assertion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertion(tupleKey AssertionTupleKey, expectation bool) *Assertion {
	this := Assertion{}
	this.TupleKey = tupleKey
	this.Expectation = expectation
	return &this
}

// NewAssertionWithDefaults instantiates a new Assertion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertionWithDefaults() *Assertion {
	this := Assertion{}
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *Assertion) GetTupleKey() AssertionTupleKey {
	if o == nil {
		var ret AssertionTupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *Assertion) GetTupleKeyOk() (*AssertionTupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *Assertion) SetTupleKey(v AssertionTupleKey) {
	o.TupleKey = v
}

// GetExpectation returns the Expectation field value
func (o *Assertion) GetExpectation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Expectation
}

// GetExpectationOk returns a tuple with the Expectation field value
// and a boolean to check if the value has been set.
func (o *Assertion) GetExpectationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expectation, true
}

// SetExpectation sets field value
func (o *Assertion) SetExpectation(v bool) {
	o.Expectation = v
}

func (o Assertion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	toSerialize["expectation"] = o.Expectation
	return json.Marshal(toSerialize)
}

type NullableAssertion struct {
	value *Assertion
	isSet bool
}

func (v NullableAssertion) Get() *Assertion {
	return v.value
}

func (v *NullableAssertion) Set(val *Assertion) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertion) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertion(val *Assertion) *NullableAssertion {
	return &NullableAssertion{value: val, isSet: true}
}

func (v NullableAssertion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
