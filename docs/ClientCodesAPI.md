# \ClientCodesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClientCodes**](ClientCodesAPI.md#ListClientCodes) | **Get** /client-codes/ | 



## ListClientCodes

> ClientCodesResp ListClientCodes(ctx).Limit(limit).Offset(offset).Type_(type_).Unused(unused).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	type_ := "type__example" // string | FC/iSCSI (optional)
	unused := true // bool | list unused client codes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientCodesAPI.ListClientCodes(context.Background()).Limit(limit).Offset(offset).Type_(type_).Unused(unused).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientCodesAPI.ListClientCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClientCodes`: ClientCodesResp
	fmt.Fprintf(os.Stdout, "Response from `ClientCodesAPI.ListClientCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClientCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **type_** | **string** | FC/iSCSI | 
 **unused** | **bool** | list unused client codes | 

### Return type

[**ClientCodesResp**](ClientCodesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

