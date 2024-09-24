# \HostsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHost**](HostsAPI.md#CreateHost) | **Post** /hosts/ | 
[**DeleteHost**](HostsAPI.md#DeleteHost) | **Delete** /hosts/{host_id} | 
[**DeleteHostsGatewayLbGroup**](HostsAPI.md#DeleteHostsGatewayLbGroup) | **Post** /hosts/:delete-hosts-gateway-lb-group | 
[**GetHost**](HostsAPI.md#GetHost) | **Get** /hosts/{host_id} | 
[**GetHostSamples**](HostsAPI.md#GetHostSamples) | **Get** /hosts/{host_id}/samples | 
[**HostDeletable**](HostsAPI.md#HostDeletable) | **Get** /hosts/{host_id}:host-deletable | 
[**ListHosts**](HostsAPI.md#ListHosts) | **Get** /hosts/ | 
[**MaintainHost**](HostsAPI.md#MaintainHost) | **Post** /hosts/{host_id}:maintain | 
[**RemoveHostsFormOspZone**](HostsAPI.md#RemoveHostsFormOspZone) | **Post** /hosts/:remove-hosts-from-osp-zone | 
[**SetHostsToOspZone**](HostsAPI.md#SetHostsToOspZone) | **Post** /hosts/:add-hosts-to-osp-zone | 
[**UnmaintainHost**](HostsAPI.md#UnmaintainHost) | **Post** /hosts/{host_id}:unmaintain | 
[**UpdateHost**](HostsAPI.md#UpdateHost) | **Patch** /hosts/{host_id} | 



## CreateHost

> HostResp CreateHost(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewHostCreateReq() // HostCreateReq | host info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.CreateHost(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.CreateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.CreateHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HostCreateReq**](HostCreateReq.md) | host info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHost

> HostResp DeleteHost(ctx, hostId).Force(force).Execute()





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
	hostId := int64(789) // int64 | host id
	force := true // bool | force delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.DeleteHost(context.Background(), hostId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.DeleteHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.DeleteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force delete or not | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostsGatewayLbGroup

> DeleteHostsGatewayLbGroup(ctx).Body(body).Execute()





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
	body := *openapiclient.NewDeleteHostsGatewayLbGroupOpReq() // DeleteHostsGatewayLbGroupOpReq | osp zone info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.DeleteHostsGatewayLbGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.DeleteHostsGatewayLbGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostsGatewayLbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteHostsGatewayLbGroupOpReq**](DeleteHostsGatewayLbGroupOpReq.md) | osp zone info | 

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


## GetHost

> HostResp GetHost(ctx, hostId).Execute()





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
	hostId := int64(789) // int64 | the host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.GetHost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | the host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostSamples

> HostSamplesResp GetHostSamples(ctx, hostId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	hostId := int64(789) // int64 | host id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.GetHostSamples(context.Background(), hostId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.GetHostSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostSamples`: HostSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.GetHostSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**HostSamplesResp**](HostSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostDeletable

> HostResp HostDeletable(ctx, hostId).Execute()





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
	hostId := int64(789) // int64 | host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.HostDeletable(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostDeletable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostDeletable`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.HostDeletable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostDeletableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> HostsResp ListHosts(ctx).Limit(limit).Offset(offset).ProtectionDomainId(protectionDomainId).ClusterId(clusterId).OspClusterId(ospClusterId).Hostname(hostname).Type_(type_).Role(role).FcAvailable(fcAvailable).ReplicationGatewayAvailable(replicationGatewayAvailable).OspZoneId(ospZoneId).DatacenterIds(datacenterIds).Usage(usage).WithS3Lb(withS3Lb).Q(q).Sort(sort).Execute()





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
	protectionDomainId := int64(789) // int64 | protection domain id (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)
	ospClusterId := "ospClusterId_example" // string | osp cluster id (optional)
	hostname := "hostname_example" // string | host name (optional)
	type_ := "type__example" // string | if it existed, value should be xdcactive (optional)
	role := "role_example" // string | filter by host role (optional)
	fcAvailable := true // bool | available host with fc port (optional)
	replicationGatewayAvailable := true // bool | available host for replication gateway (optional)
	ospZoneId := int64(789) // int64 | osp zone id (optional)
	datacenterIds := int64(789) // int64 | datacenter ids (optional)
	usage := "usage_example" // string | host usage (optional)
	withS3Lb := true // bool | host used by s3 load balancer (optional)
	q := "q_example" // string | query param of search (optional)
	sort := "sort_example" // string | sort param of search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ListHosts(context.Background()).Limit(limit).Offset(offset).ProtectionDomainId(protectionDomainId).ClusterId(clusterId).OspClusterId(ospClusterId).Hostname(hostname).Type_(type_).Role(role).FcAvailable(fcAvailable).ReplicationGatewayAvailable(replicationGatewayAvailable).OspZoneId(ospZoneId).DatacenterIds(datacenterIds).Usage(usage).WithS3Lb(withS3Lb).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: HostsResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **protectionDomainId** | **int64** | protection domain id | 
 **clusterId** | **string** | cluster id | 
 **ospClusterId** | **string** | osp cluster id | 
 **hostname** | **string** | host name | 
 **type_** | **string** | if it existed, value should be xdcactive | 
 **role** | **string** | filter by host role | 
 **fcAvailable** | **bool** | available host with fc port | 
 **replicationGatewayAvailable** | **bool** | available host for replication gateway | 
 **ospZoneId** | **int64** | osp zone id | 
 **datacenterIds** | **int64** | datacenter ids | 
 **usage** | **string** | host usage | 
 **withS3Lb** | **bool** | host used by s3 load balancer | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**HostsResp**](HostsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintainHost

> HostResp MaintainHost(ctx, hostId).Force(force).Execute()





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
	hostId := int64(789) // int64 | host id
	force := true // bool | force maintain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.MaintainHost(context.Background(), hostId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.MaintainHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintainHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.MaintainHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintainHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force maintain | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHostsFormOspZone

> RemoveHostsFormOspZone(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOpHostsZoneReq() // OpHostsZoneReq | hosts cluster osp zone info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.RemoveHostsFormOspZone(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.RemoveHostsFormOspZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHostsFormOspZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpHostsZoneReq**](OpHostsZoneReq.md) | hosts cluster osp zone info | 

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


## SetHostsToOspZone

> SetHostsToOspZone(ctx).Body(body).Execute()





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
	body := *openapiclient.NewOpHostsZoneReq() // OpHostsZoneReq | hosts cluster osp zone info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.SetHostsToOspZone(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.SetHostsToOspZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetHostsToOspZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OpHostsZoneReq**](OpHostsZoneReq.md) | hosts cluster osp zone info | 

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


## UnmaintainHost

> HostResp UnmaintainHost(ctx, hostId).Execute()





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
	hostId := int64(789) // int64 | host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UnmaintainHost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UnmaintainHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmaintainHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UnmaintainHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmaintainHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> HostResp UpdateHost(ctx, hostId).Body(body).Execute()





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
	hostId := int64(789) // int64 | host id
	body := *openapiclient.NewHostUpdateReq(*openapiclient.NewHostUpdateReqHost()) // HostUpdateReq | host info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UpdateHost(context.Background(), hostId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UpdateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHost`: HostResp
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **int64** | host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostUpdateReq**](HostUpdateReq.md) | host info | 

### Return type

[**HostResp**](HostResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

