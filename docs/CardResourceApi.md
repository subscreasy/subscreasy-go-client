# \CardResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardUsingPOST**](CardResourceApi.md#CreateCardUsingPOST) | **Post** /api/cards | createCard
[**DeleteCardUsingDELETE**](CardResourceApi.md#DeleteCardUsingDELETE) | **Delete** /api/cards/{id} | deleteCard
[**GetAllCardsUsingGET**](CardResourceApi.md#GetAllCardsUsingGET) | **Get** /api/cards | getAllCards
[**GetCardUsingGET**](CardResourceApi.md#GetCardUsingGET) | **Get** /api/cards/{id} | getCard
[**UpdateCardUsingPUT**](CardResourceApi.md#UpdateCardUsingPUT) | **Put** /api/cards | updateCard


# **CreateCardUsingPOST**
> SavedCard CreateCardUsingPOST(ctx, card)
createCard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **card** | [**SavedCard**](SavedCard.md)| card | 

### Return type

[**SavedCard**](SavedCard.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCardUsingDELETE**
> DeleteCardUsingDELETE(ctx, id)
deleteCard

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

# **GetAllCardsUsingGET**
> []SavedCard GetAllCardsUsingGET(ctx, )
getAllCards

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SavedCard**](SavedCard.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardUsingGET**
> SavedCard GetCardUsingGET(ctx, id)
getCard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**SavedCard**](SavedCard.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCardUsingPUT**
> SavedCard UpdateCardUsingPUT(ctx, card)
updateCard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **card** | [**SavedCard**](SavedCard.md)| card | 

### Return type

[**SavedCard**](SavedCard.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

