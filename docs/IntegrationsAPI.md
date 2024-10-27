# \IntegrationsAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DeleteIntegrations**](IntegrationsAPI.md#V1DeleteIntegrations) | **Delete** /v1/organizations/{oID}/projects/{pID}/integrations/{intID} | Delete Integration
[**V1ListIntegrations**](IntegrationsAPI.md#V1ListIntegrations) | **Get** /v1/organizations/{oID}/projects/{pID}/integrations | List Integrations



## V1DeleteIntegrations

> V1DeleteIntegrations(ctx, oID, pID, intID).Execute()

Delete Integration

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
	intID := "intID_example" // string | Integration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.V1DeleteIntegrations(context.Background(), oID, pID, intID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.V1DeleteIntegrations``: %v\n", err)
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
**intID** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteIntegrationsRequest struct via the builder pattern


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


## V1ListIntegrations

> V1ListIntegrations200Response V1ListIntegrations(ctx, oID, pID).Execute()

List Integrations

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
	resp, r, err := apiClient.IntegrationsAPI.V1ListIntegrations(context.Background(), oID, pID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.V1ListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListIntegrations`: V1ListIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.V1ListIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1ListIntegrations200Response**](V1ListIntegrations200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

