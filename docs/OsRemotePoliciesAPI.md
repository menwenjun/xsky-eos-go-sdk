# \OsRemotePoliciesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSRemotePolicy**](OsRemotePoliciesAPI.md#GetOSRemotePolicy) | **Get** /os-remote-policies/{policy_uuid} | 
[**ListOSRemotePolicies**](OsRemotePoliciesAPI.md#ListOSRemotePolicies) | **Get** /os-remote-policies/ | 



## GetOSRemotePolicy

> OSRemotePolicyResp GetOSRemotePolicy(ctx, policyUuid).Execute()





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
	policyUuid := "policyUuid_example" // string | policy uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsRemotePoliciesAPI.GetOSRemotePolicy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsRemotePoliciesAPI.GetOSRemotePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOSRemotePolicy`: OSRemotePolicyResp
	fmt.Fprintf(os.Stdout, "Response from `OsRemotePoliciesAPI.GetOSRemotePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | policy uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOSRemotePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSRemotePolicyResp**](OSRemotePolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOSRemotePolicies

> OSRemotePoliciesResp ListOSRemotePolicies(ctx).Limit(limit).Offset(offset).Marker(marker).ZoneUuid(zoneUuid).ClusterId(clusterId).Execute()





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
	marker := "marker_example" // string | paging param (optional)
	zoneUuid := "zoneUuid_example" // string | zone uuid (optional)
	clusterId := "clusterId_example" // string | cluster id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OsRemotePoliciesAPI.ListOSRemotePolicies(context.Background()).Limit(limit).Offset(offset).Marker(marker).ZoneUuid(zoneUuid).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OsRemotePoliciesAPI.ListOSRemotePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOSRemotePolicies`: OSRemotePoliciesResp
	fmt.Fprintf(os.Stdout, "Response from `OsRemotePoliciesAPI.ListOSRemotePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOSRemotePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | paging param | 
 **offset** | **int64** | paging param | 
 **marker** | **string** | paging param | 
 **zoneUuid** | **string** | zone uuid | 
 **clusterId** | **string** | cluster id | 

### Return type

[**OSRemotePoliciesResp**](OSRemotePoliciesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

