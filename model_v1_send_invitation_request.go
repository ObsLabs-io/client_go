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

// checks if the V1SendInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SendInvitationRequest{}

// V1SendInvitationRequest struct for V1SendInvitationRequest
type V1SendInvitationRequest struct {
	Email                string `json:"email"`
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _V1SendInvitationRequest V1SendInvitationRequest

// NewV1SendInvitationRequest instantiates a new V1SendInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SendInvitationRequest(email string, role string) *V1SendInvitationRequest {
	this := V1SendInvitationRequest{}
	this.Email = email
	this.Role = role
	return &this
}

// NewV1SendInvitationRequestWithDefaults instantiates a new V1SendInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SendInvitationRequestWithDefaults() *V1SendInvitationRequest {
	this := V1SendInvitationRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *V1SendInvitationRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *V1SendInvitationRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *V1SendInvitationRequest) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *V1SendInvitationRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *V1SendInvitationRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *V1SendInvitationRequest) SetRole(v string) {
	o.Role = v
}

func (o V1SendInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SendInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1SendInvitationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varV1SendInvitationRequest := _V1SendInvitationRequest{}

	err = json.Unmarshal(data, &varV1SendInvitationRequest)

	if err != nil {
		return err
	}

	*o = V1SendInvitationRequest(varV1SendInvitationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1SendInvitationRequest struct {
	value *V1SendInvitationRequest
	isSet bool
}

func (v NullableV1SendInvitationRequest) Get() *V1SendInvitationRequest {
	return v.value
}

func (v *NullableV1SendInvitationRequest) Set(val *V1SendInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SendInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SendInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SendInvitationRequest(val *V1SendInvitationRequest) *NullableV1SendInvitationRequest {
	return &NullableV1SendInvitationRequest{value: val, isSet: true}
}

func (v NullableV1SendInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SendInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
