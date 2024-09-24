# \OsdsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOsd**](OsdsAPI.md#ActivateOsd) | **Post** /osds/{osd_id}:activate | 
[**CreateOsd**](OsdsAPI.md#CreateOsd) | **Post** /osds/ | 
[**DeleteOsd**](OsdsAPI.md#DeleteOsd) | **Delete** /osds/{osd_id} | 
[**GetOsd**](OsdsAPI.md#GetOsd) | **Get** /osds/{osd_id} | 
[**GetOsdPredictions**](OsdsAPI.md#GetOsdPredictions) | **Get** /osds/{osd_id}/predictions | 
[**GetOsdSamples**](OsdsAPI.md#GetOsdSamples) | **Get** /osds/{osd_id}/samples | 
[**GetOsdsOverview**](OsdsAPI.md#GetOsdsOverview) | **Get** /osds/overview | 
[**ListOsds**](OsdsAPI.md#ListOsds) | **Get** /osds/ | 
[**MaintainOsd**](OsdsAPI.md#MaintainOsd) | **Post** /osds/{osd_id}:maintain | 
[**RebuildOsd**](OsdsAPI.md#RebuildOsd) | **Post** /osds/{osd_id}:rebuild | 
[**SwitchOsdRole**](OsdsAPI.md#SwitchOsdRole) | **Post** /osds/{osd_id}:switch-role | 
[**UnmaintainOsd**](OsdsAPI.md#UnmaintainOsd) | **Post** /osds/{osd_id}:unmaintain | 
[**UnsetOsdIsolation**](OsdsAPI.md#UnsetOsdIsolation) | **Post** /osds/{osd_id}:unset-isolation | 
[**UpdateOsdNumaNode**](OsdsAPI.md#UpdateOsdNumaNode) | **Post** /osds/{osd_id}:update-numa-node | 



## ActivateOsd

> OsdResp ActivateOsd(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.ActivateOsd(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.ActivateOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.ActivateOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOsd

> OsdResp CreateOsd(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOsdCreateReq() // OsdCreateReq | osd info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.CreateOsd(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.CreateOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.CreateOsd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OsdCreateReq**](OsdCreateReq.md) | osd info | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOsd

> OsdResp DeleteOsd(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.DeleteOsd(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.DeleteOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.DeleteOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsd

> OsdResp GetOsd(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.GetOsd(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.GetOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.GetOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsdPredictions

> OsdPredictionsResp GetOsdPredictions(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.GetOsdPredictions(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.GetOsdPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsdPredictions`: OsdPredictionsResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.GetOsdPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdPredictionsResp**](OsdPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsdSamples

> OsdSamplesResp GetOsdSamples(ctx, osdId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	osdId := int64(789) // int64 | osd id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.GetOsdSamples(context.Background(), osdId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.GetOsdSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsdSamples`: OsdSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.GetOsdSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OsdSamplesResp**](OsdSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsdsOverview

> OsdOverviewResp GetOsdsOverview(ctx).ClusterId(clusterId).Execute()





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
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.GetOsdsOverview(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.GetOsdsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsdsOverview`: OsdOverviewResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.GetOsdsOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | cluster id | 

### Return type

[**OsdOverviewResp**](OsdOverviewResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOsds

> OsdsResp ListOsds(ctx).Limit(limit).Offset(offset).HostId(hostId).ClusterId(clusterId).DiskIds(diskIds).PoolId(poolId).BindPoolId(bindPoolId).OsdGroupId(osdGroupId).Type_(type_).Role(role).StatusClass(statusClass).WithCompound(withCompound).WithHybrid(withHybrid).CacheDiskId(cacheDiskId).Ids(ids).Q(q).Sort(sort).Execute()





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
	hostId := int64(789) // int64 | host id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	diskIds := int64(789) // int64 | disk ids (optional)
	poolId := int64(789) // int64 | pool id (optional)
	bindPoolId := int64(789) // int64 | bind pool id (optional)
	osdGroupId := int64(789) // int64 | osd group id (optional)
	type_ := "type__example" // string | osd type: HDD, SSD, Hybrid (optional)
	role := "role_example" // string | osd role: index or data (optional)
	statusClass := "statusClass_example" // string | osd status class: active, warning, error, offline, doing (optional)
	withCompound := true // bool | with compound osd (optional)
	withHybrid := true // bool | with hybrid osd (optional)
	cacheDiskId := int64(789) // int64 | cache disk id (optional)
	ids := "ids_example" // string | comma separate osd ids (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.ListOsds(context.Background()).Limit(limit).Offset(offset).HostId(hostId).ClusterId(clusterId).DiskIds(diskIds).PoolId(poolId).BindPoolId(bindPoolId).OsdGroupId(osdGroupId).Type_(type_).Role(role).StatusClass(statusClass).WithCompound(withCompound).WithHybrid(withHybrid).CacheDiskId(cacheDiskId).Ids(ids).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.ListOsds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOsds`: OsdsResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.ListOsds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOsdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **hostId** | **int64** | host id | 
 **clusterId** | **string** | cluster id | 
 **diskIds** | **int64** | disk ids | 
 **poolId** | **int64** | pool id | 
 **bindPoolId** | **int64** | bind pool id | 
 **osdGroupId** | **int64** | osd group id | 
 **type_** | **string** | osd type: HDD, SSD, Hybrid | 
 **role** | **string** | osd role: index or data | 
 **statusClass** | **string** | osd status class: active, warning, error, offline, doing | 
 **withCompound** | **bool** | with compound osd | 
 **withHybrid** | **bool** | with hybrid osd | 
 **cacheDiskId** | **int64** | cache disk id | 
 **ids** | **string** | comma separate osd ids | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**OsdsResp**](OsdsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintainOsd

> OsdResp MaintainOsd(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.MaintainOsd(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.MaintainOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintainOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.MaintainOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintainOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildOsd

> OsdResp RebuildOsd(ctx, osdId).Body(body).Execute()





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
	osdId := int64(789) // int64 | osd id
	body := *openapiclient.NewOsdRebuildReq() // OsdRebuildReq | osd info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.RebuildOsd(context.Background(), osdId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.RebuildOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.RebuildOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdRebuildReq**](OsdRebuildReq.md) | osd info | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchOsdRole

> OsdResp SwitchOsdRole(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.SwitchOsdRole(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.SwitchOsdRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwitchOsdRole`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.SwitchOsdRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchOsdRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmaintainOsd

> OsdResp UnmaintainOsd(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.UnmaintainOsd(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.UnmaintainOsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmaintainOsd`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.UnmaintainOsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmaintainOsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetOsdIsolation

> OsdResp UnsetOsdIsolation(ctx, osdId).Execute()





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
	osdId := int64(789) // int64 | osd id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.UnsetOsdIsolation(context.Background(), osdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.UnsetOsdIsolation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetOsdIsolation`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.UnsetOsdIsolation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetOsdIsolationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOsdNumaNode

> OsdResp UpdateOsdNumaNode(ctx, osdId).Body(body).Execute()





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
	osdId := int64(789) // int64 | osd id
	body := *openapiclient.NewOsdUpdateNumaNodeReq() // OsdUpdateNumaNodeReq | osd numa node

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdsAPI.UpdateOsdNumaNode(context.Background(), osdId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdsAPI.UpdateOsdNumaNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOsdNumaNode`: OsdResp
	fmt.Fprintf(os.Stdout, "Response from `OsdsAPI.UpdateOsdNumaNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osdId** | **int64** | osd id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOsdNumaNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdUpdateNumaNodeReq**](OsdUpdateNumaNodeReq.md) | osd numa node | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

