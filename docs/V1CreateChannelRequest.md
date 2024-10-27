# V1CreateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Emails** | Pointer to [**[]V1CreateChannelRequestEmailsInner**](V1CreateChannelRequestEmailsInner.md) |  | [optional] 
**Webhooks** | Pointer to [**[]V1CreateChannelRequestWebhooksInner**](V1CreateChannelRequestWebhooksInner.md) |  | [optional] 
**Slacks** | Pointer to [**[]V1CreateChannelRequestSlacksInner**](V1CreateChannelRequestSlacksInner.md) |  | [optional] 
**Smss** | Pointer to [**[]V1CreateChannelRequestSmssInner**](V1CreateChannelRequestSmssInner.md) |  | [optional] 

## Methods

### NewV1CreateChannelRequest

`func NewV1CreateChannelRequest(name string, ) *V1CreateChannelRequest`

NewV1CreateChannelRequest instantiates a new V1CreateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CreateChannelRequestWithDefaults

`func NewV1CreateChannelRequestWithDefaults() *V1CreateChannelRequest`

NewV1CreateChannelRequestWithDefaults instantiates a new V1CreateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1CreateChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1CreateChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1CreateChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmails

`func (o *V1CreateChannelRequest) GetEmails() []V1CreateChannelRequestEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *V1CreateChannelRequest) GetEmailsOk() (*[]V1CreateChannelRequestEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *V1CreateChannelRequest) SetEmails(v []V1CreateChannelRequestEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *V1CreateChannelRequest) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetWebhooks

`func (o *V1CreateChannelRequest) GetWebhooks() []V1CreateChannelRequestWebhooksInner`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *V1CreateChannelRequest) GetWebhooksOk() (*[]V1CreateChannelRequestWebhooksInner, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *V1CreateChannelRequest) SetWebhooks(v []V1CreateChannelRequestWebhooksInner)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *V1CreateChannelRequest) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetSlacks

`func (o *V1CreateChannelRequest) GetSlacks() []V1CreateChannelRequestSlacksInner`

GetSlacks returns the Slacks field if non-nil, zero value otherwise.

### GetSlacksOk

`func (o *V1CreateChannelRequest) GetSlacksOk() (*[]V1CreateChannelRequestSlacksInner, bool)`

GetSlacksOk returns a tuple with the Slacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlacks

`func (o *V1CreateChannelRequest) SetSlacks(v []V1CreateChannelRequestSlacksInner)`

SetSlacks sets Slacks field to given value.

### HasSlacks

`func (o *V1CreateChannelRequest) HasSlacks() bool`

HasSlacks returns a boolean if a field has been set.

### GetSmss

`func (o *V1CreateChannelRequest) GetSmss() []V1CreateChannelRequestSmssInner`

GetSmss returns the Smss field if non-nil, zero value otherwise.

### GetSmssOk

`func (o *V1CreateChannelRequest) GetSmssOk() (*[]V1CreateChannelRequestSmssInner, bool)`

GetSmssOk returns a tuple with the Smss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmss

`func (o *V1CreateChannelRequest) SetSmss(v []V1CreateChannelRequestSmssInner)`

SetSmss sets Smss field to given value.

### HasSmss

`func (o *V1CreateChannelRequest) HasSmss() bool`

HasSmss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


