# OrganizationSubscriptionScheduledChangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**EffectiveAt** | **time.Time** |  | 
**ResumeAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrganizationSubscriptionScheduledChangeModel

`func NewOrganizationSubscriptionScheduledChangeModel(action string, effectiveAt time.Time, ) *OrganizationSubscriptionScheduledChangeModel`

NewOrganizationSubscriptionScheduledChangeModel instantiates a new OrganizationSubscriptionScheduledChangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSubscriptionScheduledChangeModelWithDefaults

`func NewOrganizationSubscriptionScheduledChangeModelWithDefaults() *OrganizationSubscriptionScheduledChangeModel`

NewOrganizationSubscriptionScheduledChangeModelWithDefaults instantiates a new OrganizationSubscriptionScheduledChangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrganizationSubscriptionScheduledChangeModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrganizationSubscriptionScheduledChangeModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrganizationSubscriptionScheduledChangeModel) SetAction(v string)`

SetAction sets Action field to given value.


### GetEffectiveAt

`func (o *OrganizationSubscriptionScheduledChangeModel) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *OrganizationSubscriptionScheduledChangeModel) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *OrganizationSubscriptionScheduledChangeModel) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetResumeAt

`func (o *OrganizationSubscriptionScheduledChangeModel) GetResumeAt() time.Time`

GetResumeAt returns the ResumeAt field if non-nil, zero value otherwise.

### GetResumeAtOk

`func (o *OrganizationSubscriptionScheduledChangeModel) GetResumeAtOk() (*time.Time, bool)`

GetResumeAtOk returns a tuple with the ResumeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAt

`func (o *OrganizationSubscriptionScheduledChangeModel) SetResumeAt(v time.Time)`

SetResumeAt sets ResumeAt field to given value.

### HasResumeAt

`func (o *OrganizationSubscriptionScheduledChangeModel) HasResumeAt() bool`

HasResumeAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


