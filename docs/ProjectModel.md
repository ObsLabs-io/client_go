# ProjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**Members** | [**[]ProjectMemberModel**](ProjectMemberModel.md) |  | 

## Methods

### NewProjectModel

`func NewProjectModel(id string, name string, organizationId string, members []ProjectMemberModel, ) *ProjectModel`

NewProjectModel instantiates a new ProjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectModelWithDefaults

`func NewProjectModelWithDefaults() *ProjectModel`

NewProjectModelWithDefaults instantiates a new ProjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectModel) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ProjectModel) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectModel) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectModel) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetMembers

`func (o *ProjectModel) GetMembers() []ProjectMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ProjectModel) GetMembersOk() (*[]ProjectMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ProjectModel) SetMembers(v []ProjectMemberModel)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


