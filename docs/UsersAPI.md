# \UsersAPI

All URIs are relative to *https://api.obslabs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MyUser**](UsersAPI.md#V1MyUser) | **Get** /v1/users/me | My User



## V1MyUser

> V1MyUser200Response V1MyUser(ctx).Execute()

My User

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
	resp, r, err := apiClient.UsersAPI.V1MyUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1MyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MyUser`: V1MyUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.V1MyUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MyUserRequest struct via the builder pattern


### Return type

[**V1MyUser200Response**](V1MyUser200Response.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

