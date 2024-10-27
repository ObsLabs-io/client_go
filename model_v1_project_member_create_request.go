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

// checks if the V1ProjectMemberCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ProjectMemberCreateRequest{}

// V1ProjectMemberCreateRequest struct for V1ProjectMemberCreateRequest
type V1ProjectMemberCreateRequest struct {
	UserId               string `json:"user_id"`
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _V1ProjectMemberCreateRequest V1ProjectMemberCreateRequest

// NewV1ProjectMemberCreateRequest instantiates a new V1ProjectMemberCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ProjectMemberCreateRequest(userId string, role string) *V1ProjectMemberCreateRequest {
	this := V1ProjectMemberCreateRequest{}
	this.UserId = userId
	this.Role = role
	return &this
}

// NewV1ProjectMemberCreateRequestWithDefaults instantiates a new V1ProjectMemberCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ProjectMemberCreateRequestWithDefaults() *V1ProjectMemberCreateRequest {
	this := V1ProjectMemberCreateRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *V1ProjectMemberCreateRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *V1ProjectMemberCreateRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *V1ProjectMemberCreateRequest) SetUserId(v string) {
	o.UserId = v
}

// GetRole returns the Role field value
func (o *V1ProjectMemberCreateRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *V1ProjectMemberCreateRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *V1ProjectMemberCreateRequest) SetRole(v string) {
	o.Role = v
}

func (o V1ProjectMemberCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ProjectMemberCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1ProjectMemberCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"role",
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

	varV1ProjectMemberCreateRequest := _V1ProjectMemberCreateRequest{}

	err = json.Unmarshal(data, &varV1ProjectMemberCreateRequest)

	if err != nil {
		return err
	}

	*o = V1ProjectMemberCreateRequest(varV1ProjectMemberCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1ProjectMemberCreateRequest struct {
	value *V1ProjectMemberCreateRequest
	isSet bool
}

func (v NullableV1ProjectMemberCreateRequest) Get() *V1ProjectMemberCreateRequest {
	return v.value
}

func (v *NullableV1ProjectMemberCreateRequest) Set(val *V1ProjectMemberCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ProjectMemberCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ProjectMemberCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ProjectMemberCreateRequest(val *V1ProjectMemberCreateRequest) *NullableV1ProjectMemberCreateRequest {
	return &NullableV1ProjectMemberCreateRequest{value: val, isSet: true}
}

func (v NullableV1ProjectMemberCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ProjectMemberCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}