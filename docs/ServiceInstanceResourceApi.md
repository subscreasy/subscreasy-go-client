# \ServiceInstanceResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceInstanceUsingPOST**](ServiceInstanceResourceApi.md#CreateServiceInstanceUsingPOST) | **Post** /api/service-instances | createServiceInstance
[**DeleteServiceInstanceUsingDELETE**](ServiceInstanceResourceApi.md#DeleteServiceInstanceUsingDELETE) | **Delete** /api/service-instances/{id} | deleteServiceInstance
[**GetAllServiceInstancesUsingGET**](ServiceInstanceResourceApi.md#GetAllServiceInstancesUsingGET) | **Get** /api/service-instances | getAllServiceInstances
[**GetServiceInstanceUsingGET**](ServiceInstanceResourceApi.md#GetServiceInstanceUsingGET) | **Get** /api/service-instances/{id} | getServiceInstance
[**UpdateServiceInstanceUsingPUT**](ServiceInstanceResourceApi.md#UpdateServiceInstanceUsingPUT) | **Put** /api/service-instances | updateServiceInstance


# **CreateServiceInstanceUsingPOST**
> ServiceInstance CreateServiceInstanceUsingPOST(ctx, serviceInstance)
createServiceInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceInstance** | [**ServiceInstance**](ServiceInstance.md)| serviceInstance | 

### Return type

[**ServiceInstance**](ServiceInstance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInstanceUsingDELETE**
> DeleteServiceInstanceUsingDELETE(ctx, id)
deleteServiceInstance

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

# **GetAllServiceInstancesUsingGET**
> []ServiceInstance GetAllServiceInstancesUsingGET(ctx, )
getAllServiceInstances

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ServiceInstance**](ServiceInstance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstanceUsingGET**
> ServiceInstance GetServiceInstanceUsingGET(ctx, id)
getServiceInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**ServiceInstance**](ServiceInstance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInstanceUsingPUT**
> ServiceInstance UpdateServiceInstanceUsingPUT(ctx, serviceInstance)
updateServiceInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceInstance** | [**ServiceInstance**](ServiceInstance.md)| serviceInstance | 

### Return type

[**ServiceInstance**](ServiceInstance.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

