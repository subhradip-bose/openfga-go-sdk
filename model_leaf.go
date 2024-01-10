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

// Leaf A leaf node contains either - a set of users (which may be individual users, or usersets   referencing other relations) - a computed node, which is the result of a computed userset   value in the authorization model - a tupleToUserset nodes, containing the result of expanding   a tupleToUserset value in a authorization model.
type Leaf struct {
	Users          *Users                     `json:"users,omitempty"yaml:"users,omitempty"`
	Computed       *Computed                  `json:"computed,omitempty"yaml:"computed,omitempty"`
	TupleToUserset *UsersetTreeTupleToUserset `json:"tupleToUserset,omitempty"yaml:"tupleToUserset,omitempty"`
}

// NewLeaf instantiates a new Leaf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeaf() *Leaf {
	this := Leaf{}
	return &this
}

// NewLeafWithDefaults instantiates a new Leaf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeafWithDefaults() *Leaf {
	this := Leaf{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Leaf) GetUsers() Users {
	if o == nil || o.Users == nil {
		var ret Users
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Leaf) GetUsersOk() (*Users, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Leaf) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given Users and assigns it to the Users field.
func (o *Leaf) SetUsers(v Users) {
	o.Users = &v
}

// GetComputed returns the Computed field value if set, zero value otherwise.
func (o *Leaf) GetComputed() Computed {
	if o == nil || o.Computed == nil {
		var ret Computed
		return ret
	}
	return *o.Computed
}

// GetComputedOk returns a tuple with the Computed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Leaf) GetComputedOk() (*Computed, bool) {
	if o == nil || o.Computed == nil {
		return nil, false
	}
	return o.Computed, true
}

// HasComputed returns a boolean if a field has been set.
func (o *Leaf) HasComputed() bool {
	if o != nil && o.Computed != nil {
		return true
	}

	return false
}

// SetComputed gets a reference to the given Computed and assigns it to the Computed field.
func (o *Leaf) SetComputed(v Computed) {
	o.Computed = &v
}

// GetTupleToUserset returns the TupleToUserset field value if set, zero value otherwise.
func (o *Leaf) GetTupleToUserset() UsersetTreeTupleToUserset {
	if o == nil || o.TupleToUserset == nil {
		var ret UsersetTreeTupleToUserset
		return ret
	}
	return *o.TupleToUserset
}

// GetTupleToUsersetOk returns a tuple with the TupleToUserset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Leaf) GetTupleToUsersetOk() (*UsersetTreeTupleToUserset, bool) {
	if o == nil || o.TupleToUserset == nil {
		return nil, false
	}
	return o.TupleToUserset, true
}

// HasTupleToUserset returns a boolean if a field has been set.
func (o *Leaf) HasTupleToUserset() bool {
	if o != nil && o.TupleToUserset != nil {
		return true
	}

	return false
}

// SetTupleToUserset gets a reference to the given UsersetTreeTupleToUserset and assigns it to the TupleToUserset field.
func (o *Leaf) SetTupleToUserset(v UsersetTreeTupleToUserset) {
	o.TupleToUserset = &v
}

func (o Leaf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Computed != nil {
		toSerialize["computed"] = o.Computed
	}
	if o.TupleToUserset != nil {
		toSerialize["tupleToUserset"] = o.TupleToUserset
	}
	return json.Marshal(toSerialize)
}

type NullableLeaf struct {
	value *Leaf
	isSet bool
}

func (v NullableLeaf) Get() *Leaf {
	return v.value
}

func (v *NullableLeaf) Set(val *Leaf) {
	v.value = val
	v.isSet = true
}

func (v NullableLeaf) IsSet() bool {
	return v.isSet
}

func (v *NullableLeaf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeaf(val *Leaf) *NullableLeaf {
	return &NullableLeaf{value: val, isSet: true}
}

func (v NullableLeaf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeaf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
