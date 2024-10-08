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


// OsSamplesAPIService OsSamplesAPI service
type OsSamplesAPIService service

type ApiGetOSSamplesRequest struct {
	ctx context.Context
	ApiService *OsSamplesAPIService
	time *string
	clusterId *string
}

// query time(url encode RFC3339)
func (r ApiGetOSSamplesRequest) Time(time string) ApiGetOSSamplesRequest {
	r.time = &time
	return r
}

// cluster id
func (r ApiGetOSSamplesRequest) ClusterId(clusterId string) ApiGetOSSamplesRequest {
	r.clusterId = &clusterId
	return r
}

func (r ApiGetOSSamplesRequest) Execute() (*OSSampleResp, *http.Response, error) {
	return r.ApiService.GetOSSamplesExecute(r)
}

/*
GetOSSamples Method for GetOSSamples

Get os samples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOSSamplesRequest
*/
func (a *OsSamplesAPIService) GetOSSamples(ctx context.Context) ApiGetOSSamplesRequest {
	return ApiGetOSSamplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OSSampleResp
func (a *OsSamplesAPIService) GetOSSamplesExecute(r ApiGetOSSamplesRequest) (*OSSampleResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSSampleResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsSamplesAPIService.GetOSSamples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/os-samples/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
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

type ApiGetOSSamplesByBucketNameRequest struct {
	ctx context.Context
	ApiService *OsSamplesAPIService
	userName string
	bucketName string
	time *string
	clusterId *string
}

// query time(url encode RFC3339)
func (r ApiGetOSSamplesByBucketNameRequest) Time(time string) ApiGetOSSamplesByBucketNameRequest {
	r.time = &time
	return r
}

// cluster id
func (r ApiGetOSSamplesByBucketNameRequest) ClusterId(clusterId string) ApiGetOSSamplesByBucketNameRequest {
	r.clusterId = &clusterId
	return r
}

func (r ApiGetOSSamplesByBucketNameRequest) Execute() (*OSSampleResp, *http.Response, error) {
	return r.ApiService.GetOSSamplesByBucketNameExecute(r)
}

/*
GetOSSamplesByBucketName Method for GetOSSamplesByBucketName

get os samples by os bucket name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userName os user name
 @param bucketName os bucket name
 @return ApiGetOSSamplesByBucketNameRequest
*/
func (a *OsSamplesAPIService) GetOSSamplesByBucketName(ctx context.Context, userName string, bucketName string) ApiGetOSSamplesByBucketNameRequest {
	return ApiGetOSSamplesByBucketNameRequest{
		ApiService: a,
		ctx: ctx,
		userName: userName,
		bucketName: bucketName,
	}
}

// Execute executes the request
//  @return OSSampleResp
func (a *OsSamplesAPIService) GetOSSamplesByBucketNameExecute(r ApiGetOSSamplesByBucketNameRequest) (*OSSampleResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSSampleResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsSamplesAPIService.GetOSSamplesByBucketName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/os-samples/{user_name}/{bucket_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucket_name"+"}", url.PathEscape(parameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
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

type ApiGetOSSamplesByUserNameRequest struct {
	ctx context.Context
	ApiService *OsSamplesAPIService
	userName string
	time *string
	clusterId *string
}

// query time(url encode RFC3339)
func (r ApiGetOSSamplesByUserNameRequest) Time(time string) ApiGetOSSamplesByUserNameRequest {
	r.time = &time
	return r
}

// cluster id
func (r ApiGetOSSamplesByUserNameRequest) ClusterId(clusterId string) ApiGetOSSamplesByUserNameRequest {
	r.clusterId = &clusterId
	return r
}

func (r ApiGetOSSamplesByUserNameRequest) Execute() (*OSSampleResp, *http.Response, error) {
	return r.ApiService.GetOSSamplesByUserNameExecute(r)
}

/*
GetOSSamplesByUserName Method for GetOSSamplesByUserName

Get os samples by user name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userName os user name
 @return ApiGetOSSamplesByUserNameRequest
*/
func (a *OsSamplesAPIService) GetOSSamplesByUserName(ctx context.Context, userName string) ApiGetOSSamplesByUserNameRequest {
	return ApiGetOSSamplesByUserNameRequest{
		ApiService: a,
		ctx: ctx,
		userName: userName,
	}
}

// Execute executes the request
//  @return OSSampleResp
func (a *OsSamplesAPIService) GetOSSamplesByUserNameExecute(r ApiGetOSSamplesByUserNameRequest) (*OSSampleResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OSSampleResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsSamplesAPIService.GetOSSamplesByUserName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/os-samples/{user_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_name"+"}", url.PathEscape(parameterValueToString(r.userName, "userName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cluster_id", r.clusterId, "form", "")
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
