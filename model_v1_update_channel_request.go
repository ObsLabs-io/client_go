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

// checks if the V1UpdateChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1UpdateChannelRequest{}

// V1UpdateChannelRequest struct for V1UpdateChannelRequest
type V1UpdateChannelRequest struct {
	Name *string `json:"name,omitempty"`
	Emails []V1UpdateChannelRequestEmailsInner `json:"emails,omitempty"`
	Webhooks []V1UpdateChannelRequestWebhooksInner `json:"webhooks,omitempty"`
	Slacks []V1UpdateChannelRequestSlacksInner `json:"slacks,omitempty"`
	Smss []V1UpdateChannelRequestSmssInner `json:"smss,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V1UpdateChannelRequest V1UpdateChannelRequest

// NewV1UpdateChannelRequest instantiates a new V1UpdateChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1UpdateChannelRequest() *V1UpdateChannelRequest {
	this := V1UpdateChannelRequest{}
	return &this
}

// NewV1UpdateChannelRequestWithDefaults instantiates a new V1UpdateChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UpdateChannelRequestWithDefaults() *V1UpdateChannelRequest {
	this := V1UpdateChannelRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1UpdateChannelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateChannelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1UpdateChannelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1UpdateChannelRequest) SetName(v string) {
	o.Name = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *V1UpdateChannelRequest) GetEmails() []V1UpdateChannelRequestEmailsInner {
	if o == nil || IsNil(o.Emails) {
		var ret []V1UpdateChannelRequestEmailsInner
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateChannelRequest) GetEmailsOk() ([]V1UpdateChannelRequestEmailsInner, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *V1UpdateChannelRequest) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []V1UpdateChannelRequestEmailsInner and assigns it to the Emails field.
func (o *V1UpdateChannelRequest) SetEmails(v []V1UpdateChannelRequestEmailsInner) {
	o.Emails = v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *V1UpdateChannelRequest) GetWebhooks() []V1UpdateChannelRequestWebhooksInner {
	if o == nil || IsNil(o.Webhooks) {
		var ret []V1UpdateChannelRequestWebhooksInner
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateChannelRequest) GetWebhooksOk() ([]V1UpdateChannelRequestWebhooksInner, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *V1UpdateChannelRequest) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []V1UpdateChannelRequestWebhooksInner and assigns it to the Webhooks field.
func (o *V1UpdateChannelRequest) SetWebhooks(v []V1UpdateChannelRequestWebhooksInner) {
	o.Webhooks = v
}

// GetSlacks returns the Slacks field value if set, zero value otherwise.
func (o *V1UpdateChannelRequest) GetSlacks() []V1UpdateChannelRequestSlacksInner {
	if o == nil || IsNil(o.Slacks) {
		var ret []V1UpdateChannelRequestSlacksInner
		return ret
	}
	return o.Slacks
}

// GetSlacksOk returns a tuple with the Slacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateChannelRequest) GetSlacksOk() ([]V1UpdateChannelRequestSlacksInner, bool) {
	if o == nil || IsNil(o.Slacks) {
		return nil, false
	}
	return o.Slacks, true
}

// HasSlacks returns a boolean if a field has been set.
func (o *V1UpdateChannelRequest) HasSlacks() bool {
	if o != nil && !IsNil(o.Slacks) {
		return true
	}

	return false
}

// SetSlacks gets a reference to the given []V1UpdateChannelRequestSlacksInner and assigns it to the Slacks field.
func (o *V1UpdateChannelRequest) SetSlacks(v []V1UpdateChannelRequestSlacksInner) {
	o.Slacks = v
}

// GetSmss returns the Smss field value if set, zero value otherwise.
func (o *V1UpdateChannelRequest) GetSmss() []V1UpdateChannelRequestSmssInner {
	if o == nil || IsNil(o.Smss) {
		var ret []V1UpdateChannelRequestSmssInner
		return ret
	}
	return o.Smss
}

// GetSmssOk returns a tuple with the Smss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateChannelRequest) GetSmssOk() ([]V1UpdateChannelRequestSmssInner, bool) {
	if o == nil || IsNil(o.Smss) {
		return nil, false
	}
	return o.Smss, true
}

// HasSmss returns a boolean if a field has been set.
func (o *V1UpdateChannelRequest) HasSmss() bool {
	if o != nil && !IsNil(o.Smss) {
		return true
	}

	return false
}

// SetSmss gets a reference to the given []V1UpdateChannelRequestSmssInner and assigns it to the Smss field.
func (o *V1UpdateChannelRequest) SetSmss(v []V1UpdateChannelRequestSmssInner) {
	o.Smss = v
}

func (o V1UpdateChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1UpdateChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	if !IsNil(o.Slacks) {
		toSerialize["slacks"] = o.Slacks
	}
	if !IsNil(o.Smss) {
		toSerialize["smss"] = o.Smss
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1UpdateChannelRequest) UnmarshalJSON(data []byte) (err error) {
	varV1UpdateChannelRequest := _V1UpdateChannelRequest{}

	err = json.Unmarshal(data, &varV1UpdateChannelRequest)

	if err != nil {
		return err
	}

	*o = V1UpdateChannelRequest(varV1UpdateChannelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "webhooks")
		delete(additionalProperties, "slacks")
		delete(additionalProperties, "smss")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1UpdateChannelRequest struct {
	value *V1UpdateChannelRequest
	isSet bool
}

func (v NullableV1UpdateChannelRequest) Get() *V1UpdateChannelRequest {
	return v.value
}

func (v *NullableV1UpdateChannelRequest) Set(val *V1UpdateChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1UpdateChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1UpdateChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1UpdateChannelRequest(val *V1UpdateChannelRequest) *NullableV1UpdateChannelRequest {
	return &NullableV1UpdateChannelRequest{value: val, isSet: true}
}

func (v NullableV1UpdateChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1UpdateChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


