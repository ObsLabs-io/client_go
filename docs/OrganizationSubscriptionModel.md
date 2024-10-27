# OrganizationSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Plan** | **string** |  | 
**Status** | **string** |  | 
**NextBilledAt** | Pointer to **time.Time** |  | [optional] 
**ScheduledChange** | Pointer to [**OrganizationSubscriptionScheduledChangeModel**](OrganizationSubscriptionScheduledChangeModel.md) |  | [optional] 
**PaddleCustomerId** | **string** |  | 

## Methods

### NewOrganizationSubscriptionModel

`func NewOrganizationSubscriptionModel(id string, plan string, status string, paddleCustomerId string, ) *OrganizationSubscriptionModel`

NewOrganizationSubscriptionModel instantiates a new OrganizationSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionModelWithDefaults

`func NewOrganizationSubscriptionModelWithDefaults() *OrganizationSubscriptionModel`

NewOrganizationSubscriptionModelWithDefaults instantiates a new OrganizationSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationSubscriptionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSubscriptionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSubscriptionModel) SetId(v string)`

SetId sets Id field to given value.


### GetPlan

`func (o *OrganizationSubscriptionModel) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationSubscriptionModel) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationSubscriptionModel) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetStatus

`func (o *OrganizationSubscriptionModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationSubscriptionModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationSubscriptionModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNextBilledAt

`func (o *OrganizationSubscriptionModel) GetNextBilledAt() time.Time`

GetNextBilledAt returns the NextBilledAt field if non-nil, zero value otherwise.

### GetNextBilledAtOk

`func (o *OrganizationSubscriptionModel) GetNextBilledAtOk() (*time.Time, bool)`

GetNextBilledAtOk returns a tuple with the NextBilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBilledAt

`func (o *OrganizationSubscriptionModel) SetNextBilledAt(v time.Time)`

SetNextBilledAt sets NextBilledAt field to given value.

### HasNextBilledAt

`func (o *OrganizationSubscriptionModel) HasNextBilledAt() bool`

HasNextBilledAt returns a boolean if a field has been set.

### GetScheduledChange

`func (o *OrganizationSubscriptionModel) GetScheduledChange() OrganizationSubscriptionScheduledChangeModel`

GetScheduledChange returns the ScheduledChange field if non-nil, zero value otherwise.

### GetScheduledChangeOk

`func (o *OrganizationSubscriptionModel) GetScheduledChangeOk() (*OrganizationSubscriptionScheduledChangeModel, bool)`

GetScheduledChangeOk returns a tuple with the ScheduledChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledChange

`func (o *OrganizationSubscriptionModel) SetScheduledChange(v OrganizationSubscriptionScheduledChangeModel)`

SetScheduledChange sets ScheduledChange field to given value.

### HasScheduledChange

`func (o *OrganizationSubscriptionModel) HasScheduledChange() bool`

HasScheduledChange returns a boolean if a field has been set.

### GetPaddleCustomerId

`func (o *OrganizationSubscriptionModel) GetPaddleCustomerId() string`

GetPaddleCustomerId returns the PaddleCustomerId field if non-nil, zero value otherwise.

### GetPaddleCustomerIdOk

`func (o *OrganizationSubscriptionModel) GetPaddleCustomerIdOk() (*string, bool)`

GetPaddleCustomerIdOk returns a tuple with the PaddleCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddleCustomerId

`func (o *OrganizationSubscriptionModel) SetPaddleCustomerId(v string)`

SetPaddleCustomerId sets PaddleCustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


