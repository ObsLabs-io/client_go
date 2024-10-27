# ChannelModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ProjectId** | **string** |  | 
**Emails** | [**[]ChannelEmailModel**](ChannelEmailModel.md) |  | 
**Webhooks** | [**[]ChannelWebhookModel**](ChannelWebhookModel.md) |  | 
**Slacks** | [**[]ChannelSlackModel**](ChannelSlackModel.md) |  | 
**Smss** | [**[]ChannelSmsModel**](ChannelSmsModel.md) |  | 

## Methods

### NewChannelModel

`func NewChannelModel(id string, name string, projectId string, emails []ChannelEmailModel, webhooks []ChannelWebhookModel, slacks []ChannelSlackModel, smss []ChannelSmsModel, ) *ChannelModel`

NewChannelModel instantiates a new ChannelModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelModelWithDefaults

`func NewChannelModelWithDefaults() *ChannelModel`

NewChannelModelWithDefaults instantiates a new ChannelModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ChannelModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *ChannelModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ChannelModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ChannelModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEmails

`func (o *ChannelModel) GetEmails() []ChannelEmailModel`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ChannelModel) GetEmailsOk() (*[]ChannelEmailModel, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ChannelModel) SetEmails(v []ChannelEmailModel)`

SetEmails sets Emails field to given value.


### GetWebhooks

`func (o *ChannelModel) GetWebhooks() []ChannelWebhookModel`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ChannelModel) GetWebhooksOk() (*[]ChannelWebhookModel, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ChannelModel) SetWebhooks(v []ChannelWebhookModel)`

SetWebhooks sets Webhooks field to given value.


### GetSlacks

`func (o *ChannelModel) GetSlacks() []ChannelSlackModel`

GetSlacks returns the Slacks field if non-nil, zero value otherwise.

### GetSlacksOk

`func (o *ChannelModel) GetSlacksOk() (*[]ChannelSlackModel, bool)`

GetSlacksOk returns a tuple with the Slacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlacks

`func (o *ChannelModel) SetSlacks(v []ChannelSlackModel)`

SetSlacks sets Slacks field to given value.


### GetSmss

`func (o *ChannelModel) GetSmss() []ChannelSmsModel`

GetSmss returns the Smss field if non-nil, zero value otherwise.

### GetSmssOk

`func (o *ChannelModel) GetSmssOk() (*[]ChannelSmsModel, bool)`

GetSmssOk returns a tuple with the Smss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmss

`func (o *ChannelModel) SetSmss(v []ChannelSmsModel)`

SetSmss sets Smss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


