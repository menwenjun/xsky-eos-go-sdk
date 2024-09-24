# \FcPortClientCodesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFCPortClientCodes**](FcPortClientCodesAPI.md#ListFCPortClientCodes) | **Get** /fc-port-client-codes/ | 



## ListFCPortClientCodes

> FCPortClientCodesResp ListFCPortClientCodes(ctx).Unused(unused).Limit(limit).Offset(offset).Execute()





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
	unused := true // bool | list unused client codes only (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FcPortClientCodesAPI.ListFCPortClientCodes(context.Background()).Unused(unused).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FcPortClientCodesAPI.ListFCPortClientCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFCPortClientCodes`: FCPortClientCodesResp
	fmt.Fprintf(os.Stdout, "Response from `FcPortClientCodesAPI.ListFCPortClientCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFCPortClientCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unused** | **bool** | list unused client codes only | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 

### Return type

[**FCPortClientCodesResp**](FCPortClientCodesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

