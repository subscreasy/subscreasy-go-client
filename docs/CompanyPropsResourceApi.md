# \CompanyPropsResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompanyPropsUsingPOST**](CompanyPropsResourceApi.md#CreateCompanyPropsUsingPOST) | **Post** /api/company-props | createCompanyProps
[**DeleteCompanyPropsUsingDELETE**](CompanyPropsResourceApi.md#DeleteCompanyPropsUsingDELETE) | **Delete** /api/company-props/{id} | deleteCompanyProps
[**GetAllCompanyPropsUsingGET**](CompanyPropsResourceApi.md#GetAllCompanyPropsUsingGET) | **Get** /api/company-props | getAllCompanyProps
[**GetCompanyPropsByCompanyIdUsingGET**](CompanyPropsResourceApi.md#GetCompanyPropsByCompanyIdUsingGET) | **Get** /api/company-props/company/{companyId} | getCompanyPropsByCompanyId
[**UpdateCompanyPropsUsingPUT**](CompanyPropsResourceApi.md#UpdateCompanyPropsUsingPUT) | **Put** /api/company-props | updateCompanyProps


# **CreateCompanyPropsUsingPOST**
> CompanyProps CreateCompanyPropsUsingPOST(ctx, companyProps)
createCompanyProps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **companyProps** | [**CompanyProps**](CompanyProps.md)| companyProps | 

### Return type

[**CompanyProps**](CompanyProps.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCompanyPropsUsingDELETE**
> DeleteCompanyPropsUsingDELETE(ctx, id)
deleteCompanyProps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllCompanyPropsUsingGET**
> []CompanyProps GetAllCompanyPropsUsingGET(ctx, )
getAllCompanyProps

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CompanyProps**](CompanyProps.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanyPropsByCompanyIdUsingGET**
> CompanyProps GetCompanyPropsByCompanyIdUsingGET(ctx, companyId)
getCompanyPropsByCompanyId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **companyId** | **int64**| companyId | 

### Return type

[**CompanyProps**](CompanyProps.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyPropsUsingPUT**
> CompanyProps UpdateCompanyPropsUsingPUT(ctx, companyProps)
updateCompanyProps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **companyProps** | [**CompanyProps**](CompanyProps.md)| companyProps | 

### Return type

[**CompanyProps**](CompanyProps.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

