# \AlertInfoRootAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckAlertInfoRoot**](AlertInfoRootAPI.md#AckAlertInfoRoot) | **Post** /alert-info-root/:ack | 
[**GetAlertInfoRoot**](AlertInfoRootAPI.md#GetAlertInfoRoot) | **Get** /alert-info-root/ | 
[**ResolveAlertInfoRoot**](AlertInfoRootAPI.md#ResolveAlertInfoRoot) | **Post** /alert-info-root/:resolve | 



## AckAlertInfoRoot

> AlertInfoRootResp AckAlertInfoRoot(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoRootAPI.AckAlertInfoRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoRootAPI.AckAlertInfoRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AckAlertInfoRoot`: AlertInfoRootResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoRootAPI.AckAlertInfoRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAckAlertInfoRootRequest struct via the builder pattern


### Return type

[**AlertInfoRootResp**](AlertInfoRootResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertInfoRoot

> AlertInfoRootResp GetAlertInfoRoot(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoRootAPI.GetAlertInfoRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoRootAPI.GetAlertInfoRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertInfoRoot`: AlertInfoRootResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoRootAPI.GetAlertInfoRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertInfoRootRequest struct via the builder pattern


### Return type

[**AlertInfoRootResp**](AlertInfoRootResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveAlertInfoRoot

> AlertInfoRootResp ResolveAlertInfoRoot(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertInfoRootAPI.ResolveAlertInfoRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertInfoRootAPI.ResolveAlertInfoRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAlertInfoRoot`: AlertInfoRootResp
	fmt.Fprintf(os.Stdout, "Response from `AlertInfoRootAPI.ResolveAlertInfoRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResolveAlertInfoRootRequest struct via the builder pattern


### Return type

[**AlertInfoRootResp**](AlertInfoRootResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

