# \OsdGroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOsdsToOsdGroup**](OsdGroupsAPI.md#AddOsdsToOsdGroup) | **Post** /osd-groups/{group_id}:add-osds | 
[**DisableDeviceTypeCheck**](OsdGroupsAPI.md#DisableDeviceTypeCheck) | **Post** /osd-groups/{group_id}:disable-device-type-check | 
[**EnableDeviceTypeCheck**](OsdGroupsAPI.md#EnableDeviceTypeCheck) | **Post** /osd-groups/{group_id}:enable-device-type-check | 
[**GetOsdGroup**](OsdGroupsAPI.md#GetOsdGroup) | **Get** /osd-groups/{group_id} | 
[**GetOsdGroupSamples**](OsdGroupsAPI.md#GetOsdGroupSamples) | **Get** /osd-groups/{group_id}/samples | 
[**ListOsdGroups**](OsdGroupsAPI.md#ListOsdGroups) | **Get** /osd-groups/ | 
[**RemoveOsdsFromOsdGroup**](OsdGroupsAPI.md#RemoveOsdsFromOsdGroup) | **Post** /osd-groups/{group_id}:remove-osds | 
[**ReweightOsdGroup**](OsdGroupsAPI.md#ReweightOsdGroup) | **Post** /osd-groups/{group_id}:reweight | 
[**SetOsdFullRatio**](OsdGroupsAPI.md#SetOsdFullRatio) | **Post** /osd-groups/{group_id}:set-osd-full-ratio | 
[**SetOsdGroupQos**](OsdGroupsAPI.md#SetOsdGroupQos) | **Post** /osd-groups/{group_id}:set-qos | 



## AddOsdsToOsdGroup

> OsdGroupResp AddOsdsToOsdGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | osd group id
	body := *openapiclient.NewOsdGroupOsdsUpdateReq(*openapiclient.NewOsdGroupOsdsUpdateReqOsdGroup([]int64{int64(123)})) // OsdGroupOsdsUpdateReq | osd ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.AddOsdsToOsdGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.AddOsdsToOsdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOsdsToOsdGroup`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.AddOsdsToOsdGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOsdsToOsdGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdGroupOsdsUpdateReq**](OsdGroupOsdsUpdateReq.md) | osd ids | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableDeviceTypeCheck

> OsdGroupResp DisableDeviceTypeCheck(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | osd group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.DisableDeviceTypeCheck(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.DisableDeviceTypeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableDeviceTypeCheck`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.DisableDeviceTypeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableDeviceTypeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDeviceTypeCheck

> OsdGroupResp EnableDeviceTypeCheck(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | osd group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.EnableDeviceTypeCheck(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.EnableDeviceTypeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDeviceTypeCheck`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.EnableDeviceTypeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableDeviceTypeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsdGroup

> OsdGroupResp GetOsdGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | osd group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.GetOsdGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.GetOsdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsdGroup`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.GetOsdGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsdGroupSamples

> OsdGroupSamplesResp GetOsdGroupSamples(ctx, groupId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()





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
	groupId := int64(789) // int64 | osd group id
	durationBegin := "durationBegin_example" // string | duration begin timestamp (optional)
	durationEnd := "durationEnd_example" // string | duration end timestamp (optional)
	period := "period_example" // string | samples period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.GetOsdGroupSamples(context.Background(), groupId).DurationBegin(durationBegin).DurationEnd(durationEnd).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.GetOsdGroupSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsdGroupSamples`: OsdGroupSamplesResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.GetOsdGroupSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsdGroupSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **durationBegin** | **string** | duration begin timestamp | 
 **durationEnd** | **string** | duration end timestamp | 
 **period** | **string** | samples period | 

### Return type

[**OsdGroupSamplesResp**](OsdGroupSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOsdGroups

> OsdGroupsResp ListOsdGroups(ctx).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.ListOsdGroups(context.Background()).Limit(limit).Offset(offset).Q(q).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.ListOsdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOsdGroups`: OsdGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.ListOsdGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOsdGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **q** | **string** | query param of search | 
 **sort** | **string** | sort param of search | 

### Return type

[**OsdGroupsResp**](OsdGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOsdsFromOsdGroup

> OsdGroupResp RemoveOsdsFromOsdGroup(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | osd group id
	body := *openapiclient.NewOsdGroupOsdsUpdateReq(*openapiclient.NewOsdGroupOsdsUpdateReqOsdGroup([]int64{int64(123)})) // OsdGroupOsdsUpdateReq | remove osd ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.RemoveOsdsFromOsdGroup(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.RemoveOsdsFromOsdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveOsdsFromOsdGroup`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.RemoveOsdsFromOsdGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOsdsFromOsdGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdGroupOsdsUpdateReq**](OsdGroupOsdsUpdateReq.md) | remove osd ids | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReweightOsdGroup

> OsdGroupResp ReweightOsdGroup(ctx, groupId).Execute()





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
	groupId := int64(789) // int64 | osd group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.ReweightOsdGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.ReweightOsdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReweightOsdGroup`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.ReweightOsdGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReweightOsdGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOsdFullRatio

> OsdGroupResp SetOsdFullRatio(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | osd group id
	body := *openapiclient.NewSetOsdFullRatioReq(*openapiclient.NewSetOsdFullRatioReqOsdGroup(float64(123))) // SetOsdFullRatioReq | osds full ratio

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.SetOsdFullRatio(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.SetOsdFullRatio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOsdFullRatio`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.SetOsdFullRatio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOsdFullRatioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetOsdFullRatioReq**](SetOsdFullRatioReq.md) | osds full ratio | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOsdGroupQos

> OsdGroupResp SetOsdGroupQos(ctx, groupId).Body(body).Execute()





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
	groupId := int64(789) // int64 | osd group id
	body := *openapiclient.NewOsdGroupSetQosReq(*openapiclient.NewOsdGroupSetQosReqOsdGroup(*openapiclient.NewOsdQos())) // OsdGroupSetQosReq | qos info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsdGroupsAPI.SetOsdGroupQos(context.Background(), groupId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsdGroupsAPI.SetOsdGroupQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOsdGroupQos`: OsdGroupResp
	fmt.Fprintf(os.Stdout, "Response from `OsdGroupsAPI.SetOsdGroupQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | osd group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOsdGroupQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OsdGroupSetQosReq**](OsdGroupSetQosReq.md) | qos info | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

