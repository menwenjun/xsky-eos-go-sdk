# \DfsTrashesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseDfsTrash**](DfsTrashesAPI.md#CloseDfsTrash) | **Delete** /dfs-trashes/{dfs_trash_id} | 
[**EmptyDfsTrash**](DfsTrashesAPI.md#EmptyDfsTrash) | **Post** /dfs-trashes/{dfs_trash_id}:empty | 
[**GetDfsTrash**](DfsTrashesAPI.md#GetDfsTrash) | **Get** /dfs-trashes/{dfs_trash_id} | 
[**ListDfsTrashFileJobs**](DfsTrashesAPI.md#ListDfsTrashFileJobs) | **Get** /dfs-trashes/:list_trash_file_jobs | 
[**ListDfsTrashFiles**](DfsTrashesAPI.md#ListDfsTrashFiles) | **Post** /dfs-trashes/{dfs_trash_id}:list_trash_files | 
[**ListDfsTrashes**](DfsTrashesAPI.md#ListDfsTrashes) | **Get** /dfs-trashes/ | 
[**OpenDfsTrash**](DfsTrashesAPI.md#OpenDfsTrash) | **Post** /dfs-trashes/ | 
[**RemoveDfsTrashFile**](DfsTrashesAPI.md#RemoveDfsTrashFile) | **Post** /dfs-trashes/{dfs_trash_id}:remove-file | 
[**RestoreDfsTrashFile**](DfsTrashesAPI.md#RestoreDfsTrashFile) | **Post** /dfs-trashes/{dfs_trash_id}:restore-file | 
[**SearchDfsTrashFiles**](DfsTrashesAPI.md#SearchDfsTrashFiles) | **Post** /dfs-trashes/{dfs_trash_id}:search_trash_files | 
[**UpdateDfsTrash**](DfsTrashesAPI.md#UpdateDfsTrash) | **Patch** /dfs-trashes/{dfs_trash_id} | 



## CloseDfsTrash

> DfsTrashResp CloseDfsTrash(ctx, dfsTrashId).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.CloseDfsTrash(context.Background(), dfsTrashId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.CloseDfsTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseDfsTrash`: DfsTrashResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.CloseDfsTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseDfsTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsTrashResp**](DfsTrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmptyDfsTrash

> DfsTrashResp EmptyDfsTrash(ctx, dfsTrashId).Force(force).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	force := true // bool | force empty (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.EmptyDfsTrash(context.Background(), dfsTrashId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.EmptyDfsTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmptyDfsTrash`: DfsTrashResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.EmptyDfsTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmptyDfsTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force empty | 

### Return type

[**DfsTrashResp**](DfsTrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsTrash

> DfsTrashResp GetDfsTrash(ctx, dfsTrashId).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.GetDfsTrash(context.Background(), dfsTrashId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.GetDfsTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsTrash`: DfsTrashResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.GetDfsTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsTrashResp**](DfsTrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsTrashFileJobs

> DfsTrashFileJobResp ListDfsTrashFileJobs(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).DfsTrashId(dfsTrashId).Path(path).Action(action).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)
	dfsTrashId := "dfsTrashId_example" // string | trash id (optional)
	path := "path_example" // string | path (optional)
	action := "action_example" // string | action (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.ListDfsTrashFileJobs(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).DfsTrashId(dfsTrashId).Path(path).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.ListDfsTrashFileJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsTrashFileJobs`: DfsTrashFileJobResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.ListDfsTrashFileJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsTrashFileJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 
 **dfsTrashId** | **string** | trash id | 
 **path** | **string** | path | 
 **action** | **string** | action | 

### Return type

[**DfsTrashFileJobResp**](DfsTrashFileJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsTrashFiles

> DfsTrashFilesResp ListDfsTrashFiles(ctx, dfsTrashId).Limit(limit).Path(path).Start(start).Prefix(prefix).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	limit := int64(789) // int64 | paging param (optional)
	path := "path_example" // string | parent path (optional)
	start := "start_example" // string | start file for list (optional)
	prefix := "prefix_example" // string | prefix to filter file or directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.ListDfsTrashFiles(context.Background(), dfsTrashId).Limit(limit).Path(path).Start(start).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.ListDfsTrashFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsTrashFiles`: DfsTrashFilesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.ListDfsTrashFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDfsTrashFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **path** | **string** | parent path | 
 **start** | **string** | start file for list | 
 **prefix** | **string** | prefix to filter file or directory | 

### Return type

[**DfsTrashFilesResp**](DfsTrashFilesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsTrashes

> DfsTrashesResp ListDfsTrashes(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Status(status).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	status := "status_example" // string | status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.ListDfsTrashes(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.ListDfsTrashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsTrashes`: DfsTrashesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.ListDfsTrashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsTrashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **status** | **string** | status | 

### Return type

[**DfsTrashesResp**](DfsTrashesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenDfsTrash

> DfsTrashResp OpenDfsTrash(ctx).Body(body).AllowPathCreate(allowPathCreate).Execute()





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
	body := *openapiclient.NewDfsTrashOpenReq(*openapiclient.NewDfsTrashOpenReqTrash(int64(123), int64(123), "Path_example")) // DfsTrashOpenReq | trash info
	allowPathCreate := true // bool | allow create path when not existed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.OpenDfsTrash(context.Background()).Body(body).AllowPathCreate(allowPathCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.OpenDfsTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenDfsTrash`: DfsTrashResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.OpenDfsTrash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenDfsTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsTrashOpenReq**](DfsTrashOpenReq.md) | trash info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 

### Return type

[**DfsTrashResp**](DfsTrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDfsTrashFile

> DfsTrashFileJobResp RemoveDfsTrashFile(ctx, dfsTrashId).Body(body).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	body := *openapiclient.NewDfsTrashRemoveFileReq(*openapiclient.NewDfsTrashRemoveFileReqTrash("Path_example")) // DfsTrashRemoveFileReq | trash info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.RemoveDfsTrashFile(context.Background(), dfsTrashId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.RemoveDfsTrashFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDfsTrashFile`: DfsTrashFileJobResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.RemoveDfsTrashFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDfsTrashFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsTrashRemoveFileReq**](DfsTrashRemoveFileReq.md) | trash info | 

### Return type

[**DfsTrashFileJobResp**](DfsTrashFileJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDfsTrashFile

> DfsTrashFileJobResp RestoreDfsTrashFile(ctx, dfsTrashId).Body(body).Force(force).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	body := *openapiclient.NewDfsTrashRestoreFileReq(*openapiclient.NewDfsTrashRestoreFileReqTrash("Path_example")) // DfsTrashRestoreFileReq | trash info
	force := true // bool | force restore (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.RestoreDfsTrashFile(context.Background(), dfsTrashId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.RestoreDfsTrashFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreDfsTrashFile`: DfsTrashFileJobResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.RestoreDfsTrashFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDfsTrashFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsTrashRestoreFileReq**](DfsTrashRestoreFileReq.md) | trash info | 
 **force** | **bool** | force restore | 

### Return type

[**DfsTrashFileJobResp**](DfsTrashFileJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDfsTrashFiles

> DfsTrashFilesResp SearchDfsTrashFiles(ctx, dfsTrashId).Limit(limit).Path(path).Start(start).Prefix(prefix).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	limit := int64(789) // int64 | paging param (optional)
	path := "path_example" // string | subdirectory of trash path (optional)
	start := "start_example" // string | start file for list (optional)
	prefix := "prefix_example" // string | prefix to filter file or directory (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.SearchDfsTrashFiles(context.Background(), dfsTrashId).Limit(limit).Path(path).Start(start).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.SearchDfsTrashFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDfsTrashFiles`: DfsTrashFilesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.SearchDfsTrashFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDfsTrashFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | paging param | 
 **path** | **string** | subdirectory of trash path | 
 **start** | **string** | start file for list | 
 **prefix** | **string** | prefix to filter file or directory | 

### Return type

[**DfsTrashFilesResp**](DfsTrashFilesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsTrash

> DfsTrashResp UpdateDfsTrash(ctx, dfsTrashId).Body(body).Execute()





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
	dfsTrashId := int64(789) // int64 | trash id
	body := *openapiclient.NewDfsTrashUpdateReq(*openapiclient.NewDfsTrashUpdateReqTrash(int64(123))) // DfsTrashUpdateReq | trash info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsTrashesAPI.UpdateDfsTrash(context.Background(), dfsTrashId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsTrashesAPI.UpdateDfsTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsTrash`: DfsTrashResp
	fmt.Fprintf(os.Stdout, "Response from `DfsTrashesAPI.UpdateDfsTrash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsTrashId** | **int64** | trash id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsTrashUpdateReq**](DfsTrashUpdateReq.md) | trash info | 

### Return type

[**DfsTrashResp**](DfsTrashResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

