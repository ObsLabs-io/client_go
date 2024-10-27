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

// checks if the UserProjectRoleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProjectRoleModel{}

// UserProjectRoleModel struct for UserProjectRoleModel
type UserProjectRoleModel struct {
	ProjectId            string `json:"project_id"`
	Role                 string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _UserProjectRoleModel UserProjectRoleModel

// NewUserProjectRoleModel instantiates a new UserProjectRoleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProjectRoleModel(projectId string, role string) *UserProjectRoleModel {
	this := UserProjectRoleModel{}
	this.ProjectId = projectId
	this.Role = role
	return &this
}

// NewUserProjectRoleModelWithDefaults instantiates a new UserProjectRoleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProjectRoleModelWithDefaults() *UserProjectRoleModel {
	this := UserProjectRoleModel{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *UserProjectRoleModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UserProjectRoleModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UserProjectRoleModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRole returns the Role field value
func (o *UserProjectRoleModel) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserProjectRoleModel) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserProjectRoleModel) SetRole(v string) {
	o.Role = v
}

func (o UserProjectRoleModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProjectRoleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserProjectRoleModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varUserProjectRoleModel := _UserProjectRoleModel{}

	err = json.Unmarshal(data, &varUserProjectRoleModel)

	if err != nil {
		return err
	}

	*o = UserProjectRoleModel(varUserProjectRoleModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserProjectRoleModel struct {
	value *UserProjectRoleModel
	isSet bool
}

func (v NullableUserProjectRoleModel) Get() *UserProjectRoleModel {
	return v.value
}

func (v *NullableUserProjectRoleModel) Set(val *UserProjectRoleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProjectRoleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProjectRoleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProjectRoleModel(val *UserProjectRoleModel) *NullableUserProjectRoleModel {
	return &NullableUserProjectRoleModel{value: val, isSet: true}
}

func (v NullableUserProjectRoleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProjectRoleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}