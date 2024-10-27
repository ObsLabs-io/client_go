# V1UpdateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]V1UpdateChannelRequestEmailsInner**](V1UpdateChannelRequestEmailsInner.md) |  | [optional] 
**Webhooks** | Pointer to [**[]V1UpdateChannelRequestWebhooksInner**](V1UpdateChannelRequestWebhooksInner.md) |  | [optional] 
**Slacks** | Pointer to [**[]V1UpdateChannelRequestSlacksInner**](V1UpdateChannelRequestSlacksInner.md) |  | [optional] 
**Smss** | Pointer to [**[]V1UpdateChannelRequestSmssInner**](V1UpdateChannelRequestSmssInner.md) |  | [optional] 

## Methods

### NewV1UpdateChannelRequest

`func NewV1UpdateChannelRequest() *V1UpdateChannelRequest`

NewV1UpdateChannelRequest instantiates a new V1UpdateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UpdateChannelRequestWithDefaults

`func NewV1UpdateChannelRequestWithDefaults() *V1UpdateChannelRequest`

NewV1UpdateChannelRequestWithDefaults instantiates a new V1UpdateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1UpdateChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UpdateChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UpdateChannelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1UpdateChannelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmails

`func (o *V1UpdateChannelRequest) GetEmails() []V1UpdateChannelRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *V1UpdateChannelRequest) GetEmailsOk() (*[]V1UpdateChannelRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *V1UpdateChannelRequest) SetEmails(v []V1UpdateChannelRequestEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *V1UpdateChannelRequest) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetWebhooks

`func (o *V1UpdateChannelRequest) GetWebhooks() []V1UpdateChannelRequestWebhooksInner`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *V1UpdateChannelRequest) GetWebhooksOk() (*[]V1UpdateChannelRequestWebhooksInner, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *V1UpdateChannelRequest) SetWebhooks(v []V1UpdateChannelRequestWebhooksInner)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *V1UpdateChannelRequest) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetSlacks

`func (o *V1UpdateChannelRequest) GetSlacks() []V1UpdateChannelRequestSlacksInner`

GetSlacks returns the Slacks field if non-nil, zero value otherwise.

### GetSlacksOk

`func (o *V1UpdateChannelRequest) GetSlacksOk() (*[]V1UpdateChannelRequestSlacksInner, bool)`

GetSlacksOk returns a tuple with the Slacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlacks

`func (o *V1UpdateChannelRequest) SetSlacks(v []V1UpdateChannelRequestSlacksInner)`

SetSlacks sets Slacks field to given value.

### HasSlacks

`func (o *V1UpdateChannelRequest) HasSlacks() bool`

HasSlacks returns a boolean if a field has been set.

### GetSmss

`func (o *V1UpdateChannelRequest) GetSmss() []V1UpdateChannelRequestSmssInner`

GetSmss returns the Smss field if non-nil, zero value otherwise.

### GetSmssOk

`func (o *V1UpdateChannelRequest) GetSmssOk() (*[]V1UpdateChannelRequestSmssInner, bool)`

GetSmssOk returns a tuple with the Smss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmss

`func (o *V1UpdateChannelRequest) SetSmss(v []V1UpdateChannelRequestSmssInner)`

SetSmss sets Smss field to given value.

### HasSmss

`func (o *V1UpdateChannelRequest) HasSmss() bool`

HasSmss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


