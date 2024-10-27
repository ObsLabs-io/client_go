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

// checks if the V1CreateOrganization201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateOrganization201Response{}

// V1CreateOrganization201Response struct for V1CreateOrganization201Response
type V1CreateOrganization201Response struct {
	Organization OrganizationModel `json:"organization"`
	AdditionalProperties map[string]interface{}
}

type _V1CreateOrganization201Response V1CreateOrganization201Response

// NewV1CreateOrganization201Response instantiates a new V1CreateOrganization201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateOrganization201Response(organization OrganizationModel) *V1CreateOrganization201Response {
	this := V1CreateOrganization201Response{}
	this.Organization = organization
	return &this
}

// NewV1CreateOrganization201ResponseWithDefaults instantiates a new V1CreateOrganization201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateOrganization201ResponseWithDefaults() *V1CreateOrganization201Response {
	this := V1CreateOrganization201Response{}
	return &this
}

// GetOrganization returns the Organization field value
func (o *V1CreateOrganization201Response) GetOrganization() OrganizationModel {
	if o == nil {
		var ret OrganizationModel
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *V1CreateOrganization201Response) GetOrganizationOk() (*OrganizationModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *V1CreateOrganization201Response) SetOrganization(v OrganizationModel) {
	o.Organization = v
}

func (o V1CreateOrganization201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateOrganization201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organization"] = o.Organization

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1CreateOrganization201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organization",
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

	varV1CreateOrganization201Response := _V1CreateOrganization201Response{}

	err = json.Unmarshal(data, &varV1CreateOrganization201Response)

	if err != nil {
		return err
	}

	*o = V1CreateOrganization201Response(varV1CreateOrganization201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1CreateOrganization201Response struct {
	value *V1CreateOrganization201Response
	isSet bool
}

func (v NullableV1CreateOrganization201Response) Get() *V1CreateOrganization201Response {
	return v.value
}

func (v *NullableV1CreateOrganization201Response) Set(val *V1CreateOrganization201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateOrganization201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateOrganization201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateOrganization201Response(val *V1CreateOrganization201Response) *NullableV1CreateOrganization201Response {
	return &NullableV1CreateOrganization201Response{value: val, isSet: true}
}

func (v NullableV1CreateOrganization201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateOrganization201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


