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

// checks if the V1ListProjects200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ListProjects200Response{}

// V1ListProjects200Response struct for V1ListProjects200Response
type V1ListProjects200Response struct {
	Projects             []ProjectModel `json:"projects"`
	AdditionalProperties map[string]interface{}
}

type _V1ListProjects200Response V1ListProjects200Response

// NewV1ListProjects200Response instantiates a new V1ListProjects200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListProjects200Response(projects []ProjectModel) *V1ListProjects200Response {
	this := V1ListProjects200Response{}
	this.Projects = projects
	return &this
}

// NewV1ListProjects200ResponseWithDefaults instantiates a new V1ListProjects200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListProjects200ResponseWithDefaults() *V1ListProjects200Response {
	this := V1ListProjects200Response{}
	return &this
}

// GetProjects returns the Projects field value
func (o *V1ListProjects200Response) GetProjects() []ProjectModel {
	if o == nil {
		var ret []ProjectModel
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *V1ListProjects200Response) GetProjectsOk() ([]ProjectModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *V1ListProjects200Response) SetProjects(v []ProjectModel) {
	o.Projects = v
}

func (o V1ListProjects200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ListProjects200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projects"] = o.Projects

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1ListProjects200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projects",
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

	varV1ListProjects200Response := _V1ListProjects200Response{}

	err = json.Unmarshal(data, &varV1ListProjects200Response)

	if err != nil {
		return err
	}

	*o = V1ListProjects200Response(varV1ListProjects200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "projects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1ListProjects200Response struct {
	value *V1ListProjects200Response
	isSet bool
}

func (v NullableV1ListProjects200Response) Get() *V1ListProjects200Response {
	return v.value
}

func (v *NullableV1ListProjects200Response) Set(val *V1ListProjects200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListProjects200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListProjects200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListProjects200Response(val *V1ListProjects200Response) *NullableV1ListProjects200Response {
	return &NullableV1ListProjects200Response{value: val, isSet: true}
}

func (v NullableV1ListProjects200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListProjects200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
