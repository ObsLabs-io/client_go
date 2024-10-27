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

// checks if the V1SendInvitation201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SendInvitation201Response{}

// V1SendInvitation201Response struct for V1SendInvitation201Response
type V1SendInvitation201Response struct {
	Invitation           InvitationModel `json:"invitation"`
	AdditionalProperties map[string]interface{}
}

type _V1SendInvitation201Response V1SendInvitation201Response

// NewV1SendInvitation201Response instantiates a new V1SendInvitation201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SendInvitation201Response(invitation InvitationModel) *V1SendInvitation201Response {
	this := V1SendInvitation201Response{}
	this.Invitation = invitation
	return &this
}

// NewV1SendInvitation201ResponseWithDefaults instantiates a new V1SendInvitation201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SendInvitation201ResponseWithDefaults() *V1SendInvitation201Response {
	this := V1SendInvitation201Response{}
	return &this
}

// GetInvitation returns the Invitation field value
func (o *V1SendInvitation201Response) GetInvitation() InvitationModel {
	if o == nil {
		var ret InvitationModel
		return ret
	}

	return o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value
// and a boolean to check if the value has been set.
func (o *V1SendInvitation201Response) GetInvitationOk() (*InvitationModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invitation, true
}

// SetInvitation sets field value
func (o *V1SendInvitation201Response) SetInvitation(v InvitationModel) {
	o.Invitation = v
}

func (o V1SendInvitation201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SendInvitation201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invitation"] = o.Invitation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1SendInvitation201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invitation",
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

	varV1SendInvitation201Response := _V1SendInvitation201Response{}

	err = json.Unmarshal(data, &varV1SendInvitation201Response)

	if err != nil {
		return err
	}

	*o = V1SendInvitation201Response(varV1SendInvitation201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invitation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1SendInvitation201Response struct {
	value *V1SendInvitation201Response
	isSet bool
}

func (v NullableV1SendInvitation201Response) Get() *V1SendInvitation201Response {
	return v.value
}

func (v *NullableV1SendInvitation201Response) Set(val *V1SendInvitation201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SendInvitation201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SendInvitation201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SendInvitation201Response(val *V1SendInvitation201Response) *NullableV1SendInvitation201Response {
	return &NullableV1SendInvitation201Response{value: val, isSet: true}
}

func (v NullableV1SendInvitation201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SendInvitation201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
