# \ServiceOfferingResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceOfferingUsingPOST**](ServiceOfferingResourceApi.md#CreateServiceOfferingUsingPOST) | **Post** /api/service-offerings | createServiceOffering
[**DeleteServiceOfferingUsingDELETE**](ServiceOfferingResourceApi.md#DeleteServiceOfferingUsingDELETE) | **Delete** /api/service-offerings/{id} | deleteServiceOffering
[**GetAllServiceOfferingsUsingGET**](ServiceOfferingResourceApi.md#GetAllServiceOfferingsUsingGET) | **Get** /api/service-offerings | getAllServiceOfferings
[**GetServiceOfferingUsingGET**](ServiceOfferingResourceApi.md#GetServiceOfferingUsingGET) | **Get** /api/service-offerings/{id} | getServiceOffering
[**UpdateServiceOfferingUsingPUT**](ServiceOfferingResourceApi.md#UpdateServiceOfferingUsingPUT) | **Put** /api/service-offerings | updateServiceOffering


# **CreateServiceOfferingUsingPOST**
> ServiceOffering CreateServiceOfferingUsingPOST(ctx, serviceOffering)
createServiceOffering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceOffering** | [**ServiceOffering**](ServiceOffering.md)| serviceOffering | 

### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceOfferingUsingDELETE**
> DeleteServiceOfferingUsingDELETE(ctx, id)
deleteServiceOffering

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

# **GetAllServiceOfferingsUsingGET**
> []ServiceOffering GetAllServiceOfferingsUsingGET(ctx, )
getAllServiceOfferings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ServiceOffering**](ServiceOffering.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceOfferingUsingGET**
> ServiceOffering GetServiceOfferingUsingGET(ctx, id)
getServiceOffering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceOfferingUsingPUT**
> ServiceOffering UpdateServiceOfferingUsingPUT(ctx, serviceOffering)
updateServiceOffering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceOffering** | [**ServiceOffering**](ServiceOffering.md)| serviceOffering | 

### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

