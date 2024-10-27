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

// checks if the V1CreateChannelRequestEmailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateChannelRequestEmailsInner{}

// V1CreateChannelRequestEmailsInner struct for V1CreateChannelRequestEmailsInner
type V1CreateChannelRequestEmailsInner struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _V1CreateChannelRequestEmailsInner V1CreateChannelRequestEmailsInner

// NewV1CreateChannelRequestEmailsInner instantiates a new V1CreateChannelRequestEmailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateChannelRequestEmailsInner(name string, email string) *V1CreateChannelRequestEmailsInner {
	this := V1CreateChannelRequestEmailsInner{}
	this.Name = name
	this.Email = email
	return &this
}

// NewV1CreateChannelRequestEmailsInnerWithDefaults instantiates a new V1CreateChannelRequestEmailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateChannelRequestEmailsInnerWithDefaults() *V1CreateChannelRequestEmailsInner {
	this := V1CreateChannelRequestEmailsInner{}
	return &this
}

// GetName returns the Name field value
func (o *V1CreateChannelRequestEmailsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1CreateChannelRequestEmailsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1CreateChannelRequestEmailsInner) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *V1CreateChannelRequestEmailsInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *V1CreateChannelRequestEmailsInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *V1CreateChannelRequestEmailsInner) SetEmail(v string) {
	o.Email = v
}

func (o V1CreateChannelRequestEmailsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateChannelRequestEmailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1CreateChannelRequestEmailsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
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

	varV1CreateChannelRequestEmailsInner := _V1CreateChannelRequestEmailsInner{}

	err = json.Unmarshal(data, &varV1CreateChannelRequestEmailsInner)

	if err != nil {
		return err
	}

	*o = V1CreateChannelRequestEmailsInner(varV1CreateChannelRequestEmailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1CreateChannelRequestEmailsInner struct {
	value *V1CreateChannelRequestEmailsInner
	isSet bool
}

func (v NullableV1CreateChannelRequestEmailsInner) Get() *V1CreateChannelRequestEmailsInner {
	return v.value
}

func (v *NullableV1CreateChannelRequestEmailsInner) Set(val *V1CreateChannelRequestEmailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateChannelRequestEmailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateChannelRequestEmailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateChannelRequestEmailsInner(val *V1CreateChannelRequestEmailsInner) *NullableV1CreateChannelRequestEmailsInner {
	return &NullableV1CreateChannelRequestEmailsInner{value: val, isSet: true}
}

func (v NullableV1CreateChannelRequestEmailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateChannelRequestEmailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
