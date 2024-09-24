# \DfsFilesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Chmod**](DfsFilesAPI.md#Chmod) | **Post** /dfs-files/:chmod | 
[**DeleteDfsFile**](DfsFilesAPI.md#DeleteDfsFile) | **Post** /dfs-files:delete | 
[**DownloadDfsFile**](DfsFilesAPI.md#DownloadDfsFile) | **Get** /dfs-files/:download | 
[**FindDfsFiles**](DfsFilesAPI.md#FindDfsFiles) | **Get** /dfs-files/:find | 
[**GetDfsLogReport**](DfsFilesAPI.md#GetDfsLogReport) | **Get** /dfs-files/log-report | 
[**ListDfsFiles**](DfsFilesAPI.md#ListDfsFiles) | **Post** /dfs-files/ | 
[**ListDfsFilesWithResources**](DfsFilesAPI.md#ListDfsFilesWithResources) | **Get** /dfs-files/:list | 
[**RenameDfsFile**](DfsFilesAPI.md#RenameDfsFile) | **Post** /dfs-files/:rename | 
[**StatDfsFile**](DfsFilesAPI.md#StatDfsFile) | **Post** /dfs-files/:stat | 



## Chmod

> Chmod(ctx).Body(body).Execute()





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
	body := *openapiclient.NewChmodReq(*openapiclient.NewChmodReqFile(int64(123), "Mode_example", "Path_example")) // ChmodReq | path and permission info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DfsFilesAPI.Chmod(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.Chmod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChmodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChmodReq**](ChmodReq.md) | path and permission info | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsFile

> DeleteDfsFile(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDfsFileDeleteReq(*openapiclient.NewDfsFileDeleteReqFile(int64(123), "Path_example", "PrivilegedToken_example")) // DfsFileDeleteReq | active pool ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DfsFilesAPI.DeleteDfsFile(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.DeleteDfsFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsFileDeleteReq**](DfsFileDeleteReq.md) | active pool ids | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDfsFile

> *os.File DownloadDfsFile(ctx).Path(path).RootfsId(rootfsId).Execute()





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
	path := "path_example" // string | dfs file path
	rootfsId := int64(789) // int64 | dfs rootfs id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.DownloadDfsFile(context.Background()).Path(path).RootfsId(rootfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.DownloadDfsFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDfsFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.DownloadDfsFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDfsFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | dfs file path | 
 **rootfsId** | **int64** | dfs rootfs id | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDfsFiles

> DfsFileBasesResp FindDfsFiles(ctx).Path(path).RootfsId(rootfsId).Type_(type_).Execute()





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
	path := "path_example" // string | dfs dirctory path
	rootfsId := int64(789) // int64 | dfs rootfs id (optional)
	type_ := "type__example" // string | filter type, 'file' or 'dir' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.FindDfsFiles(context.Background()).Path(path).RootfsId(rootfsId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.FindDfsFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindDfsFiles`: DfsFileBasesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.FindDfsFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindDfsFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | dfs dirctory path | 
 **rootfsId** | **int64** | dfs rootfs id | 
 **type_** | **string** | filter type, &#39;file&#39; or &#39;dir&#39; | 

### Return type

[**DfsFileBasesResp**](DfsFileBasesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsLogReport

> *os.File GetDfsLogReport(ctx).DfsRootfsId(dfsRootfsId).Paths(paths).Execute()





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
	dfsRootfsId := int64(789) // int64 | dfs rootfs id
	paths := "paths_example" // string | dfs log path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.GetDfsLogReport(context.Background()).DfsRootfsId(dfsRootfsId).Paths(paths).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.GetDfsLogReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsLogReport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.GetDfsLogReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsLogReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **paths** | **string** | dfs log path | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsFiles

> DfsFilesResp ListDfsFiles(ctx).RootfsId(rootfsId).Limit(limit).Path(path).Start(start).Prefix(prefix).Type_(type_).Worm(worm).Reverse(reverse).PageUp(pageUp).Hidden(hidden).Execute()





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
	rootfsId := int64(789) // int64 | dfs rootfs id (optional)
	limit := int64(789) // int64 | paging param (optional)
	path := "path_example" // string | parent path (optional)
	start := "start_example" // string | start file for list (optional)
	prefix := "prefix_example" // string | prefix to filter file or directory (optional)
	type_ := "type__example" // string | filter by file or dir type (optional)
	worm := true // bool | check dir worm before enable filter (optional)
	reverse := true // bool | sort reverse (optional)
	pageUp := true // bool | query last page instead of next (optional)
	hidden := true // bool | show hidden files (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.ListDfsFiles(context.Background()).RootfsId(rootfsId).Limit(limit).Path(path).Start(start).Prefix(prefix).Type_(type_).Worm(worm).Reverse(reverse).PageUp(pageUp).Hidden(hidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.ListDfsFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsFiles`: DfsFilesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.ListDfsFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootfsId** | **int64** | dfs rootfs id | 
 **limit** | **int64** | paging param | 
 **path** | **string** | parent path | 
 **start** | **string** | start file for list | 
 **prefix** | **string** | prefix to filter file or directory | 
 **type_** | **string** | filter by file or dir type | 
 **worm** | **bool** | check dir worm before enable filter | 
 **reverse** | **bool** | sort reverse | 
 **pageUp** | **bool** | query last page instead of next | 
 **hidden** | **bool** | show hidden files | 

### Return type

[**DfsFilesResp**](DfsFilesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsFilesWithResources

> DfsListFilesWithResourcesResp ListDfsFilesWithResources(ctx).ResourceTypes(resourceTypes).RootfsId(rootfsId).Keyword(keyword).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()





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
	resourceTypes := "resourceTypes_example" // string | resource types
	rootfsId := int64(789) // int64 | dfs rootfs id (optional)
	keyword := "keyword_example" // string | keyword (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.ListDfsFilesWithResources(context.Background()).ResourceTypes(resourceTypes).RootfsId(rootfsId).Keyword(keyword).Limit(limit).Offset(offset).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.ListDfsFilesWithResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsFilesWithResources`: DfsListFilesWithResourcesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.ListDfsFilesWithResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsFilesWithResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceTypes** | **string** | resource types | 
 **rootfsId** | **int64** | dfs rootfs id | 
 **keyword** | **string** | keyword | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsListFilesWithResourcesResp**](DfsListFilesWithResourcesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameDfsFile

> RenameDfsFile(ctx).Body(body).Execute()





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
	body := *openapiclient.NewRenameDfsFileReq(*openapiclient.NewRenameDfsFileReqFile(int64(123), "DstPath_example", "SrcPath_example")) // RenameDfsFileReq | src path and dst path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DfsFilesAPI.RenameDfsFile(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.RenameDfsFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameDfsFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RenameDfsFileReq**](RenameDfsFileReq.md) | src path and dst path | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatDfsFile

> DfsFileResp StatDfsFile(ctx).RootfsId(rootfsId).Path(path).Execute()





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
	rootfsId := int64(789) // int64 | dfs rootfs id (optional)
	path := "path_example" // string | parent path (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsFilesAPI.StatDfsFile(context.Background()).RootfsId(rootfsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsFilesAPI.StatDfsFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatDfsFile`: DfsFileResp
	fmt.Fprintf(os.Stdout, "Response from `DfsFilesAPI.StatDfsFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatDfsFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootfsId** | **int64** | dfs rootfs id | 
 **path** | **string** | parent path | 

### Return type

[**DfsFileResp**](DfsFileResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

