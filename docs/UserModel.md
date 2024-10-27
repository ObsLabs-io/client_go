# UserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**OrganizationsRole** | [**[]UserOrganizationRoleModel**](UserOrganizationRoleModel.md) |  | 
**ProjectsRole** | [**[]UserProjectRoleModel**](UserProjectRoleModel.md) |  | 

## Methods

### NewUserModel

`func NewUserModel(id string, name string, email string, organizationsRole []UserOrganizationRoleModel, projectsRole []UserProjectRoleModel, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserModel) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationsRole

`func (o *UserModel) GetOrganizationsRole() []UserOrganizationRoleModel`

GetOrganizationsRole returns the OrganizationsRole field if non-nil, zero value otherwise.

### GetOrganizationsRoleOk

`func (o *UserModel) GetOrganizationsRoleOk() (*[]UserOrganizationRoleModel, bool)`

GetOrganizationsRoleOk returns a tuple with the OrganizationsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsRole

`func (o *UserModel) SetOrganizationsRole(v []UserOrganizationRoleModel)`

SetOrganizationsRole sets OrganizationsRole field to given value.


### GetProjectsRole

`func (o *UserModel) GetProjectsRole() []UserProjectRoleModel`

GetProjectsRole returns the ProjectsRole field if non-nil, zero value otherwise.

### GetProjectsRoleOk

`func (o *UserModel) GetProjectsRoleOk() (*[]UserProjectRoleModel, bool)`

GetProjectsRoleOk returns a tuple with the ProjectsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsRole

`func (o *UserModel) SetProjectsRole(v []UserProjectRoleModel)`

SetProjectsRole sets ProjectsRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


