# \ProbesAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateProbe**](ProbesAPI.md#V1CreateProbe) | **Post** /v1/organizations/{oID}/projects/{pID}/probes | Create Probe
[**V1DeleteProbe**](ProbesAPI.md#V1DeleteProbe) | **Delete** /v1/organizations/{oID}/projects/{pID}/probes/{prID} | Delete Probe
[**V1GetProbe**](ProbesAPI.md#V1GetProbe) | **Get** /v1/organizations/{oID}/projects/{pID}/probes/{prID} | Get Probe
[**V1ListProbes**](ProbesAPI.md#V1ListProbes) | **Get** /v1/organizations/{oID}/projects/{pID}/probes | List Probes
[**V1UpdateProbe**](ProbesAPI.md#V1UpdateProbe) | **Patch** /v1/organizations/{oID}/projects/{pID}/probes/{prID} | Update Probe



## V1CreateProbe

> V1CreateProbe201Response V1CreateProbe(ctx, oID, pID).V1CreateProbeRequest(v1CreateProbeRequest).Execute()

Create Probe



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
	v1CreateProbeRequest := *openapiclient.NewV1CreateProbeRequest("Name_example", "Type_example") // V1CreateProbeRequest | Create Probe (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProbesAPI.V1CreateProbe(context.Background(), oID, pID).V1CreateProbeRequest(v1CreateProbeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.V1CreateProbe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateProbe`: V1CreateProbe201Response
	fmt.Fprintf(os.Stdout, "Response from `ProbesAPI.V1CreateProbe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateProbeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v1CreateProbeRequest** | [**V1CreateProbeRequest**](V1CreateProbeRequest.md) | Create Probe | 

### Return type

[**V1CreateProbe201Response**](V1CreateProbe201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteProbe

> V1DeleteProbe(ctx, oID, pID, prID).Execute()

Delete Probe

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
	prID := "ECD3gsVktn1E_xuPYMlrig" // string | Probe ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProbesAPI.V1DeleteProbe(context.Background(), oID, pID, prID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.V1DeleteProbe``: %v\n", err)
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
**prID** | **string** | Probe ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeleteProbeRequest struct via the builder pattern


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


## V1GetProbe

> V1CreateProbe201Response V1GetProbe(ctx, oID, pID, prID).Execute()

Get Probe

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
	prID := "ECD3gsVktn1E_xuPYMlrig" // string | Probe ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProbesAPI.V1GetProbe(context.Background(), oID, pID, prID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.V1GetProbe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetProbe`: V1CreateProbe201Response
	fmt.Fprintf(os.Stdout, "Response from `ProbesAPI.V1GetProbe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**prID** | **string** | Probe ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetProbeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V1CreateProbe201Response**](V1CreateProbe201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListProbes

> V1ListProbes200Response V1ListProbes(ctx, oID, pID).Execute()

List Probes



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
	resp, r, err := apiClient.ProbesAPI.V1ListProbes(context.Background(), oID, pID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.V1ListProbes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListProbes`: V1ListProbes200Response
	fmt.Fprintf(os.Stdout, "Response from `ProbesAPI.V1ListProbes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ListProbesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V1ListProbes200Response**](V1ListProbes200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateProbe

> V1CreateProbe201Response V1UpdateProbe(ctx, oID, pID, prID).V1UpdateProbeRequest(v1UpdateProbeRequest).Execute()

Update Probe

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
	prID := "ECD3gsVktn1E_xuPYMlrig" // string | Probe ID
	v1UpdateProbeRequest := *openapiclient.NewV1UpdateProbeRequest() // V1UpdateProbeRequest | Update Probe (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProbesAPI.V1UpdateProbe(context.Background(), oID, pID, prID).V1UpdateProbeRequest(v1UpdateProbeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.V1UpdateProbe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UpdateProbe`: V1CreateProbe201Response
	fmt.Fprintf(os.Stdout, "Response from `ProbesAPI.V1UpdateProbe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oID** | **string** | Organization ID | 
**pID** | **string** | Project ID | 
**prID** | **string** | Probe ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UpdateProbeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v1UpdateProbeRequest** | [**V1UpdateProbeRequest**](V1UpdateProbeRequest.md) | Update Probe | 

### Return type

[**V1CreateProbe201Response**](V1CreateProbe201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

