/*
XMS API

XMS is the controller of distributed storage system

API version: XSCALEROS_6.4.000.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// NetworkAddressesAPIService NetworkAddressesAPI service
type NetworkAddressesAPIService service

type ApiGetNetworkAddressRequest struct {
	ctx context.Context
	ApiService *NetworkAddressesAPIService
	networkAddressId int64
}

func (r ApiGetNetworkAddressRequest) Execute() (*NetworkAddressResp, *http.Response, error) {
	return r.ApiService.GetNetworkAddressExecute(r)
}

/*
GetNetworkAddress Method for GetNetworkAddress

Get a network address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkAddressId network address id
 @return ApiGetNetworkAddressRequest
*/
func (a *NetworkAddressesAPIService) GetNetworkAddress(ctx context.Context, networkAddressId int64) ApiGetNetworkAddressRequest {
	return ApiGetNetworkAddressRequest{
		ApiService: a,
		ctx: ctx,
		networkAddressId: networkAddressId,
	}
}

// Execute executes the request
//  @return NetworkAddressResp
func (a *NetworkAddressesAPIService) GetNetworkAddressExecute(r ApiGetNetworkAddressRequest) (*NetworkAddressResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NetworkAddressResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkAddressesAPIService.GetNetworkAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-addresses/{network_address_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"network_address_id"+"}", url.PathEscape(parameterValueToString(r.networkAddressId, "networkAddressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListNetworkAddressesRequest struct {
	ctx context.Context
	ApiService *NetworkAddressesAPIService
	limit *int64
	offset *int64
	networkInterfaceId *int64
	hostId *int64
	role *string
	vipGroupId *int64
}

// paging param
func (r ApiListNetworkAddressesRequest) Limit(limit int64) ApiListNetworkAddressesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListNetworkAddressesRequest) Offset(offset int64) ApiListNetworkAddressesRequest {
	r.offset = &offset
	return r
}

// network interface id
func (r ApiListNetworkAddressesRequest) NetworkInterfaceId(networkInterfaceId int64) ApiListNetworkAddressesRequest {
	r.networkInterfaceId = &networkInterfaceId
	return r
}

// host id
func (r ApiListNetworkAddressesRequest) HostId(hostId int64) ApiListNetworkAddressesRequest {
	r.hostId = &hostId
	return r
}

// network address role
func (r ApiListNetworkAddressesRequest) Role(role string) ApiListNetworkAddressesRequest {
	r.role = &role
	return r
}

// vip group id
func (r ApiListNetworkAddressesRequest) VipGroupId(vipGroupId int64) ApiListNetworkAddressesRequest {
	r.vipGroupId = &vipGroupId
	return r
}

func (r ApiListNetworkAddressesRequest) Execute() (*NetworkAddressesResp, *http.Response, error) {
	return r.ApiService.ListNetworkAddressesExecute(r)
}

/*
ListNetworkAddresses Method for ListNetworkAddresses

List network addresses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkAddressesRequest
*/
func (a *NetworkAddressesAPIService) ListNetworkAddresses(ctx context.Context) ApiListNetworkAddressesRequest {
	return ApiListNetworkAddressesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkAddressesResp
func (a *NetworkAddressesAPIService) ListNetworkAddressesExecute(r ApiListNetworkAddressesRequest) (*NetworkAddressesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NetworkAddressesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkAddressesAPIService.ListNetworkAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-addresses/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.networkInterfaceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network_interface_id", r.networkInterfaceId, "form", "")
	}
	if r.hostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host_id", r.hostId, "form", "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "form", "")
	}
	if r.vipGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vip_group_id", r.vipGroupId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Xms-Auth-Token"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenInQuery"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
