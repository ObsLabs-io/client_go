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

// checks if the V1CreateApiKey201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateApiKey201Response{}

// V1CreateApiKey201Response struct for V1CreateApiKey201Response
type V1CreateApiKey201Response struct {
	ApiKey               ApiKeyModel `json:"api_key"`
	AdditionalProperties map[string]interface{}
}

type _V1CreateApiKey201Response V1CreateApiKey201Response

// NewV1CreateApiKey201Response instantiates a new V1CreateApiKey201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateApiKey201Response(apiKey ApiKeyModel) *V1CreateApiKey201Response {
	this := V1CreateApiKey201Response{}
	this.ApiKey = apiKey
	return &this
}

// NewV1CreateApiKey201ResponseWithDefaults instantiates a new V1CreateApiKey201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateApiKey201ResponseWithDefaults() *V1CreateApiKey201Response {
	this := V1CreateApiKey201Response{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *V1CreateApiKey201Response) GetApiKey() ApiKeyModel {
	if o == nil {
		var ret ApiKeyModel
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *V1CreateApiKey201Response) GetApiKeyOk() (*ApiKeyModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *V1CreateApiKey201Response) SetApiKey(v ApiKeyModel) {
	o.ApiKey = v
}

func (o V1CreateApiKey201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateApiKey201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_key"] = o.ApiKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1CreateApiKey201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api_key",
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

	varV1CreateApiKey201Response := _V1CreateApiKey201Response{}

	err = json.Unmarshal(data, &varV1CreateApiKey201Response)

	if err != nil {
		return err
	}

	*o = V1CreateApiKey201Response(varV1CreateApiKey201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1CreateApiKey201Response struct {
	value *V1CreateApiKey201Response
	isSet bool
}

func (v NullableV1CreateApiKey201Response) Get() *V1CreateApiKey201Response {
	return v.value
}

func (v *NullableV1CreateApiKey201Response) Set(val *V1CreateApiKey201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateApiKey201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateApiKey201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateApiKey201Response(val *V1CreateApiKey201Response) *NullableV1CreateApiKey201Response {
	return &NullableV1CreateApiKey201Response{value: val, isSet: true}
}

func (v NullableV1CreateApiKey201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateApiKey201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}