# \OspMetadataServersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOspMetadataServer**](OspMetadataServersAPI.md#DeleteOspMetadataServer) | **Delete** /osp-metadata-servers/{id} | 
[**DisableOspMetadataServer**](OspMetadataServersAPI.md#DisableOspMetadataServer) | **Post** /osp-metadata-servers/{id}:disable | 
[**EnableOspMetadataServer**](OspMetadataServersAPI.md#EnableOspMetadataServer) | **Post** /osp-metadata-servers/{id}:enable | 
[**GetOspMetadataServer**](OspMetadataServersAPI.md#GetOspMetadataServer) | **Get** /osp-metadata-servers/{id} | 
[**GetOspMetadataServerSamples**](OspMetadataServersAPI.md#GetOspMetadataServerSamples) | **Get** /osp-metadata-servers/{id}/samples | 
[**ListOspMetadataServers**](OspMetadataServersAPI.md#ListOspMetadataServers) | **Get** /osp-metadata-servers/ | 
[**SetOspMetadataServerRole**](OspMetadataServersAPI.md#SetOspMetadataServerRole) | **Post** /osp-metadata-servers/{id}:set-role | 



## DeleteOspMetadataServer

> OspMetadataServerResp DeleteOspMetadataServer(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata server id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.DeleteOspMetadataServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.DeleteOspMetadataServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOspMetadataServer`: OspMetadataServerResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.DeleteOspMetadataServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOspMetadataServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataServerResp**](OspMetadataServerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableOspMetadataServer

> OspMetadataServerResp DisableOspMetadataServer(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata server id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.DisableOspMetadataServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.DisableOspMetadataServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableOspMetadataServer`: OspMetadataServerResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.DisableOspMetadataServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableOspMetadataServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataServerResp**](OspMetadataServerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOspMetadataServer

> OspMetadataServerResp EnableOspMetadataServer(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata server id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.EnableOspMetadataServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.EnableOspMetadataServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableOspMetadataServer`: OspMetadataServerResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.EnableOspMetadataServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableOspMetadataServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataServerResp**](OspMetadataServerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspMetadataServer

> OspMetadataServerResp GetOspMetadataServer(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata server id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.GetOspMetadataServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.GetOspMetadataServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspMetadataServer`: OspMetadataServerResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.GetOspMetadataServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspMetadataServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataServerResp**](OspMetadataServerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspMetadataServerSamples

> OspMetadataServerSamplesResp GetOspMetadataServerSamples(ctx, id).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	id := int64(789) // int64 | osp metadata server id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.GetOspMetadataServerSamples(context.Background(), id).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.GetOspMetadataServerSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspMetadataServerSamples`: OspMetadataServerSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.GetOspMetadataServerSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspMetadataServerSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OspMetadataServerSamplesResp**](OspMetadataServerSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOspMetadataServers

> OspMetadataServersResp ListOspMetadataServers(ctx).Limit(limit).Offset(offset).ClusterId(clusterId).OspMetadataClusterId(ospMetadataClusterId).Status(status).Execute()





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
	clusterId := int64(789) // int64 | The id of osp metadata server's cluster (optional)
	ospMetadataClusterId := int64(789) // int64 | The id of osp metadata server's metadata cluster (optional)
	status := "status_example" // string | status of osp metadata server (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.ListOspMetadataServers(context.Background()).Limit(limit).Offset(offset).ClusterId(clusterId).OspMetadataClusterId(ospMetadataClusterId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.ListOspMetadataServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOspMetadataServers`: OspMetadataServersResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.ListOspMetadataServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOspMetadataServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **clusterId** | **int64** | The id of osp metadata server&#39;s cluster | 
 **ospMetadataClusterId** | **int64** | The id of osp metadata server&#39;s metadata cluster | 
 **status** | **string** | status of osp metadata server | 

### Return type

[**OspMetadataServersResp**](OspMetadataServersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOspMetadataServerRole

> OspMetadataServerResp SetOspMetadataServerRole(ctx, id).Execute()





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
	id := int64(789) // int64 | osp metadata server id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspMetadataServersAPI.SetOspMetadataServerRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspMetadataServersAPI.SetOspMetadataServerRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOspMetadataServerRole`: OspMetadataServerResp
	fmt.Fprintf(os.Stdout, "Response from `OspMetadataServersAPI.SetOspMetadataServerRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | osp metadata server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOspMetadataServerRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OspMetadataServerResp**](OspMetadataServerResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

