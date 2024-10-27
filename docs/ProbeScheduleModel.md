# ProbeScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**FailureThreshold** | **int32** |  | [default to 1]
**SuccessThreshold** | **int32** |  | [default to 1]
**Expression** | Pointer to **string** | Required if type is \&quot;cron\&quot; | [optional] 
**PeriodSeconds** | Pointer to **int32** | Required if type is \&quot;period\&quot; | [optional] [default to 360]

## Methods

### NewProbeScheduleModel

`func NewProbeScheduleModel(type_ string, failureThreshold int32, successThreshold int32, ) *ProbeScheduleModel`

NewProbeScheduleModel instantiates a new ProbeScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeScheduleModelWithDefaults

`func NewProbeScheduleModelWithDefaults() *ProbeScheduleModel`

NewProbeScheduleModelWithDefaults instantiates a new ProbeScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProbeScheduleModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProbeScheduleModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProbeScheduleModel) SetType(v string)`

SetType sets Type field to given value.


### GetFailureThreshold

`func (o *ProbeScheduleModel) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *ProbeScheduleModel) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *ProbeScheduleModel) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.


### GetSuccessThreshold

`func (o *ProbeScheduleModel) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *ProbeScheduleModel) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *ProbeScheduleModel) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.


### GetExpression

`func (o *ProbeScheduleModel) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ProbeScheduleModel) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ProbeScheduleModel) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ProbeScheduleModel) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *ProbeScheduleModel) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *ProbeScheduleModel) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *ProbeScheduleModel) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *ProbeScheduleModel) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


