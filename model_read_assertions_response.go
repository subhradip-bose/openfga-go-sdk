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

// ReadAssertionsResponse struct for ReadAssertionsResponse
type ReadAssertionsResponse struct {
	AuthorizationModelId string       `json:"authorization_model_id"yaml:"authorization_model_id"`
	Assertions           *[]Assertion `json:"assertions,omitempty"yaml:"assertions,omitempty"`
}

// NewReadAssertionsResponse instantiates a new ReadAssertionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAssertionsResponse(authorizationModelId string) *ReadAssertionsResponse {
	this := ReadAssertionsResponse{}
	this.AuthorizationModelId = authorizationModelId
	return &this
}

// NewReadAssertionsResponseWithDefaults instantiates a new ReadAssertionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAssertionsResponseWithDefaults() *ReadAssertionsResponse {
	this := ReadAssertionsResponse{}
	return &this
}

// GetAuthorizationModelId returns the AuthorizationModelId field value
func (o *ReadAssertionsResponse) GetAuthorizationModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value
// and a boolean to check if the value has been set.
func (o *ReadAssertionsResponse) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationModelId, true
}

// SetAuthorizationModelId sets field value
func (o *ReadAssertionsResponse) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = v
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *ReadAssertionsResponse) GetAssertions() []Assertion {
	if o == nil || o.Assertions == nil {
		var ret []Assertion
		return ret
	}
	return *o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAssertionsResponse) GetAssertionsOk() (*[]Assertion, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *ReadAssertionsResponse) HasAssertions() bool {
	if o != nil && o.Assertions != nil {
		return true
	}

	return false
}

// SetAssertions gets a reference to the given []Assertion and assigns it to the Assertions field.
func (o *ReadAssertionsResponse) SetAssertions(v []Assertion) {
	o.Assertions = &v
}

func (o ReadAssertionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorization_model_id"] = o.AuthorizationModelId
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	return json.Marshal(toSerialize)
}

type NullableReadAssertionsResponse struct {
	value *ReadAssertionsResponse
	isSet bool
}

func (v NullableReadAssertionsResponse) Get() *ReadAssertionsResponse {
	return v.value
}

func (v *NullableReadAssertionsResponse) Set(val *ReadAssertionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAssertionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAssertionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAssertionsResponse(val *ReadAssertionsResponse) *NullableReadAssertionsResponse {
	return &NullableReadAssertionsResponse{value: val, isSet: true}
}

func (v NullableReadAssertionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAssertionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
