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

// checks if the ProjectMemberModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMemberModel{}

// ProjectMemberModel struct for ProjectMemberModel
type ProjectMemberModel struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Role                 string `json:"role"`
	Email                string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _ProjectMemberModel ProjectMemberModel

// NewProjectMemberModel instantiates a new ProjectMemberModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMemberModel(id string, name string, role string, email string) *ProjectMemberModel {
	this := ProjectMemberModel{}
	this.Id = id
	this.Name = name
	this.Role = role
	this.Email = email
	return &this
}

// NewProjectMemberModelWithDefaults instantiates a new ProjectMemberModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMemberModelWithDefaults() *ProjectMemberModel {
	this := ProjectMemberModel{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectMemberModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectMemberModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectMemberModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectMemberModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMemberModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMemberModel) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *ProjectMemberModel) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProjectMemberModel) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProjectMemberModel) SetRole(v string) {
	o.Role = v
}

// GetEmail returns the Email field value
func (o *ProjectMemberModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ProjectMemberModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ProjectMemberModel) SetEmail(v string) {
	o.Email = v
}

func (o ProjectMemberModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMemberModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectMemberModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"role",
		"email",
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

	varProjectMemberModel := _ProjectMemberModel{}

	err = json.Unmarshal(data, &varProjectMemberModel)

	if err != nil {
		return err
	}

	*o = ProjectMemberModel(varProjectMemberModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectMemberModel struct {
	value *ProjectMemberModel
	isSet bool
}

func (v NullableProjectMemberModel) Get() *ProjectMemberModel {
	return v.value
}

func (v *NullableProjectMemberModel) Set(val *ProjectMemberModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMemberModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMemberModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMemberModel(val *ProjectMemberModel) *NullableProjectMemberModel {
	return &NullableProjectMemberModel{value: val, isSet: true}
}

func (v NullableProjectMemberModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMemberModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}