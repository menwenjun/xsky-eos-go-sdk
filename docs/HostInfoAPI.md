# \HostInfoAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostInfo**](HostInfoAPI.md#GetHostInfo) | **Get** /host-info/ | 



## GetHostInfo

> HostInfoResp GetHostInfo(ctx).Ip(ip).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/menwenjun/xsky-eos-go-sdk"
)

func main() {
	ip := "ip_example" // string | host ip

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostInfoAPI.GetHostInfo(context.Background()).Ip(ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostInfoAPI.GetHostInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostInfo`: HostInfoResp
	fmt.Fprintf(os.Stdout, "Response from `HostInfoAPI.GetHostInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | host ip | 

### Return type

[**HostInfoResp**](HostInfoResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

