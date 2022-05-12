/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// ReadChangesResponse struct for ReadChangesResponse
type ReadChangesResponse struct {
	Changes           *[]TupleChange `json:"changes,omitempty"`
	ContinuationToken *string        `json:"continuation_token,omitempty"`
}

// NewReadChangesResponse instantiates a new ReadChangesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadChangesResponse() *ReadChangesResponse {
	this := ReadChangesResponse{}
	return &this
}

// NewReadChangesResponseWithDefaults instantiates a new ReadChangesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadChangesResponseWithDefaults() *ReadChangesResponse {
	this := ReadChangesResponse{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *ReadChangesResponse) GetChanges() []TupleChange {
	if o == nil || o.Changes == nil {
		var ret []TupleChange
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadChangesResponse) GetChangesOk() (*[]TupleChange, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *ReadChangesResponse) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []TupleChange and assigns it to the Changes field.
func (o *ReadChangesResponse) SetChanges(v []TupleChange) {
	o.Changes = &v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *ReadChangesResponse) GetContinuationToken() string {
	if o == nil || o.ContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadChangesResponse) GetContinuationTokenOk() (*string, bool) {
	if o == nil || o.ContinuationToken == nil {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *ReadChangesResponse) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken != nil {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *ReadChangesResponse) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o ReadChangesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.ContinuationToken != nil {
		toSerialize["continuation_token"] = o.ContinuationToken
	}
	return json.Marshal(toSerialize)
}

type NullableReadChangesResponse struct {
	value *ReadChangesResponse
	isSet bool
}

func (v NullableReadChangesResponse) Get() *ReadChangesResponse {
	return v.value
}

func (v *NullableReadChangesResponse) Set(val *ReadChangesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadChangesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadChangesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadChangesResponse(val *ReadChangesResponse) *NullableReadChangesResponse {
	return &NullableReadChangesResponse{value: val, isSet: true}
}

func (v NullableReadChangesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadChangesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
