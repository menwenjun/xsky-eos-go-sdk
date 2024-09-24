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
	"os"
)


// LicensesAPIService LicensesAPI service
type LicensesAPIService service

type ApiDisableLicenseRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	licenseId int64
	force *bool
}

// force disable not expired license
func (r ApiDisableLicenseRequest) Force(force bool) ApiDisableLicenseRequest {
	r.force = &force
	return r
}

func (r ApiDisableLicenseRequest) Execute() (*LicenseResp, *http.Response, error) {
	return r.ApiService.DisableLicenseExecute(r)
}

/*
DisableLicense Method for DisableLicense

disable license

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param licenseId the license id
 @return ApiDisableLicenseRequest
*/
func (a *LicensesAPIService) DisableLicense(ctx context.Context, licenseId int64) ApiDisableLicenseRequest {
	return ApiDisableLicenseRequest{
		ApiService: a,
		ctx: ctx,
		licenseId: licenseId,
	}
}

// Execute executes the request
//  @return LicenseResp
func (a *LicensesAPIService) DisableLicenseExecute(r ApiDisableLicenseRequest) (*LicenseResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.DisableLicense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/{license_id}:disable"
	localVarPath = strings.Replace(localVarPath, "{"+"license_id"+"}", url.PathEscape(parameterValueToString(r.licenseId, "licenseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
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

type ApiDownloadLicenseKeyRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
}

func (r ApiDownloadLicenseKeyRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.DownloadLicenseKeyExecute(r)
}

/*
DownloadLicenseKey Method for DownloadLicenseKey

get license key file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDownloadLicenseKeyRequest
*/
func (a *LicensesAPIService) DownloadLicenseKey(ctx context.Context) ApiDownloadLicenseKeyRequest {
	return ApiDownloadLicenseKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *LicensesAPIService) DownloadLicenseKeyExecute(r ApiDownloadLicenseKeyRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.DownloadLicenseKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/key"

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiGetLicenseRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	licenseId int64
}

func (r ApiGetLicenseRequest) Execute() (*LicenseResp, *http.Response, error) {
	return r.ApiService.GetLicenseExecute(r)
}

/*
GetLicense Method for GetLicense

get license

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param licenseId the license id
 @return ApiGetLicenseRequest
*/
func (a *LicensesAPIService) GetLicense(ctx context.Context, licenseId int64) ApiGetLicenseRequest {
	return ApiGetLicenseRequest{
		ApiService: a,
		ctx: ctx,
		licenseId: licenseId,
	}
}

// Execute executes the request
//  @return LicenseResp
func (a *LicensesAPIService) GetLicenseExecute(r ApiGetLicenseRequest) (*LicenseResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.GetLicense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/{license_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"license_id"+"}", url.PathEscape(parameterValueToString(r.licenseId, "licenseId")), -1)

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

type ApiGetLicenseSummaryRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
}

func (r ApiGetLicenseSummaryRequest) Execute() (*LicenseSummaryResp, *http.Response, error) {
	return r.ApiService.GetLicenseSummaryExecute(r)
}

/*
GetLicenseSummary Method for GetLicenseSummary

Get licenses sumary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLicenseSummaryRequest
*/
func (a *LicensesAPIService) GetLicenseSummary(ctx context.Context) ApiGetLicenseSummaryRequest {
	return ApiGetLicenseSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LicenseSummaryResp
func (a *LicensesAPIService) GetLicenseSummaryExecute(r ApiGetLicenseSummaryRequest) (*LicenseSummaryResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseSummaryResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.GetLicenseSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/summary"

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

type ApiListLicensesRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	limit *int64
	offset *int64
	q *string
	sort *string
	active *bool
}

// paging param
func (r ApiListLicensesRequest) Limit(limit int64) ApiListLicensesRequest {
	r.limit = &limit
	return r
}

// paging param
func (r ApiListLicensesRequest) Offset(offset int64) ApiListLicensesRequest {
	r.offset = &offset
	return r
}

// query param of search
func (r ApiListLicensesRequest) Q(q string) ApiListLicensesRequest {
	r.q = &q
	return r
}

// sort param of search
func (r ApiListLicensesRequest) Sort(sort string) ApiListLicensesRequest {
	r.sort = &sort
	return r
}

// filter license by active status
func (r ApiListLicensesRequest) Active(active bool) ApiListLicensesRequest {
	r.active = &active
	return r
}

func (r ApiListLicensesRequest) Execute() (*LicensesResp, *http.Response, error) {
	return r.ApiService.ListLicensesExecute(r)
}

/*
ListLicenses Method for ListLicenses

List licenses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListLicensesRequest
*/
func (a *LicensesAPIService) ListLicenses(ctx context.Context) ApiListLicensesRequest {
	return ApiListLicensesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LicensesResp
func (a *LicensesAPIService) ListLicensesExecute(r ApiListLicensesRequest) (*LicensesResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicensesResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.ListLicenses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
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

type ApiRegisterLicenseRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	license *os.File
	force *bool
	dryRun *bool
}

// license info
func (r ApiRegisterLicenseRequest) License(license *os.File) ApiRegisterLicenseRequest {
	r.license = license
	return r
}

// force activate
func (r ApiRegisterLicenseRequest) Force(force bool) ApiRegisterLicenseRequest {
	r.force = &force
	return r
}

// dry run upload license file
func (r ApiRegisterLicenseRequest) DryRun(dryRun bool) ApiRegisterLicenseRequest {
	r.dryRun = &dryRun
	return r
}

func (r ApiRegisterLicenseRequest) Execute() (*LicenseResp, *http.Response, error) {
	return r.ApiService.RegisterLicenseExecute(r)
}

/*
RegisterLicense Method for RegisterLicense

Activate product license

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterLicenseRequest
*/
func (a *LicensesAPIService) RegisterLicense(ctx context.Context) ApiRegisterLicenseRequest {
	return ApiRegisterLicenseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LicenseResp
func (a *LicensesAPIService) RegisterLicenseExecute(r ApiRegisterLicenseRequest) (*LicenseResp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseResp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.RegisterLicense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.license == nil {
		return localVarReturnValue, nil, reportError("license is required and must be specified")
	}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
	}
	if r.dryRun != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dry_run", r.dryRun, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var licenseLocalVarFormFileName string
	var licenseLocalVarFileName     string
	var licenseLocalVarFileBytes    []byte

	licenseLocalVarFormFileName = "license"
	licenseLocalVarFile := r.license

	if licenseLocalVarFile != nil {
		fbs, _ := io.ReadAll(licenseLocalVarFile)

		licenseLocalVarFileBytes = fbs
		licenseLocalVarFileName = licenseLocalVarFile.Name()
		licenseLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: licenseLocalVarFileBytes, fileName: licenseLocalVarFileName, formFileName: licenseLocalVarFormFileName})
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
