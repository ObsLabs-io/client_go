# \ChannelsAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateChannel**](ChannelsAPI.md#V1CreateChannel) | **Post** /v1/organizations/{oID}/projects/{pID}/channels | Create Channel
[**V1DeleteChannel**](ChannelsAPI.md#V1DeleteChannel) | **Delete** /v1/organizations/{oID}/projects/{pID}/channels/{chID} | Delete Channel
[**V1GetChannel**](ChannelsAPI.md#V1GetChannel) | **Get** /v1/organizations/{oID}/projects/{pID}/channels/{chID} | Get Channel
[**V1ListChannels**](ChannelsAPI.md#V1ListChannels) | **Get** /v1/organizations/{oID}/projects/{pID}/channels | List Channels
[**V1UpdateChannel**](ChannelsAPI.md#V1UpdateChannel) | **Patch** /v1/organizations/{oID}/projects/{pID}/channels/{chID} | Update Channel



## V1CreateChannel

> V1CreateChannel201Response V1CreateChannel(ctx, oID, pID).V1CreateChannelRequest(v1CreateChannelRequest).Execute()

Create Channel



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
	v1CreateChannelRequest := *openapiclient.NewV1CreateChannelRequest("Name_example") // V1CreateChannelRequest | Create Channel (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.V1CreateChannel(context.Background(), oID, pID).V1CreateChannelRequest(v1CreateChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.V1CreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateChannel`: V1CreateChannel201Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.V1CreateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1CreateChannelRequest** | [**V1CreateChannelRequest**](V1CreateChannelRequest.md) | Create Channel | 

### Return type

[**V1CreateChannel201Response**](V1CreateChannel201Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteChannel

> V1DeleteChannel(ctx, oID, pID, chID).Execute()

Delete Channel



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
	chID := "chID_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelsAPI.V1DeleteChannel(context.Background(), oID, pID, chID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.V1DeleteChannel``: %v\n", err)
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
**chID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteChannelRequest struct via the builder pattern


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


## V1GetChannel

> V1CreateChannel201Response V1GetChannel(ctx, oID, pID, chID).Execute()

Get Channel



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
	chID := "chID_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.V1GetChannel(context.Background(), oID, pID, chID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.V1GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetChannel`: V1CreateChannel201Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.V1GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**chID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1CreateChannel201Response**](V1CreateChannel201Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListChannels

> V1ListChannels200Response V1ListChannels(ctx, oID, pID).Execute()

List Channels



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
	resp, r, err := apiClient.ChannelsAPI.V1ListChannels(context.Background(), oID, pID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.V1ListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListChannels`: V1ListChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.V1ListChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1ListChannels200Response**](V1ListChannels200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateChannel

> V1CreateChannel201Response V1UpdateChannel(ctx, oID, pID, chID).V1UpdateChannelRequest(v1UpdateChannelRequest).Execute()

Update Channel



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
	chID := "chID_example" // string | Channel ID
	v1UpdateChannelRequest := *openapiclient.NewV1UpdateChannelRequest() // V1UpdateChannelRequest | Update Channel (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.V1UpdateChannel(context.Background(), oID, pID, chID).V1UpdateChannelRequest(v1UpdateChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.V1UpdateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UpdateChannel`: V1CreateChannel201Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.V1UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**chID** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1UpdateChannelRequest** | [**V1UpdateChannelRequest**](V1UpdateChannelRequest.md) | Update Channel | 

### Return type

[**V1CreateChannel201Response**](V1CreateChannel201Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

