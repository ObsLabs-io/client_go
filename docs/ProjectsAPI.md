# \ProjectsAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateProject**](ProjectsAPI.md#V1CreateProject) | **Post** /v1/organizations/{oID}/projects | Create Project
[**V1DeleteProject**](ProjectsAPI.md#V1DeleteProject) | **Delete** /v1/organizations/{oID}/projects/{pID} | Delete Project
[**V1GetProject**](ProjectsAPI.md#V1GetProject) | **Get** /v1/organizations/{oID}/projects/{pID} | Get Project
[**V1ListProjects**](ProjectsAPI.md#V1ListProjects) | **Get** /v1/organizations/{oID}/projects | List Projects
[**V1ProjectMemberCreate**](ProjectsAPI.md#V1ProjectMemberCreate) | **Post** /v1/organizations/{oID}/projects/{pID}/members | Add Project Member
[**V1ProjectMemberDelete**](ProjectsAPI.md#V1ProjectMemberDelete) | **Delete** /v1/organizations/{oID}/projects/{pID}/members/{uID} | Delete Project Member
[**V1ProjectMemberUpdate**](ProjectsAPI.md#V1ProjectMemberUpdate) | **Patch** /v1/organizations/{oID}/projects/{pID}/members/{uID} | Update Project Member
[**V1UpdateProject**](ProjectsAPI.md#V1UpdateProject) | **Patch** /v1/organizations/{oID}/projects/{pID} | Update Project



## V1CreateProject

> V1CreateProject201Response V1CreateProject(ctx, oID).V1CreateProjectRequest(v1CreateProjectRequest).Execute()

Create Project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	v1CreateProjectRequest := *openapiclient.NewV1CreateProjectRequest("Name_example") // V1CreateProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1CreateProject(context.Background(), oID).V1CreateProjectRequest(v1CreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateProject`: V1CreateProject201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1CreateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1CreateProjectRequest** | [**V1CreateProjectRequest**](V1CreateProjectRequest.md) |  | 

### Return type

[**V1CreateProject201Response**](V1CreateProject201Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteProject

> V1DeleteProject(ctx, oID, pID).Execute()

Delete Project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.V1DeleteProject(context.Background(), oID, pID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetProject

> V1CreateProject201Response V1GetProject(ctx, oID, pID).Execute()

Get Project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1GetProject(context.Background(), oID, pID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1GetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetProject`: V1CreateProject201Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1CreateProject201Response**](V1CreateProject201Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListProjects

> V1ListProjects200Response V1ListProjects(ctx, oID).Execute()

List Projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1ListProjects(context.Background(), oID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1ListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListProjects`: V1ListProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1ListProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListProjects200Response**](V1ListProjects200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectMemberCreate

> V1ProjectMemberCreate200Response V1ProjectMemberCreate(ctx, oID, pID).V1ProjectMemberCreateRequest(v1ProjectMemberCreateRequest).Execute()

Add Project Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID
	v1ProjectMemberCreateRequest := *openapiclient.NewV1ProjectMemberCreateRequest("UserId_example", "Role_example") // V1ProjectMemberCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1ProjectMemberCreate(context.Background(), oID, pID).V1ProjectMemberCreateRequest(v1ProjectMemberCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1ProjectMemberCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ProjectMemberCreate`: V1ProjectMemberCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1ProjectMemberCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectMemberCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1ProjectMemberCreateRequest** | [**V1ProjectMemberCreateRequest**](V1ProjectMemberCreateRequest.md) |  | 

### Return type

[**V1ProjectMemberCreate200Response**](V1ProjectMemberCreate200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectMemberDelete

> V1ProjectMemberDelete(ctx, oID, pID, uID).Execute()

Delete Project Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID
	uID := "iaI_KnKzSNxMhqwNKfvEuw" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.V1ProjectMemberDelete(context.Background(), oID, pID, uID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1ProjectMemberDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**uID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectMemberDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectMemberUpdate

> V1ProjectMemberCreate200Response V1ProjectMemberUpdate(ctx, oID, pID, uID).V1ProjectMemberUpdateRequest(v1ProjectMemberUpdateRequest).Execute()

Update Project Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID
	uID := "iaI_KnKzSNxMhqwNKfvEuw" // string | User ID
	v1ProjectMemberUpdateRequest := *openapiclient.NewV1ProjectMemberUpdateRequest("Role_example") // V1ProjectMemberUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1ProjectMemberUpdate(context.Background(), oID, pID, uID).V1ProjectMemberUpdateRequest(v1ProjectMemberUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1ProjectMemberUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ProjectMemberUpdate`: V1ProjectMemberCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1ProjectMemberUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**uID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectMemberUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1ProjectMemberUpdateRequest** | [**V1ProjectMemberUpdateRequest**](V1ProjectMemberUpdateRequest.md) |  | 

### Return type

[**V1ProjectMemberCreate200Response**](V1ProjectMemberCreate200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateProject

> V1UpdateProject200Response V1UpdateProject(ctx, oID, pID).V1UpdateProjectRequest(v1UpdateProjectRequest).Execute()

Update Project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ObsLabs-io/client_go"
)

func main() {
	oID := "NTz6OnG_yqlm_0qUmjHJjg" // string | Organization ID
	pID := "1_g7EQvDIFjREnMBuRQ30Q" // string | Project ID
	v1UpdateProjectRequest := *openapiclient.NewV1UpdateProjectRequest() // V1UpdateProjectRequest | Update Project (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.V1UpdateProject(context.Background(), oID, pID).V1UpdateProjectRequest(v1UpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.V1UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UpdateProject`: V1UpdateProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.V1UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1UpdateProjectRequest** | [**V1UpdateProjectRequest**](V1UpdateProjectRequest.md) | Update Project | 

### Return type

[**V1UpdateProject200Response**](V1UpdateProject200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

