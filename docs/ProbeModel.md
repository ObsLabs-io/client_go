# ProbeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Url** | **string** |  | 
**Status** | **string** |  | 
**StatusCheckedAt** | **NullableTime** |  | 
**StatusChangedAt** | **NullableTime** |  | 
**Schedule** | [**ProbeScheduleModel**](ProbeScheduleModel.md) |  | 
**Channels** | [**[]ProbeChannelModel**](ProbeChannelModel.md) |  | 

## Methods

### NewProbeModel

`func NewProbeModel(id string, name string, type_ string, url string, status string, statusCheckedAt NullableTime, statusChangedAt NullableTime, schedule ProbeScheduleModel, channels []ProbeChannelModel, ) *ProbeModel`

NewProbeModel instantiates a new ProbeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeModelWithDefaults

`func NewProbeModelWithDefaults() *ProbeModel`

NewProbeModelWithDefaults instantiates a new ProbeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProbeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProbeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProbeModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProbeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProbeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProbeModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ProbeModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProbeModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProbeModel) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *ProbeModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProbeModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProbeModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ProbeModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProbeModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProbeModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusCheckedAt

`func (o *ProbeModel) GetStatusCheckedAt() time.Time`

GetStatusCheckedAt returns the StatusCheckedAt field if non-nil, zero value otherwise.

### GetStatusCheckedAtOk

`func (o *ProbeModel) GetStatusCheckedAtOk() (*time.Time, bool)`

GetStatusCheckedAtOk returns a tuple with the StatusCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCheckedAt

`func (o *ProbeModel) SetStatusCheckedAt(v time.Time)`

SetStatusCheckedAt sets StatusCheckedAt field to given value.


### SetStatusCheckedAtNil

`func (o *ProbeModel) SetStatusCheckedAtNil(b bool)`

 SetStatusCheckedAtNil sets the value for StatusCheckedAt to be an explicit nil

### UnsetStatusCheckedAt
`func (o *ProbeModel) UnsetStatusCheckedAt()`

UnsetStatusCheckedAt ensures that no value is present for StatusCheckedAt, not even an explicit nil
### GetStatusChangedAt

`func (o *ProbeModel) GetStatusChangedAt() time.Time`

GetStatusChangedAt returns the StatusChangedAt field if non-nil, zero value otherwise.

### GetStatusChangedAtOk

`func (o *ProbeModel) GetStatusChangedAtOk() (*time.Time, bool)`

GetStatusChangedAtOk returns a tuple with the StatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedAt

`func (o *ProbeModel) SetStatusChangedAt(v time.Time)`

SetStatusChangedAt sets StatusChangedAt field to given value.


### SetStatusChangedAtNil

`func (o *ProbeModel) SetStatusChangedAtNil(b bool)`

 SetStatusChangedAtNil sets the value for StatusChangedAt to be an explicit nil

### UnsetStatusChangedAt
`func (o *ProbeModel) UnsetStatusChangedAt()`

UnsetStatusChangedAt ensures that no value is present for StatusChangedAt, not even an explicit nil
### GetSchedule

`func (o *ProbeModel) GetSchedule() ProbeScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ProbeModel) GetScheduleOk() (*ProbeScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ProbeModel) SetSchedule(v ProbeScheduleModel)`

SetSchedule sets Schedule field to given value.


### GetChannels

`func (o *ProbeModel) GetChannels() []ProbeChannelModel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ProbeModel) GetChannelsOk() (*[]ProbeChannelModel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ProbeModel) SetChannels(v []ProbeChannelModel)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


