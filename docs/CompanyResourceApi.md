# \CompanyResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompanyUsingPOST**](CompanyResourceApi.md#CreateCompanyUsingPOST) | **Post** /api/companies | createCompany
[**DeleteCompanyUsingDELETE**](CompanyResourceApi.md#DeleteCompanyUsingDELETE) | **Delete** /api/companies/{id} | deleteCompany
[**GetAllCompaniesUsingGET**](CompanyResourceApi.md#GetAllCompaniesUsingGET) | **Get** /api/companies | getAllCompanies
[**GetCompanyUsingGET**](CompanyResourceApi.md#GetCompanyUsingGET) | **Get** /api/companies/{id} | getCompany
[**UpdateCompanyUsingPUT**](CompanyResourceApi.md#UpdateCompanyUsingPUT) | **Put** /api/companies | updateCompany


# **CreateCompanyUsingPOST**
> Company CreateCompanyUsingPOST(ctx, company)
createCompany

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **company** | [**Company**](Company.md)| company | 

### Return type

[**Company**](Company.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCompanyUsingDELETE**
> DeleteCompanyUsingDELETE(ctx, id)
deleteCompany

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

# **GetAllCompaniesUsingGET**
> []Company GetAllCompaniesUsingGET(ctx, )
getAllCompanies

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Company**](Company.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompanyUsingGET**
> Company GetCompanyUsingGET(ctx, id)
getCompany

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Company**](Company.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompanyUsingPUT**
> Company UpdateCompanyUsingPUT(ctx, company)
updateCompany

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **company** | [**Company**](Company.md)| company | 

### Return type

[**Company**](Company.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

