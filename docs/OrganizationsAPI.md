# \OrganizationsAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateOrganization**](OrganizationsAPI.md#V1CreateOrganization) | **Post** /v1/organizations | Create Organization
[**V1DeleteOrganization**](OrganizationsAPI.md#V1DeleteOrganization) | **Delete** /v1/organizations/{oID} | Delete Organization
[**V1GetOrganization**](OrganizationsAPI.md#V1GetOrganization) | **Get** /v1/organizations/{oID} | Get Organization
[**V1ListOrganizations**](OrganizationsAPI.md#V1ListOrganizations) | **Get** /v1/organizations | List Organizations
[**V1OrganizationMemberDelete**](OrganizationsAPI.md#V1OrganizationMemberDelete) | **Delete** /v1/organizations/{oID}/members/{uID} | Delete Organization Memeber
[**V1OrganizationMemberUpdate**](OrganizationsAPI.md#V1OrganizationMemberUpdate) | **Patch** /v1/organizations/{oID}/members/{uID} | Update Organization Member
[**V1UpdateOrganization**](OrganizationsAPI.md#V1UpdateOrganization) | **Patch** /v1/organizations/{oID} | Update Organization



## V1CreateOrganization

> V1CreateOrganization201Response V1CreateOrganization(ctx).V1CreateOrganizationRequest(v1CreateOrganizationRequest).Execute()

Create Organization

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
	v1CreateOrganizationRequest := *openapiclient.NewV1CreateOrganizationRequest("Name_example") // V1CreateOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.V1CreateOrganization(context.Background()).V1CreateOrganizationRequest(v1CreateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateOrganization`: V1CreateOrganization201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.V1CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateOrganizationRequest** | [**V1CreateOrganizationRequest**](V1CreateOrganizationRequest.md) |  | 

### Return type

[**V1CreateOrganization201Response**](V1CreateOrganization201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteOrganization

> V1DeleteOrganization(ctx, oID).Execute()

Delete Organization



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
	r, err := apiClient.OrganizationsAPI.V1DeleteOrganization(context.Background(), oID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetOrganization

> V1CreateOrganization201Response V1GetOrganization(ctx, oID).Execute()

Get Organization



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
	resp, r, err := apiClient.OrganizationsAPI.V1GetOrganization(context.Background(), oID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetOrganization`: V1CreateOrganization201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.V1GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1CreateOrganization201Response**](V1CreateOrganization201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListOrganizations

> V1ListOrganizations200Response V1ListOrganizations(ctx).Execute()

List Organizations

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.V1ListOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1ListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListOrganizations`: V1ListOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.V1ListOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListOrganizationsRequest struct via the builder pattern


### Return type

[**V1ListOrganizations200Response**](V1ListOrganizations200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationMemberDelete

> V1OrganizationMemberDelete(ctx, oID, uID).Execute()

Delete Organization Memeber



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
	uID := "iaI_KnKzSNxMhqwNKfvEuw" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.V1OrganizationMemberDelete(context.Background(), oID, uID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1OrganizationMemberDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**uID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationMemberDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationMemberUpdate

> V1OrganizationMemberUpdate200Response V1OrganizationMemberUpdate(ctx, oID, uID).V1OrganizationMemberUpdateRequest(v1OrganizationMemberUpdateRequest).Execute()

Update Organization Member

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
	uID := "iaI_KnKzSNxMhqwNKfvEuw" // string | User ID
	v1OrganizationMemberUpdateRequest := *openapiclient.NewV1OrganizationMemberUpdateRequest("Role_example") // V1OrganizationMemberUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.V1OrganizationMemberUpdate(context.Background(), oID, uID).V1OrganizationMemberUpdateRequest(v1OrganizationMemberUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1OrganizationMemberUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OrganizationMemberUpdate`: V1OrganizationMemberUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.V1OrganizationMemberUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**uID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationMemberUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1OrganizationMemberUpdateRequest** | [**V1OrganizationMemberUpdateRequest**](V1OrganizationMemberUpdateRequest.md) |  | 

### Return type

[**V1OrganizationMemberUpdate200Response**](V1OrganizationMemberUpdate200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateOrganization

> V1CreateOrganization201Response V1UpdateOrganization(ctx, oID).V1UpdateOrganizationRequest(v1UpdateOrganizationRequest).Execute()

Update Organization



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
	v1UpdateOrganizationRequest := *openapiclient.NewV1UpdateOrganizationRequest() // V1UpdateOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.V1UpdateOrganization(context.Background(), oID).V1UpdateOrganizationRequest(v1UpdateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.V1UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UpdateOrganization`: V1CreateOrganization201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.V1UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1UpdateOrganizationRequest** | [**V1UpdateOrganizationRequest**](V1UpdateOrganizationRequest.md) |  | 

### Return type

[**V1CreateOrganization201Response**](V1CreateOrganization201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

