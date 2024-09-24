# \ClientGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientGroup**](ClientGroupsAPI.md#CreateClientGroup) | **Post** /client-groups/ | 
[**DeleteClientGroup**](ClientGroupsAPI.md#DeleteClientGroup) | **Delete** /client-groups/{client_group_id} | 
[**GetClientGroup**](ClientGroupsAPI.md#GetClientGroup) | **Get** /client-groups/{client_group_id} | 
[**ListClientGroups**](ClientGroupsAPI.md#ListClientGroups) | **Get** /client-groups/ | 
[**UpdateClientGroup**](ClientGroupsAPI.md#UpdateClientGroup) | **Patch** /client-groups/{client_group_id} | 



## CreateClientGroup

> ClientGroupResp CreateClientGroup(ctx).Client(client).ClusterId(clusterId).Execute()





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
	client := *openapiclient.NewClientGroupCreateReq() // ClientGroupCreateReq | client group info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientGroupsAPI.CreateClientGroup(context.Background()).Client(client).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientGroupsAPI.CreateClientGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientGroup`: ClientGroupResp
	fmt.Fprintf(os.Stdout, "Response from `ClientGroupsAPI.CreateClientGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client** | [**ClientGroupCreateReq**](ClientGroupCreateReq.md) | client group info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientGroup

> DeleteClientGroup(ctx, clientGroupId).Execute()





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
	clientGroupId := int64(789) // int64 | client group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientGroupsAPI.DeleteClientGroup(context.Background(), clientGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientGroupsAPI.DeleteClientGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientGroupId** | **int64** | client group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientGroup

> ClientGroupResp GetClientGroup(ctx, clientGroupId).Execute()





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
	clientGroupId := int64(789) // int64 | client group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientGroupsAPI.GetClientGroup(context.Background(), clientGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientGroupsAPI.GetClientGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientGroup`: ClientGroupResp
	fmt.Fprintf(os.Stdout, "Response from `ClientGroupsAPI.GetClientGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientGroupId** | **int64** | client group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClientGroups

> ClientGroupsResp ListClientGroups(ctx).Limit(limit).Offset(offset).Uid(uid).BlockVolumeId(blockVolumeId).Q(q).Sort(sort).Execute()





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
	uid := "uid_example" // string | client group uid (optional)
	blockVolumeId := int64(789) // int64 | related block volume id (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientGroupsAPI.ListClientGroups(context.Background()).Limit(limit).Offset(offset).Uid(uid).BlockVolumeId(blockVolumeId).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientGroupsAPI.ListClientGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClientGroups`: ClientGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `ClientGroupsAPI.ListClientGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClientGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **uid** | **string** | client group uid | 
 **blockVolumeId** | **int64** | related block volume id | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**ClientGroupsResp**](ClientGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientGroup

> ClientGroupResp UpdateClientGroup(ctx, clientGroupId).Client(client).Execute()





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
	clientGroupId := int64(789) // int64 | client group id
	client := *openapiclient.NewClientGroupUpdateReq() // ClientGroupUpdateReq | client group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientGroupsAPI.UpdateClientGroup(context.Background(), clientGroupId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientGroupsAPI.UpdateClientGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientGroup`: ClientGroupResp
	fmt.Fprintf(os.Stdout, "Response from `ClientGroupsAPI.UpdateClientGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientGroupId** | **int64** | client group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | [**ClientGroupUpdateReq**](ClientGroupUpdateReq.md) | client group info | 

### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

