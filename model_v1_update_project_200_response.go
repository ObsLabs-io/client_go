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

// checks if the V1UpdateProject200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1UpdateProject200Response{}

// V1UpdateProject200Response struct for V1UpdateProject200Response
type V1UpdateProject200Response struct {
	Project *ProjectModel `json:"project,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1UpdateProject200Response V1UpdateProject200Response

// NewV1UpdateProject200Response instantiates a new V1UpdateProject200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1UpdateProject200Response() *V1UpdateProject200Response {
	this := V1UpdateProject200Response{}
	return &this
}

// NewV1UpdateProject200ResponseWithDefaults instantiates a new V1UpdateProject200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UpdateProject200ResponseWithDefaults() *V1UpdateProject200Response {
	this := V1UpdateProject200Response{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *V1UpdateProject200Response) GetProject() ProjectModel {
	if o == nil || IsNil(o.Project) {
		var ret ProjectModel
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateProject200Response) GetProjectOk() (*ProjectModel, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *V1UpdateProject200Response) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectModel and assigns it to the Project field.
func (o *V1UpdateProject200Response) SetProject(v ProjectModel) {
	o.Project = &v
}

func (o V1UpdateProject200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1UpdateProject200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1UpdateProject200Response) UnmarshalJSON(data []byte) (err error) {
	varV1UpdateProject200Response := _V1UpdateProject200Response{}

	err = json.Unmarshal(data, &varV1UpdateProject200Response)

	if err != nil {
		return err
	}

	*o = V1UpdateProject200Response(varV1UpdateProject200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1UpdateProject200Response struct {
	value *V1UpdateProject200Response
	isSet bool
}

func (v NullableV1UpdateProject200Response) Get() *V1UpdateProject200Response {
	return v.value
}

func (v *NullableV1UpdateProject200Response) Set(val *V1UpdateProject200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1UpdateProject200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1UpdateProject200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1UpdateProject200Response(val *V1UpdateProject200Response) *NullableV1UpdateProject200Response {
	return &NullableV1UpdateProject200Response{value: val, isSet: true}
}

func (v NullableV1UpdateProject200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1UpdateProject200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


