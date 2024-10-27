# IntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Slack** | Pointer to [**IntegrationSlackModel**](IntegrationSlackModel.md) |  | [optional] 

## Methods

### NewIntegrationModel

`func NewIntegrationModel(id string, type_ string, ) *IntegrationModel`

NewIntegrationModel instantiates a new IntegrationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationModelWithDefaults

`func NewIntegrationModelWithDefaults() *IntegrationModel`

NewIntegrationModelWithDefaults instantiates a new IntegrationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationModel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IntegrationModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationModel) SetType(v string)`

SetType sets Type field to given value.


### GetSlack

`func (o *IntegrationModel) GetSlack() IntegrationSlackModel`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *IntegrationModel) GetSlackOk() (*IntegrationSlackModel, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *IntegrationModel) SetSlack(v IntegrationSlackModel)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *IntegrationModel) HasSlack() bool`

HasSlack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


