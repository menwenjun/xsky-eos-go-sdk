# \LunsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLun**](LunsAPI.md#GetLun) | **Get** /luns/{lun_id} | 
[**ListLuns**](LunsAPI.md#ListLuns) | **Get** /luns/ | 



## GetLun

> LunResp GetLun(ctx, lunId).Execute()





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
	lunId := int64(789) // int64 | lun id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LunsAPI.GetLun(context.Background(), lunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LunsAPI.GetLun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLun`: LunResp
	fmt.Fprintf(os.Stdout, "Response from `LunsAPI.GetLun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lunId** | **int64** | lun id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LunResp**](LunResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLuns

> LunsResp ListLuns(ctx).Limit(limit).Offset(offset).MappingGroupId(mappingGroupId).VolumeId(volumeId).AccessPathId(accessPathId).ClusterId(clusterId).Execute()





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
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	mappingGroupId := int64(789) // int64 | mapping group id (optional)
	volumeId := int64(789) // int64 | volume id (optional)
	accessPathId := int64(789) // int64 | access path id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LunsAPI.ListLuns(context.Background()).Limit(limit).Offset(offset).MappingGroupId(mappingGroupId).VolumeId(volumeId).AccessPathId(accessPathId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LunsAPI.ListLuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLuns`: LunsResp
	fmt.Fprintf(os.Stdout, "Response from `LunsAPI.ListLuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **mappingGroupId** | **int64** | mapping group id | 
 **volumeId** | **int64** | volume id | 
 **accessPathId** | **int64** | access path id | 
 **clusterId** | **string** | cluster id | 

### Return type

[**LunsResp**](LunsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

