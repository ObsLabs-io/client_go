/*
ObsLabs API

# Authentication  ObsLabs uses basic auth to authenticate the API. You can create API keys in the account settings. Use your API key as the basic auth password. The username should be left blank (notice the colon sign before api-key that must be included). All requests must be made over https.  Example usage: ```bash curl -u :<YOUR API KEY> https://api.obslabs.io/v1/users/me ```  # Errors  The API returns a structured error response in case of failure. Below is the format of the error response object:  ```json {   \"error\": {     \"status\": 400,     \"code\": \"VALIDATION\",     \"message\": \"Validation errors occurred.\",     \"details\": [       {         \"field\": \"email\",         \"issue\": \"The email address is not in a valid format.\"       },       {         \"field\": \"password\",         \"issue\": \"The password must be at least 8 characters long.\"       }     ]   } } 

API version: 1.0
Contact: contact@obslabs.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client_go

import (
	"encoding/json"
	"fmt"
)

// checks if the ProbeChannelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProbeChannelModel{}

// ProbeChannelModel struct for ProbeChannelModel
type ProbeChannelModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ProbeChannelModel ProbeChannelModel

// NewProbeChannelModel instantiates a new ProbeChannelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProbeChannelModel(id string, name string) *ProbeChannelModel {
	this := ProbeChannelModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProbeChannelModelWithDefaults instantiates a new ProbeChannelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProbeChannelModelWithDefaults() *ProbeChannelModel {
	this := ProbeChannelModel{}
	return &this
}

// GetId returns the Id field value
func (o *ProbeChannelModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProbeChannelModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProbeChannelModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProbeChannelModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProbeChannelModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProbeChannelModel) SetName(v string) {
	o.Name = v
}

func (o ProbeChannelModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProbeChannelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProbeChannelModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProbeChannelModel := _ProbeChannelModel{}

	err = json.Unmarshal(data, &varProbeChannelModel)

	if err != nil {
		return err
	}

	*o = ProbeChannelModel(varProbeChannelModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProbeChannelModel struct {
	value *ProbeChannelModel
	isSet bool
}

func (v NullableProbeChannelModel) Get() *ProbeChannelModel {
	return v.value
}

func (v *NullableProbeChannelModel) Set(val *ProbeChannelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProbeChannelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProbeChannelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbeChannelModel(val *ProbeChannelModel) *NullableProbeChannelModel {
	return &NullableProbeChannelModel{value: val, isSet: true}
}

func (v NullableProbeChannelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbeChannelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


