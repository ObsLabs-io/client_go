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

// checks if the V1AcceptInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AcceptInvitationRequest{}

// V1AcceptInvitationRequest struct for V1AcceptInvitationRequest
type V1AcceptInvitationRequest struct {
	SecretKey string `json:"secret_key"`
	AdditionalProperties map[string]interface{}
}

type _V1AcceptInvitationRequest V1AcceptInvitationRequest

// NewV1AcceptInvitationRequest instantiates a new V1AcceptInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AcceptInvitationRequest(secretKey string) *V1AcceptInvitationRequest {
	this := V1AcceptInvitationRequest{}
	this.SecretKey = secretKey
	return &this
}

// NewV1AcceptInvitationRequestWithDefaults instantiates a new V1AcceptInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AcceptInvitationRequestWithDefaults() *V1AcceptInvitationRequest {
	this := V1AcceptInvitationRequest{}
	return &this
}

// GetSecretKey returns the SecretKey field value
func (o *V1AcceptInvitationRequest) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *V1AcceptInvitationRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *V1AcceptInvitationRequest) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o V1AcceptInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AcceptInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret_key"] = o.SecretKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1AcceptInvitationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret_key",
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

	varV1AcceptInvitationRequest := _V1AcceptInvitationRequest{}

	err = json.Unmarshal(data, &varV1AcceptInvitationRequest)

	if err != nil {
		return err
	}

	*o = V1AcceptInvitationRequest(varV1AcceptInvitationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "secret_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1AcceptInvitationRequest struct {
	value *V1AcceptInvitationRequest
	isSet bool
}

func (v NullableV1AcceptInvitationRequest) Get() *V1AcceptInvitationRequest {
	return v.value
}

func (v *NullableV1AcceptInvitationRequest) Set(val *V1AcceptInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AcceptInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AcceptInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AcceptInvitationRequest(val *V1AcceptInvitationRequest) *NullableV1AcceptInvitationRequest {
	return &NullableV1AcceptInvitationRequest{value: val, isSet: true}
}

func (v NullableV1AcceptInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AcceptInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


