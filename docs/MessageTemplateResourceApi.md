# \MessageTemplateResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageTemplateUsingPOST**](MessageTemplateResourceApi.md#CreateMessageTemplateUsingPOST) | **Post** /api/message-templates | createMessageTemplate
[**DeleteMessageTemplateUsingDELETE**](MessageTemplateResourceApi.md#DeleteMessageTemplateUsingDELETE) | **Delete** /api/message-templates/{id} | deleteMessageTemplate
[**GetAllMessageTemplatesUsingGET**](MessageTemplateResourceApi.md#GetAllMessageTemplatesUsingGET) | **Get** /api/message-templates | getAllMessageTemplates
[**GetMessageTemplateUsingGET1**](MessageTemplateResourceApi.md#GetMessageTemplateUsingGET1) | **Get** /api/message-templates/{id} | getMessageTemplate
[**UpdateMessageTemplateUsingPUT**](MessageTemplateResourceApi.md#UpdateMessageTemplateUsingPUT) | **Put** /api/message-templates | updateMessageTemplate


# **CreateMessageTemplateUsingPOST**
> MessageTemplate CreateMessageTemplateUsingPOST(ctx, messageTemplate)
createMessageTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **messageTemplate** | [**MessageTemplate**](MessageTemplate.md)| messageTemplate | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessageTemplateUsingDELETE**
> DeleteMessageTemplateUsingDELETE(ctx, id)
deleteMessageTemplate

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

# **GetAllMessageTemplatesUsingGET**
> []MessageTemplate GetAllMessageTemplatesUsingGET(ctx, )
getAllMessageTemplates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MessageTemplate**](MessageTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageTemplateUsingGET1**
> MessageTemplate GetMessageTemplateUsingGET1(ctx, id)
getMessageTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMessageTemplateUsingPUT**
> MessageTemplate UpdateMessageTemplateUsingPUT(ctx, messageTemplate)
updateMessageTemplate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **messageTemplate** | [**MessageTemplate**](MessageTemplate.md)| messageTemplate | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

