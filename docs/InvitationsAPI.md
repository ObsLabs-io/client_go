# \InvitationsAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AcceptInvitation**](InvitationsAPI.md#V1AcceptInvitation) | **Post** /v1/organizations/{oID}/invitations/{invID}/accept | Accept Invitation
[**V1GetInvitation**](InvitationsAPI.md#V1GetInvitation) | **Get** /v1/organizations/{oID}/invitations/{invID} | Get Invitation
[**V1ListInvitations**](InvitationsAPI.md#V1ListInvitations) | **Get** /v1/organizations/{oID}/invitations | List Invitations
[**V1ResendInvitation**](InvitationsAPI.md#V1ResendInvitation) | **Post** /v1/organizations/{oID}/invitations/{invID}/resend | Resend Invitation
[**V1RevokeInvitation**](InvitationsAPI.md#V1RevokeInvitation) | **Delete** /v1/organizations/{oID}/invitations/{invID} | Revoke Invitation
[**V1SendInvitation**](InvitationsAPI.md#V1SendInvitation) | **Post** /v1/organizations/{oID}/invitations | Send Invitation



## V1AcceptInvitation

> V1AcceptInvitation(ctx, oID, invID).V1AcceptInvitationRequest(v1AcceptInvitationRequest).Execute()

Accept Invitation

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
	invID := "2pg7EQvDIFjREnMBuRQ30Q" // string | Invitation ID
	v1AcceptInvitationRequest := *openapiclient.NewV1AcceptInvitationRequest("SecretKey_example") // V1AcceptInvitationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvitationsAPI.V1AcceptInvitation(context.Background(), oID, invID).V1AcceptInvitationRequest(v1AcceptInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1AcceptInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**invID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1AcceptInvitationRequest** | [**V1AcceptInvitationRequest**](V1AcceptInvitationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetInvitation

> V1SendInvitation201Response V1GetInvitation(ctx, oID, invID).Execute()

Get Invitation

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
	invID := "2pg7EQvDIFjREnMBuRQ30Q" // string | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.V1GetInvitation(context.Background(), oID, invID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1GetInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetInvitation`: V1SendInvitation201Response
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.V1GetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**invID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1SendInvitation201Response**](V1SendInvitation201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListInvitations

> V1ListInvitations200Response V1ListInvitations(ctx, oID).Execute()

List Invitations



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
	resp, r, err := apiClient.InvitationsAPI.V1ListInvitations(context.Background(), oID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1ListInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListInvitations`: V1ListInvitations200Response
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.V1ListInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1ListInvitations200Response**](V1ListInvitations200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResendInvitation

> V1SendInvitation201Response V1ResendInvitation(ctx, oID, invID).Execute()

Resend Invitation



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
	invID := "2pg7EQvDIFjREnMBuRQ30Q" // string | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.V1ResendInvitation(context.Background(), oID, invID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1ResendInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResendInvitation`: V1SendInvitation201Response
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.V1ResendInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**invID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ResendInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1SendInvitation201Response**](V1SendInvitation201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1RevokeInvitation

> V1RevokeInvitation(ctx, oID, invID).Execute()

Revoke Invitation

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
	invID := "2pg7EQvDIFjREnMBuRQ30Q" // string | Invitation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvitationsAPI.V1RevokeInvitation(context.Background(), oID, invID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1RevokeInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**invID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1RevokeInvitationRequest struct via the builder pattern


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


## V1SendInvitation

> V1SendInvitation201Response V1SendInvitation(ctx, oID).V1SendInvitationRequest(v1SendInvitationRequest).Execute()

Send Invitation



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
	v1SendInvitationRequest := *openapiclient.NewV1SendInvitationRequest("Email_example", "Role_example") // V1SendInvitationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.V1SendInvitation(context.Background(), oID).V1SendInvitationRequest(v1SendInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.V1SendInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SendInvitation`: V1SendInvitation201Response
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.V1SendInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SendInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v1SendInvitationRequest** | [**V1SendInvitationRequest**](V1SendInvitationRequest.md) |  | 

### Return type

[**V1SendInvitation201Response**](V1SendInvitation201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

