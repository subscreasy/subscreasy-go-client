# \OfferResourceApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOfferUsingPOST**](OfferResourceApi.md#CreateOfferUsingPOST) | **Post** /api/offers | createOffer
[**DeleteOfferUsingDELETE**](OfferResourceApi.md#DeleteOfferUsingDELETE) | **Delete** /api/offers/{id} | deleteOffer
[**GetAllOffersUsingGET**](OfferResourceApi.md#GetAllOffersUsingGET) | **Get** /api/offers | getAllOffers
[**GetOfferUsingGET**](OfferResourceApi.md#GetOfferUsingGET) | **Get** /api/offers/{id} | getOffer
[**UpdateOfferUsingPUT**](OfferResourceApi.md#UpdateOfferUsingPUT) | **Put** /api/offers | updateOffer


# **CreateOfferUsingPOST**
> Offer CreateOfferUsingPOST(ctx, offer)
createOffer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **offer** | [**Offer**](Offer.md)| offer | 

### Return type

[**Offer**](Offer.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOfferUsingDELETE**
> DeleteOfferUsingDELETE(ctx, id)
deleteOffer

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

# **GetAllOffersUsingGET**
> []Offer GetAllOffersUsingGET(ctx, )
getAllOffers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Offer**](Offer.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOfferUsingGET**
> Offer GetOfferUsingGET(ctx, id)
getOffer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| id | 

### Return type

[**Offer**](Offer.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOfferUsingPUT**
> Offer UpdateOfferUsingPUT(ctx, offer)
updateOffer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **offer** | [**Offer**](Offer.md)| offer | 

### Return type

[**Offer**](Offer.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

