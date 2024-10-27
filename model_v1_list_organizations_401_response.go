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

// checks if the V1ListOrganizations401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ListOrganizations401Response{}

// V1ListOrganizations401Response struct for V1ListOrganizations401Response
type V1ListOrganizations401Response struct {
	Error                ErrorModel `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _V1ListOrganizations401Response V1ListOrganizations401Response

// NewV1ListOrganizations401Response instantiates a new V1ListOrganizations401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListOrganizations401Response(error_ ErrorModel) *V1ListOrganizations401Response {
	this := V1ListOrganizations401Response{}
	this.Error = error_
	return &this
}

// NewV1ListOrganizations401ResponseWithDefaults instantiates a new V1ListOrganizations401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListOrganizations401ResponseWithDefaults() *V1ListOrganizations401Response {
	this := V1ListOrganizations401Response{}
	return &this
}

// GetError returns the Error field value
func (o *V1ListOrganizations401Response) GetError() ErrorModel {
	if o == nil {
		var ret ErrorModel
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *V1ListOrganizations401Response) GetErrorOk() (*ErrorModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *V1ListOrganizations401Response) SetError(v ErrorModel) {
	o.Error = v
}

func (o V1ListOrganizations401Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ListOrganizations401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1ListOrganizations401Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varV1ListOrganizations401Response := _V1ListOrganizations401Response{}

	err = json.Unmarshal(data, &varV1ListOrganizations401Response)

	if err != nil {
		return err
	}

	*o = V1ListOrganizations401Response(varV1ListOrganizations401Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1ListOrganizations401Response struct {
	value *V1ListOrganizations401Response
	isSet bool
}

func (v NullableV1ListOrganizations401Response) Get() *V1ListOrganizations401Response {
	return v.value
}

func (v *NullableV1ListOrganizations401Response) Set(val *V1ListOrganizations401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListOrganizations401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListOrganizations401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListOrganizations401Response(val *V1ListOrganizations401Response) *NullableV1ListOrganizations401Response {
	return &NullableV1ListOrganizations401Response{value: val, isSet: true}
}

func (v NullableV1ListOrganizations401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListOrganizations401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
