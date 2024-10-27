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

// checks if the V1UpdateProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1UpdateProjectRequest{}

// V1UpdateProjectRequest struct for V1UpdateProjectRequest
type V1UpdateProjectRequest struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1UpdateProjectRequest V1UpdateProjectRequest

// NewV1UpdateProjectRequest instantiates a new V1UpdateProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1UpdateProjectRequest() *V1UpdateProjectRequest {
	this := V1UpdateProjectRequest{}
	return &this
}

// NewV1UpdateProjectRequestWithDefaults instantiates a new V1UpdateProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UpdateProjectRequestWithDefaults() *V1UpdateProjectRequest {
	this := V1UpdateProjectRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1UpdateProjectRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateProjectRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1UpdateProjectRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1UpdateProjectRequest) SetName(v string) {
	o.Name = &v
}

func (o V1UpdateProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1UpdateProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1UpdateProjectRequest) UnmarshalJSON(data []byte) (err error) {
	varV1UpdateProjectRequest := _V1UpdateProjectRequest{}

	err = json.Unmarshal(data, &varV1UpdateProjectRequest)

	if err != nil {
		return err
	}

	*o = V1UpdateProjectRequest(varV1UpdateProjectRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1UpdateProjectRequest struct {
	value *V1UpdateProjectRequest
	isSet bool
}

func (v NullableV1UpdateProjectRequest) Get() *V1UpdateProjectRequest {
	return v.value
}

func (v *NullableV1UpdateProjectRequest) Set(val *V1UpdateProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1UpdateProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1UpdateProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1UpdateProjectRequest(val *V1UpdateProjectRequest) *NullableV1UpdateProjectRequest {
	return &NullableV1UpdateProjectRequest{value: val, isSet: true}
}

func (v NullableV1UpdateProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1UpdateProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}