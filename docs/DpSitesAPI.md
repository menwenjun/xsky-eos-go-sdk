# \DpSitesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpSite**](DpSitesAPI.md#CreateDpSite) | **Post** /dp-sites/ | 
[**DeleteDpSite**](DpSitesAPI.md#DeleteDpSite) | **Delete** /dp-sites/{site_id} | 
[**GetBackupBlockSnapshots**](DpSitesAPI.md#GetBackupBlockSnapshots) | **Get** /dp-sites/{site_id}/backup-block-snapshots | 
[**GetBackupBlockVolumes**](DpSitesAPI.md#GetBackupBlockVolumes) | **Get** /dp-sites/{site_id}/backup-block-volumes | 
[**GetBackupClusters**](DpSitesAPI.md#GetBackupClusters) | **Get** /dp-sites/{site_id}/backup-clusters | 
[**GetDpSite**](DpSitesAPI.md#GetDpSite) | **Get** /dp-sites/{site_id} | 
[**ListDpSites**](DpSitesAPI.md#ListDpSites) | **Get** /dp-sites/ | 
[**UpdateDpSite**](DpSitesAPI.md#UpdateDpSite) | **Patch** /dp-sites/{site_id} | 



## CreateDpSite

> DpSiteResp CreateDpSite(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDpSiteCreateReq() // DpSiteCreateReq | dp site info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.CreateDpSite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.CreateDpSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDpSite`: DpSiteResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.CreateDpSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DpSiteCreateReq**](DpSiteCreateReq.md) | dp site info | 

### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDpSite

> DeleteDpSite(ctx, siteId).Execute()





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
	siteId := int64(789) // int64 | dp site id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DpSitesAPI.DeleteDpSite(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.DeleteDpSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | dp site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDpSiteRequest struct via the builder pattern


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


## GetBackupBlockSnapshots

> DpBackupBlockSnapshotsResp GetBackupBlockSnapshots(ctx, siteId).DpGatewayId(dpGatewayId).ClusterFsId(clusterFsId).BlockVolumeName(blockVolumeName).After(after).Execute()





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
	siteId := int64(789) // int64 | dp site id
	dpGatewayId := int64(789) // int64 | the dp gateway to execute the query
	clusterFsId := "clusterFsId_example" // string | cluster fs id
	blockVolumeName := "blockVolumeName_example" // string | block volume name
	after := "after_example" // string | continuation token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.GetBackupBlockSnapshots(context.Background(), siteId).DpGatewayId(dpGatewayId).ClusterFsId(clusterFsId).BlockVolumeName(blockVolumeName).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.GetBackupBlockSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupBlockSnapshots`: DpBackupBlockSnapshotsResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.GetBackupBlockSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | dp site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupBlockSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dpGatewayId** | **int64** | the dp gateway to execute the query | 
 **clusterFsId** | **string** | cluster fs id | 
 **blockVolumeName** | **string** | block volume name | 
 **after** | **string** | continuation token | 

### Return type

[**DpBackupBlockSnapshotsResp**](DpBackupBlockSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupBlockVolumes

> DpBackupBlockVolumesResp GetBackupBlockVolumes(ctx, siteId).DpGatewayId(dpGatewayId).ClusterFsId(clusterFsId).After(after).Execute()





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
	siteId := int64(789) // int64 | dp site id
	dpGatewayId := int64(789) // int64 | the dp gateway to execute the query
	clusterFsId := "clusterFsId_example" // string | cluster fs id
	after := "after_example" // string | continuation token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.GetBackupBlockVolumes(context.Background(), siteId).DpGatewayId(dpGatewayId).ClusterFsId(clusterFsId).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.GetBackupBlockVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupBlockVolumes`: DpBackupBlockVolumesResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.GetBackupBlockVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | dp site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupBlockVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dpGatewayId** | **int64** | the dp gateway to execute the query | 
 **clusterFsId** | **string** | cluster fs id | 
 **after** | **string** | continuation token | 

### Return type

[**DpBackupBlockVolumesResp**](DpBackupBlockVolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupClusters

> DpBackupClustersResp GetBackupClusters(ctx, siteId).DpGatewayId(dpGatewayId).Execute()





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
	siteId := int64(789) // int64 | dp site id
	dpGatewayId := int64(789) // int64 | the dp gateway to execute the query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.GetBackupClusters(context.Background(), siteId).DpGatewayId(dpGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.GetBackupClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupClusters`: DpBackupClustersResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.GetBackupClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | dp site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dpGatewayId** | **int64** | the dp gateway to execute the query | 

### Return type

[**DpBackupClustersResp**](DpBackupClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDpSite

> DpSiteResp GetDpSite(ctx, siteId).Execute()





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
	siteId := int64(789) // int64 | protection site id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.GetDpSite(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.GetDpSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDpSite`: DpSiteResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.GetDpSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | protection site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDpSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpSites

> DpSitesResp ListDpSites(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ProtectionType(protectionType).Execute()





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
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)
	protectionType := "protectionType_example" // string | protection type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.ListDpSites(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ProtectionType(protectionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.ListDpSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDpSites`: DpSitesResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.ListDpSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **protectionType** | **string** | protection type | 

### Return type

[**DpSitesResp**](DpSitesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDpSite

> DpSiteResp UpdateDpSite(ctx, siteId).Body(body).Execute()





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
	siteId := int64(789) // int64 | dp site id
	body := *openapiclient.NewDpSiteUpdateReq() // DpSiteUpdateReq | dp site info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DpSitesAPI.UpdateDpSite(context.Background(), siteId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DpSitesAPI.UpdateDpSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDpSite`: DpSiteResp
	fmt.Fprintf(os.Stdout, "Response from `DpSitesAPI.UpdateDpSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **int64** | dp site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDpSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DpSiteUpdateReq**](DpSiteUpdateReq.md) | dp site info | 

### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

