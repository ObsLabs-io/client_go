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
)

// checks if the V1UpdateOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1UpdateOrganizationRequest{}

// V1UpdateOrganizationRequest struct for V1UpdateOrganizationRequest
type V1UpdateOrganizationRequest struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1UpdateOrganizationRequest V1UpdateOrganizationRequest

// NewV1UpdateOrganizationRequest instantiates a new V1UpdateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1UpdateOrganizationRequest() *V1UpdateOrganizationRequest {
	this := V1UpdateOrganizationRequest{}
	return &this
}

// NewV1UpdateOrganizationRequestWithDefaults instantiates a new V1UpdateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UpdateOrganizationRequestWithDefaults() *V1UpdateOrganizationRequest {
	this := V1UpdateOrganizationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1UpdateOrganizationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1UpdateOrganizationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1UpdateOrganizationRequest) SetName(v string) {
	o.Name = &v
}

func (o V1UpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1UpdateOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1UpdateOrganizationRequest) UnmarshalJSON(data []byte) (err error) {
	varV1UpdateOrganizationRequest := _V1UpdateOrganizationRequest{}

	err = json.Unmarshal(data, &varV1UpdateOrganizationRequest)

	if err != nil {
		return err
	}

	*o = V1UpdateOrganizationRequest(varV1UpdateOrganizationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1UpdateOrganizationRequest struct {
	value *V1UpdateOrganizationRequest
	isSet bool
}

func (v NullableV1UpdateOrganizationRequest) Get() *V1UpdateOrganizationRequest {
	return v.value
}

func (v *NullableV1UpdateOrganizationRequest) Set(val *V1UpdateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1UpdateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1UpdateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1UpdateOrganizationRequest(val *V1UpdateOrganizationRequest) *NullableV1UpdateOrganizationRequest {
	return &NullableV1UpdateOrganizationRequest{value: val, isSet: true}
}

func (v NullableV1UpdateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1UpdateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
