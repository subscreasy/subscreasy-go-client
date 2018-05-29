# \InvoiceResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoiceUsingPOST**](InvoiceResourceApi.md#CreateInvoiceUsingPOST) | **Post** /api/invoices | createInvoice
[**DeleteInvoiceUsingDELETE**](InvoiceResourceApi.md#DeleteInvoiceUsingDELETE) | **Delete** /api/invoices/{id} | deleteInvoice
[**GetAllInvoicesUsingGET**](InvoiceResourceApi.md#GetAllInvoicesUsingGET) | **Get** /api/invoices | getAllInvoices
[**GetInvoiceBySubscriberUsingGET**](InvoiceResourceApi.md#GetInvoiceBySubscriberUsingGET) | **Get** /api/invoices/subscriber/{subscriberSecureId} | getInvoiceBySubscriber
[**GetInvoiceUsingGET**](InvoiceResourceApi.md#GetInvoiceUsingGET) | **Get** /api/invoices/{id} | getInvoice
[**UpdateInvoiceUsingPUT**](InvoiceResourceApi.md#UpdateInvoiceUsingPUT) | **Put** /api/invoices | updateInvoice


# **CreateInvoiceUsingPOST**
> Invoice CreateInvoiceUsingPOST(ctx, invoice)
createInvoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **invoice** | [**Invoice**](Invoice.md)| invoice | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInvoiceUsingDELETE**
> DeleteInvoiceUsingDELETE(ctx, id)
deleteInvoice

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

# **GetAllInvoicesUsingGET**
> []Invoice GetAllInvoicesUsingGET(ctx, )
getAllInvoices

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceBySubscriberUsingGET**
> []Invoice GetInvoiceBySubscriberUsingGET(ctx, subscriberSecureId)
getInvoiceBySubscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subscriberSecureId** | **string**| subscriberSecureId | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceUsingGET**
> Invoice GetInvoiceUsingGET(ctx, id)
getInvoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInvoiceUsingPUT**
> Invoice UpdateInvoiceUsingPUT(ctx, invoice)
updateInvoice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **invoice** | [**Invoice**](Invoice.md)| invoice | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

