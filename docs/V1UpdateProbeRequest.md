# V1UpdateProbeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Only pull probes can have updated URL | [optional] 
**Schedule** | Pointer to [**ProbeScheduleModel**](ProbeScheduleModel.md) |  | [optional] 
**Channels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1UpdateProbeRequest

`func NewV1UpdateProbeRequest() *V1UpdateProbeRequest`

NewV1UpdateProbeRequest instantiates a new V1UpdateProbeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UpdateProbeRequestWithDefaults

`func NewV1UpdateProbeRequestWithDefaults() *V1UpdateProbeRequest`

NewV1UpdateProbeRequestWithDefaults instantiates a new V1UpdateProbeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1UpdateProbeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UpdateProbeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UpdateProbeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1UpdateProbeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *V1UpdateProbeRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1UpdateProbeRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1UpdateProbeRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1UpdateProbeRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSchedule

`func (o *V1UpdateProbeRequest) GetSchedule() ProbeScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *V1UpdateProbeRequest) GetScheduleOk() (*ProbeScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *V1UpdateProbeRequest) SetSchedule(v ProbeScheduleModel)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *V1UpdateProbeRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetChannels

`func (o *V1UpdateProbeRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *V1UpdateProbeRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *V1UpdateProbeRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *V1UpdateProbeRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


