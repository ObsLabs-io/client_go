# OrganizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IsTrial** | **bool** |  | 
**TrialEndsAt** | Pointer to **time.Time** |  | [optional] 
**Members** | [**[]OrganizationMemberModel**](OrganizationMemberModel.md) |  | 
**Projects** | [**[]OrganizationProjectModel**](OrganizationProjectModel.md) |  | 
**Subscription** | Pointer to [**OrganizationSubscriptionModel**](OrganizationSubscriptionModel.md) |  | [optional] 

## Methods

### NewOrganizationModel

`func NewOrganizationModel(id string, name string, isTrial bool, members []OrganizationMemberModel, projects []OrganizationProjectModel, ) *OrganizationModel`

NewOrganizationModel instantiates a new OrganizationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationModelWithDefaults

`func NewOrganizationModelWithDefaults() *OrganizationModel`

NewOrganizationModelWithDefaults instantiates a new OrganizationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsTrial

`func (o *OrganizationModel) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *OrganizationModel) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *OrganizationModel) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.


### GetTrialEndsAt

`func (o *OrganizationModel) GetTrialEndsAt() time.Time`

GetTrialEndsAt returns the TrialEndsAt field if non-nil, zero value otherwise.

### GetTrialEndsAtOk

`func (o *OrganizationModel) GetTrialEndsAtOk() (*time.Time, bool)`

GetTrialEndsAtOk returns a tuple with the TrialEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndsAt

`func (o *OrganizationModel) SetTrialEndsAt(v time.Time)`

SetTrialEndsAt sets TrialEndsAt field to given value.

### HasTrialEndsAt

`func (o *OrganizationModel) HasTrialEndsAt() bool`

HasTrialEndsAt returns a boolean if a field has been set.

### GetMembers

`func (o *OrganizationModel) GetMembers() []OrganizationMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationModel) GetMembersOk() (*[]OrganizationMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationModel) SetMembers(v []OrganizationMemberModel)`

SetMembers sets Members field to given value.


### GetProjects

`func (o *OrganizationModel) GetProjects() []OrganizationProjectModel`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationModel) GetProjectsOk() (*[]OrganizationProjectModel, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationModel) SetProjects(v []OrganizationProjectModel)`

SetProjects sets Projects field to given value.


### GetSubscription

`func (o *OrganizationModel) GetSubscription() OrganizationSubscriptionModel`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *OrganizationModel) GetSubscriptionOk() (*OrganizationSubscriptionModel, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *OrganizationModel) SetSubscription(v OrganizationSubscriptionModel)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *OrganizationModel) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


