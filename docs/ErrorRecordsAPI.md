# \ErrorRecordsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateErrorRecord**](ErrorRecordsAPI.md#CreateErrorRecord) | **Post** /error-records/ | 



## CreateErrorRecord

> ErrorRecordResp CreateErrorRecord(ctx).Body(body).Execute()





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
	body := *openapiclient.NewErrorRecordCreateReq() // ErrorRecordCreateReq | error record info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ErrorRecordsAPI.CreateErrorRecord(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErrorRecordsAPI.CreateErrorRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateErrorRecord`: ErrorRecordResp
	fmt.Fprintf(os.Stdout, "Response from `ErrorRecordsAPI.CreateErrorRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateErrorRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ErrorRecordCreateReq**](ErrorRecordCreateReq.md) | error record info | 

### Return type

[**ErrorRecordResp**](ErrorRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

