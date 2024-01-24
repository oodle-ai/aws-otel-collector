/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesClusterEntities struct for KubernetesClusterEntities
type KubernetesClusterEntities struct {
	Nodepools *KubernetesNodePools `json:"nodepools,omitempty"`
}

// NewKubernetesClusterEntities instantiates a new KubernetesClusterEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterEntities() *KubernetesClusterEntities {
	this := KubernetesClusterEntities{}

	return &this
}

// NewKubernetesClusterEntitiesWithDefaults instantiates a new KubernetesClusterEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterEntitiesWithDefaults() *KubernetesClusterEntities {
	this := KubernetesClusterEntities{}
	return &this
}

// GetNodepools returns the Nodepools field value
// If the value is explicit nil, nil is returned
func (o *KubernetesClusterEntities) GetNodepools() *KubernetesNodePools {
	if o == nil {
		return nil
	}

	return o.Nodepools

}

// GetNodepoolsOk returns a tuple with the Nodepools field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterEntities) GetNodepoolsOk() (*KubernetesNodePools, bool) {
	if o == nil {
		return nil, false
	}

	return o.Nodepools, true
}

// SetNodepools sets field value
func (o *KubernetesClusterEntities) SetNodepools(v KubernetesNodePools) {

	o.Nodepools = &v

}

// HasNodepools returns a boolean if a field has been set.
func (o *KubernetesClusterEntities) HasNodepools() bool {
	if o != nil && o.Nodepools != nil {
		return true
	}

	return false
}

func (o KubernetesClusterEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nodepools != nil {
		toSerialize["nodepools"] = o.Nodepools
	}

	return json.Marshal(toSerialize)
}

type NullableKubernetesClusterEntities struct {
	value *KubernetesClusterEntities
	isSet bool
}

func (v NullableKubernetesClusterEntities) Get() *KubernetesClusterEntities {
	return v.value
}

func (v *NullableKubernetesClusterEntities) Set(val *KubernetesClusterEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterEntities(val *KubernetesClusterEntities) *NullableKubernetesClusterEntities {
	return &NullableKubernetesClusterEntities{value: val, isSet: true}
}

func (v NullableKubernetesClusterEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
