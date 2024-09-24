# \S3LoadBalancerGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddS3LoadBalancersToGroup**](S3LoadBalancerGroupsAPI.md#AddS3LoadBalancersToGroup) | **Put** /s3-load-balancer-groups/{group_id}/s3-load-balancers | 
[**CreateS3LoadBalancerGroup**](S3LoadBalancerGroupsAPI.md#CreateS3LoadBalancerGroup) | **Post** /s3-load-balancer-groups/ | 
[**DeleteS3LoadBalancerGroup**](S3LoadBalancerGroupsAPI.md#DeleteS3LoadBalancerGroup) | **Delete** /s3-load-balancer-groups/{group_id} | 
[**GetS3LoadBalancerGroup**](S3LoadBalancerGroupsAPI.md#GetS3LoadBalancerGroup) | **Get** /s3-load-balancer-groups/{group_id} | 
[**ListS3LoadBalancerGroups**](S3LoadBalancerGroupsAPI.md#ListS3LoadBalancerGroups) | **Get** /s3-load-balancer-groups/ | 
[**RedeployS3LoadBalancerGroup**](S3LoadBalancerGroupsAPI.md#RedeployS3LoadBalancerGroup) | **Post** /s3-load-balancer-groups/{group_id}:redeploy | 
[**RegisterS3LoadBalancerService**](S3LoadBalancerGroupsAPI.md#RegisterS3LoadBalancerService) | **Post** /s3-load-balancer-groups/register-service | 
[**RemoveS3LoadBalancerService**](S3LoadBalancerGroupsAPI.md#RemoveS3LoadBalancerService) | **Delete** /s3-load-balancer-groups/remove-service | 
[**RemoveS3LoadBalancersFromGroup**](S3LoadBalancerGroupsAPI.md#RemoveS3LoadBalancersFromGroup) | **Delete** /s3-load-balancer-groups/{group_id}/s3-load-balancers | 
[**UpdateS3LoadBalancerGroup**](S3LoadBalancerGroupsAPI.md#UpdateS3LoadBalancerGroup) | **Patch** /s3-load-balancer-groups/{group_id} | 



## AddS3LoadBalancersToGroup

> S3LoadBalancerGroupResp AddS3LoadBalancersToGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | group id
	body := *openapiclient.NewS3LoadBalancersAddReq([]openapiclient.S3LoadBalancersAddReqLoadBalancersElt{*openapiclient.NewS3LoadBalancersAddReqLoadBalancersElt(int64(123), "Vip_example")}) // S3LoadBalancersAddReq | load balancer ids to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.AddS3LoadBalancersToGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.AddS3LoadBalancersToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddS3LoadBalancersToGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.AddS3LoadBalancersToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddS3LoadBalancersToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3LoadBalancersAddReq**](S3LoadBalancersAddReq.md) | load balancer ids to add | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateS3LoadBalancerGroup

> S3LoadBalancerGroupResp CreateS3LoadBalancerGroup(ctx).Body(body).ClusterId(clusterId).Execute()





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
	body := *openapiclient.NewS3LoadBalancerGroupCreateReq(*openapiclient.NewS3LoadBalancerGroupCreateReqGroup("Name_example", []openapiclient.S3LoadBalancerGroupCreateReqGroupLoadBalancersElt{*openapiclient.NewS3LoadBalancerGroupCreateReqGroupLoadBalancersElt(int64(123), "Vip_example")})) // S3LoadBalancerGroupCreateReq | s3 load balancer group info
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.CreateS3LoadBalancerGroup(context.Background()).Body(body).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.CreateS3LoadBalancerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateS3LoadBalancerGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.CreateS3LoadBalancerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateS3LoadBalancerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**S3LoadBalancerGroupCreateReq**](S3LoadBalancerGroupCreateReq.md) | s3 load balancer group info | 
 **clusterId** | **string** | cluster id | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteS3LoadBalancerGroup

> S3LoadBalancerGroupResp DeleteS3LoadBalancerGroup(ctx, groupId).Force(force).Execute()





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
	groupId := int64(789) // int64 | s3 load balancer group id
	force := true // bool | forcedly delete or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.DeleteS3LoadBalancerGroup(context.Background(), groupId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.DeleteS3LoadBalancerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteS3LoadBalancerGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.DeleteS3LoadBalancerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | s3 load balancer group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3LoadBalancerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | forcedly delete or not | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3LoadBalancerGroup

> S3LoadBalancerGroupResp GetS3LoadBalancerGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | s3 load balancer group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.GetS3LoadBalancerGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.GetS3LoadBalancerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3LoadBalancerGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.GetS3LoadBalancerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | s3 load balancer group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3LoadBalancerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListS3LoadBalancerGroups

> S3LoadBalancerGroupsResp ListS3LoadBalancerGroups(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).OspZoneId(ospZoneId).Execute()





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
	ospZoneId := int64(789) // int64 | osp zone id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.ListS3LoadBalancerGroups(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).ClusterId(clusterId).OspZoneId(ospZoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.ListS3LoadBalancerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListS3LoadBalancerGroups`: S3LoadBalancerGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.ListS3LoadBalancerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListS3LoadBalancerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 
 **clusterId** | **string** | cluster id | 
 **ospZoneId** | **int64** | osp zone id | 

### Return type

[**S3LoadBalancerGroupsResp**](S3LoadBalancerGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployS3LoadBalancerGroup

> S3LoadBalancerGroupResp RedeployS3LoadBalancerGroup(ctx, groupId).ReloadKeepalived(reloadKeepalived).Execute()





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
	groupId := int64(789) // int64 | s3 load balancer group id
	reloadKeepalived := true // bool | reload keepalived (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.RedeployS3LoadBalancerGroup(context.Background(), groupId).ReloadKeepalived(reloadKeepalived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.RedeployS3LoadBalancerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeployS3LoadBalancerGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.RedeployS3LoadBalancerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | s3 load balancer group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeployS3LoadBalancerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reloadKeepalived** | **bool** | reload keepalived | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterS3LoadBalancerService

> S3LoadBalancerServiceResp RegisterS3LoadBalancerService(ctx).Body(body).Execute()





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
	body := *openapiclient.NewS3LoadBalancerRegisterServiceReq([]openapiclient.S3LoadBalancerRegisterServiceReqRegisterRulesElt{*openapiclient.NewS3LoadBalancerRegisterServiceReqRegisterRulesElt([]string{"RecognitionRules_example"}, []openapiclient.S3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt{*openapiclient.NewS3LoadBalancerRegisterServiceReqRegisterRulesEltTargetEndpointsElt([]string{"Endpoints_example"}, int64(123))})}) // S3LoadBalancerRegisterServiceReq | register s3 load balancer service info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.RegisterS3LoadBalancerService(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.RegisterS3LoadBalancerService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterS3LoadBalancerService`: S3LoadBalancerServiceResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.RegisterS3LoadBalancerService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterS3LoadBalancerServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**S3LoadBalancerRegisterServiceReq**](S3LoadBalancerRegisterServiceReq.md) | register s3 load balancer service info | 

### Return type

[**S3LoadBalancerServiceResp**](S3LoadBalancerServiceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveS3LoadBalancerService

> S3LoadBalancerServiceResp RemoveS3LoadBalancerService(ctx).Body(body).Execute()





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
	body := *openapiclient.NewS3LoadBalancerRemoveServiceReq([]openapiclient.S3LoadBalancerRemoveServiceReqRemoveRulesElt{*openapiclient.NewS3LoadBalancerRemoveServiceReqRemoveRulesElt(int64(123), []string{"Rules_example"})}) // S3LoadBalancerRemoveServiceReq | remove s3 load balancer service info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.RemoveS3LoadBalancerService(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.RemoveS3LoadBalancerService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveS3LoadBalancerService`: S3LoadBalancerServiceResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.RemoveS3LoadBalancerService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveS3LoadBalancerServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**S3LoadBalancerRemoveServiceReq**](S3LoadBalancerRemoveServiceReq.md) | remove s3 load balancer service info | 

### Return type

[**S3LoadBalancerServiceResp**](S3LoadBalancerServiceResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveS3LoadBalancersFromGroup

> S3LoadBalancerGroupResp RemoveS3LoadBalancersFromGroup(ctx, groupId).Body(body).Force(force).Execute()





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
	groupId := int64(789) // int64 | group id
	body := *openapiclient.NewS3LoadBalancersRemoveReq() // S3LoadBalancersRemoveReq | load balancer ids to remove
	force := true // bool | forcedly remove or not (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.RemoveS3LoadBalancersFromGroup(context.Background(), groupId).Body(body).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.RemoveS3LoadBalancersFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveS3LoadBalancersFromGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.RemoveS3LoadBalancersFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveS3LoadBalancersFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3LoadBalancersRemoveReq**](S3LoadBalancersRemoveReq.md) | load balancer ids to remove | 
 **force** | **bool** | forcedly remove or not | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateS3LoadBalancerGroup

> S3LoadBalancerGroupResp UpdateS3LoadBalancerGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | s3 load balancer group id
	body := *openapiclient.NewS3LoadBalancerGroupUpdateReq(*openapiclient.NewS3LoadBalancerGroupUpdateReqGroup()) // S3LoadBalancerGroupUpdateReq | s3 load balancer group info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3LoadBalancerGroupsAPI.UpdateS3LoadBalancerGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3LoadBalancerGroupsAPI.UpdateS3LoadBalancerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateS3LoadBalancerGroup`: S3LoadBalancerGroupResp
	fmt.Fprintf(os.Stdout, "Response from `S3LoadBalancerGroupsAPI.UpdateS3LoadBalancerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | s3 load balancer group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateS3LoadBalancerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3LoadBalancerGroupUpdateReq**](S3LoadBalancerGroupUpdateReq.md) | s3 load balancer group info | 

### Return type

[**S3LoadBalancerGroupResp**](S3LoadBalancerGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

