# \DfsQuotasAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDfsQuota**](DfsQuotasAPI.md#CreateDfsQuota) | **Post** /dfs-quotas/ | 
[**DeleteDfsQuota**](DfsQuotasAPI.md#DeleteDfsQuota) | **Delete** /dfs-quotas/{dfs_quota_id} | 
[**DfsQuotaOverview**](DfsQuotasAPI.md#DfsQuotaOverview) | **Get** /dfs-quotas/overview | 
[**GetDfsQuota**](DfsQuotasAPI.md#GetDfsQuota) | **Get** /dfs-quotas/{dfs_quota_id} | 
[**GetDfsQuotaPredictions**](DfsQuotasAPI.md#GetDfsQuotaPredictions) | **Get** /dfs-quotas/{dfs_quota_id}/predictions | 
[**GetDfsQuotaSamples**](DfsQuotasAPI.md#GetDfsQuotaSamples) | **Get** /dfs-quotas/{dfs_quota_id}/samples | 
[**ListDfsQuotas**](DfsQuotasAPI.md#ListDfsQuotas) | **Get** /dfs-quotas/ | 
[**PathValidator**](DfsQuotasAPI.md#PathValidator) | **Get** /dfs-quotas/path-validator | 
[**UpdateDfsQuota**](DfsQuotasAPI.md#UpdateDfsQuota) | **Patch** /dfs-quotas/{dfs_quota_id} | 



## CreateDfsQuota

> DfsQuotaResp CreateDfsQuota(ctx).Body(body).AllowPathCreate(allowPathCreate).IgnoreDirCheck(ignoreDirCheck).Execute()





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
	body := *openapiclient.NewDfsQuotaCreateReq(*openapiclient.NewDfsQuotaCreateReqQuota(int64(123), "Path_example")) // DfsQuotaCreateReq | quota info
	allowPathCreate := true // bool | allow create path when not existed (optional)
	ignoreDirCheck := true // bool | ignore directory check (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.CreateDfsQuota(context.Background()).Body(body).AllowPathCreate(allowPathCreate).IgnoreDirCheck(ignoreDirCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.CreateDfsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDfsQuota`: DfsQuotaResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.CreateDfsQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDfsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DfsQuotaCreateReq**](DfsQuotaCreateReq.md) | quota info | 
 **allowPathCreate** | **bool** | allow create path when not existed | 
 **ignoreDirCheck** | **bool** | ignore directory check | 

### Return type

[**DfsQuotaResp**](DfsQuotaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDfsQuota

> DfsQuotaResp DeleteDfsQuota(ctx, dfsQuotaId).Execute()





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
	dfsQuotaId := int64(789) // int64 | dfs quota id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.DeleteDfsQuota(context.Background(), dfsQuotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.DeleteDfsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDfsQuota`: DfsQuotaResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.DeleteDfsQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQuotaId** | **int64** | dfs quota id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDfsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQuotaResp**](DfsQuotaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DfsQuotaOverview

> DfsQuotaOverviewResp DfsQuotaOverview(ctx).Execute()





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
	resp, r, err := apiClient.DfsQuotasAPI.DfsQuotaOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.DfsQuotaOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DfsQuotaOverview`: DfsQuotaOverviewResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.DfsQuotaOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDfsQuotaOverviewRequest struct via the builder pattern


### Return type

[**DfsQuotaOverviewResp**](DfsQuotaOverviewResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQuota

> DfsQuotaResp GetDfsQuota(ctx, dfsQuotaId).Execute()





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
	dfsQuotaId := int64(789) // int64 | dfs quota id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.GetDfsQuota(context.Background(), dfsQuotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.GetDfsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQuota`: DfsQuotaResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.GetDfsQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQuotaId** | **int64** | dfs quota id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQuotaResp**](DfsQuotaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQuotaPredictions

> DfsQuotaPredictionsResp GetDfsQuotaPredictions(ctx, dfsQuotaId).Execute()





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
	dfsQuotaId := int64(789) // int64 | dfs quota id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.GetDfsQuotaPredictions(context.Background(), dfsQuotaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.GetDfsQuotaPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQuotaPredictions`: DfsQuotaPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.GetDfsQuotaPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQuotaId** | **int64** | dfs quota id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQuotaPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DfsQuotaPredictionsResp**](DfsQuotaPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDfsQuotaSamples

> DfsQuotaSamplesResp GetDfsQuotaSamples(ctx, dfsQuotaId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	dfsQuotaId := int64(789) // int64 | dfs quota id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.GetDfsQuotaSamples(context.Background(), dfsQuotaId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.GetDfsQuotaSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDfsQuotaSamples`: DfsQuotaSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.GetDfsQuotaSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQuotaId** | **int64** | dfs quota id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDfsQuotaSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**DfsQuotaSamplesResp**](DfsQuotaSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDfsQuotas

> DfsQuotasResp ListDfsQuotas(ctx).Path(path).Type_(type_).DomainUserName(domainUserName).DomainUserGroupName(domainUserGroupName).FsUserId(fsUserId).FsUserGroupId(fsUserGroupId).DfsGatewayGroupId(dfsGatewayGroupId).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()





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
	path := "path_example" // string | dfs quota path (optional)
	type_ := "type__example" // string | dfs quota type (optional)
	domainUserName := "domainUserName_example" // string | dfs quota domain user name (optional)
	domainUserGroupName := "domainUserGroupName_example" // string | dfs quota domain user group name (optional)
	fsUserId := int64(789) // int64 | fs user id (optional)
	fsUserGroupId := int64(789) // int64 | fs user group id (optional)
	dfsGatewayGroupId := int64(789) // int64 | dfs gateway group id (optional)
	limit := int64(789) // int64 | paging param (optional)
	offset := int64(789) // int64 | paging param (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.ListDfsQuotas(context.Background()).Path(path).Type_(type_).DomainUserName(domainUserName).DomainUserGroupName(domainUserGroupName).FsUserId(fsUserId).FsUserGroupId(fsUserGroupId).DfsGatewayGroupId(dfsGatewayGroupId).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.ListDfsQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDfsQuotas`: DfsQuotasResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.ListDfsQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDfsQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | dfs quota path | 
 **type_** | **string** | dfs quota type | 
 **domainUserName** | **string** | dfs quota domain user name | 
 **domainUserGroupName** | **string** | dfs quota domain user group name | 
 **fsUserId** | **int64** | fs user id | 
 **fsUserGroupId** | **int64** | fs user group id | 
 **dfsGatewayGroupId** | **int64** | dfs gateway group id | 
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 

### Return type

[**DfsQuotasResp**](DfsQuotasResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PathValidator

> DfsQuotaPathValidateResp PathValidator(ctx).DfsRootfsId(dfsRootfsId).Path(path).Execute()





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
	path := "path_example" // string | dfs quota path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.PathValidator(context.Background()).DfsRootfsId(dfsRootfsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.PathValidator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PathValidator`: DfsQuotaPathValidateResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.PathValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPathValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dfsRootfsId** | **int64** | dfs rootfs id | 
 **path** | **string** | dfs quota path | 

### Return type

[**DfsQuotaPathValidateResp**](DfsQuotaPathValidateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDfsQuota

> DfsQuotaResp UpdateDfsQuota(ctx, dfsQuotaId).Body(body).Execute()





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
	dfsQuotaId := int64(789) // int64 | quota id
	body := *openapiclient.NewDfsQuotaUpdateReq() // DfsQuotaUpdateReq | dfs quota info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DfsQuotasAPI.UpdateDfsQuota(context.Background(), dfsQuotaId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DfsQuotasAPI.UpdateDfsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDfsQuota`: DfsQuotaResp
	fmt.Fprintf(os.Stdout, "Response from `DfsQuotasAPI.UpdateDfsQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dfsQuotaId** | **int64** | quota id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDfsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DfsQuotaUpdateReq**](DfsQuotaUpdateReq.md) | dfs quota info | 

### Return type

[**DfsQuotaResp**](DfsQuotaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

