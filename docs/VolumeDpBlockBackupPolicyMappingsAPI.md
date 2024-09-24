# \VolumeDpBlockBackupPolicyMappingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVolumeDpBlockBackupPolicyMappings**](VolumeDpBlockBackupPolicyMappingsAPI.md#ListVolumeDpBlockBackupPolicyMappings) | **Get** /volume-dp-block-backup-policy-mappings/ | 



## ListVolumeDpBlockBackupPolicyMappings

> VolumeDpBlockBackupPolicyMappingsResp ListVolumeDpBlockBackupPolicyMappings(ctx).BlockVolumeId(blockVolumeId).DpBlockBackupPolicyId(dpBlockBackupPolicyId).Execute()





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
	blockVolumeId := int64(789) // int64 | show mappings of specific block volume (optional)
	dpBlockBackupPolicyId := int64(789) // int64 | show mappings of specific block volume (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeDpBlockBackupPolicyMappingsAPI.ListVolumeDpBlockBackupPolicyMappings(context.Background()).BlockVolumeId(blockVolumeId).DpBlockBackupPolicyId(dpBlockBackupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeDpBlockBackupPolicyMappingsAPI.ListVolumeDpBlockBackupPolicyMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeDpBlockBackupPolicyMappings`: VolumeDpBlockBackupPolicyMappingsResp
	fmt.Fprintf(os.Stdout, "Response from `VolumeDpBlockBackupPolicyMappingsAPI.ListVolumeDpBlockBackupPolicyMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeDpBlockBackupPolicyMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64** | show mappings of specific block volume | 
 **dpBlockBackupPolicyId** | **int64** | show mappings of specific block volume | 

### Return type

[**VolumeDpBlockBackupPolicyMappingsResp**](VolumeDpBlockBackupPolicyMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

