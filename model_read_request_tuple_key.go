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

// ReadRequestTupleKey struct for ReadRequestTupleKey
type ReadRequestTupleKey struct {
	User     *string `json:"user,omitempty"yaml:"user,omitempty"`
	Relation *string `json:"relation,omitempty"yaml:"relation,omitempty"`
	Object   *string `json:"object,omitempty"yaml:"object,omitempty"`
}

// NewReadRequestTupleKey instantiates a new ReadRequestTupleKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRequestTupleKey() *ReadRequestTupleKey {
	this := ReadRequestTupleKey{}
	return &this
}

// NewReadRequestTupleKeyWithDefaults instantiates a new ReadRequestTupleKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRequestTupleKeyWithDefaults() *ReadRequestTupleKey {
	this := ReadRequestTupleKey{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ReadRequestTupleKey) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequestTupleKey) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ReadRequestTupleKey) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ReadRequestTupleKey) SetUser(v string) {
	o.User = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *ReadRequestTupleKey) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequestTupleKey) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *ReadRequestTupleKey) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *ReadRequestTupleKey) SetRelation(v string) {
	o.Relation = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ReadRequestTupleKey) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRequestTupleKey) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ReadRequestTupleKey) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ReadRequestTupleKey) SetObject(v string) {
	o.Object = &v
}

func (o ReadRequestTupleKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableReadRequestTupleKey struct {
	value *ReadRequestTupleKey
	isSet bool
}

func (v NullableReadRequestTupleKey) Get() *ReadRequestTupleKey {
	return v.value
}

func (v *NullableReadRequestTupleKey) Set(val *ReadRequestTupleKey) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRequestTupleKey) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRequestTupleKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRequestTupleKey(val *ReadRequestTupleKey) *NullableReadRequestTupleKey {
	return &NullableReadRequestTupleKey{value: val, isSet: true}
}

func (v NullableReadRequestTupleKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRequestTupleKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
