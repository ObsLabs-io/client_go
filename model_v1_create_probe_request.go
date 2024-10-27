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

// checks if the V1CreateProbeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateProbeRequest{}

// V1CreateProbeRequest struct for V1CreateProbeRequest
type V1CreateProbeRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Required for pull probes. Autogenerated for push probes.
	Url                  *string             `json:"url,omitempty"`
	Schedule             *ProbeScheduleModel `json:"schedule,omitempty"`
	Channels             []string            `json:"channels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1CreateProbeRequest V1CreateProbeRequest

// NewV1CreateProbeRequest instantiates a new V1CreateProbeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateProbeRequest(name string, type_ string) *V1CreateProbeRequest {
	this := V1CreateProbeRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewV1CreateProbeRequestWithDefaults instantiates a new V1CreateProbeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateProbeRequestWithDefaults() *V1CreateProbeRequest {
	this := V1CreateProbeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *V1CreateProbeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1CreateProbeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1CreateProbeRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *V1CreateProbeRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1CreateProbeRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1CreateProbeRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V1CreateProbeRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateProbeRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V1CreateProbeRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V1CreateProbeRequest) SetUrl(v string) {
	o.Url = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *V1CreateProbeRequest) GetSchedule() ProbeScheduleModel {
	if o == nil || IsNil(o.Schedule) {
		var ret ProbeScheduleModel
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateProbeRequest) GetScheduleOk() (*ProbeScheduleModel, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *V1CreateProbeRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ProbeScheduleModel and assigns it to the Schedule field.
func (o *V1CreateProbeRequest) SetSchedule(v ProbeScheduleModel) {
	o.Schedule = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *V1CreateProbeRequest) GetChannels() []string {
	if o == nil || IsNil(o.Channels) {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateProbeRequest) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *V1CreateProbeRequest) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *V1CreateProbeRequest) SetChannels(v []string) {
	o.Channels = v
}

func (o V1CreateProbeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateProbeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1CreateProbeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1CreateProbeRequest := _V1CreateProbeRequest{}

	err = json.Unmarshal(data, &varV1CreateProbeRequest)

	if err != nil {
		return err
	}

	*o = V1CreateProbeRequest(varV1CreateProbeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "channels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1CreateProbeRequest struct {
	value *V1CreateProbeRequest
	isSet bool
}

func (v NullableV1CreateProbeRequest) Get() *V1CreateProbeRequest {
	return v.value
}

func (v *NullableV1CreateProbeRequest) Set(val *V1CreateProbeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateProbeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateProbeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateProbeRequest(val *V1CreateProbeRequest) *NullableV1CreateProbeRequest {
	return &NullableV1CreateProbeRequest{value: val, isSet: true}
}

func (v NullableV1CreateProbeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateProbeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
