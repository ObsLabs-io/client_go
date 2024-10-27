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

// checks if the UserOrganizationRoleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserOrganizationRoleModel{}

// UserOrganizationRoleModel struct for UserOrganizationRoleModel
type UserOrganizationRoleModel struct {
	OrganizationId       string `json:"organization_id"`
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _UserOrganizationRoleModel UserOrganizationRoleModel

// NewUserOrganizationRoleModel instantiates a new UserOrganizationRoleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOrganizationRoleModel(organizationId string, role string) *UserOrganizationRoleModel {
	this := UserOrganizationRoleModel{}
	this.OrganizationId = organizationId
	this.Role = role
	return &this
}

// NewUserOrganizationRoleModelWithDefaults instantiates a new UserOrganizationRoleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOrganizationRoleModelWithDefaults() *UserOrganizationRoleModel {
	this := UserOrganizationRoleModel{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value
func (o *UserOrganizationRoleModel) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *UserOrganizationRoleModel) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *UserOrganizationRoleModel) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetRole returns the Role field value
func (o *UserOrganizationRoleModel) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserOrganizationRoleModel) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserOrganizationRoleModel) SetRole(v string) {
	o.Role = v
}

func (o UserOrganizationRoleModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserOrganizationRoleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organization_id"] = o.OrganizationId
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserOrganizationRoleModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organization_id",
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

	varUserOrganizationRoleModel := _UserOrganizationRoleModel{}

	err = json.Unmarshal(data, &varUserOrganizationRoleModel)

	if err != nil {
		return err
	}

	*o = UserOrganizationRoleModel(varUserOrganizationRoleModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserOrganizationRoleModel struct {
	value *UserOrganizationRoleModel
	isSet bool
}

func (v NullableUserOrganizationRoleModel) Get() *UserOrganizationRoleModel {
	return v.value
}

func (v *NullableUserOrganizationRoleModel) Set(val *UserOrganizationRoleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOrganizationRoleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOrganizationRoleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOrganizationRoleModel(val *UserOrganizationRoleModel) *NullableUserOrganizationRoleModel {
	return &NullableUserOrganizationRoleModel{value: val, isSet: true}
}

func (v NullableUserOrganizationRoleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOrganizationRoleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}