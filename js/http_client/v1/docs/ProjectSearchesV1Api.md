# PolyaxonSdk.ProjectSearchesV1Api

Polyaxon sdk

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createProjectSearch**](ProjectSearchesV1Api.md#createProjectSearch) | **POST** /api/v1/{owner}/{project}/searches | Create project search
[**deleteProjectSearch**](ProjectSearchesV1Api.md#deleteProjectSearch) | **DELETE** /api/v1/{owner}/{entity}/searches/{uuid} | Delete project search
[**getProjectSearch**](ProjectSearchesV1Api.md#getProjectSearch) | **GET** /api/v1/{owner}/{entity}/searches/{uuid} | Get project search
[**listProjectSearchNames**](ProjectSearchesV1Api.md#listProjectSearchNames) | **GET** /api/v1/{owner}/{name}/searches/names | List project search names
[**listProjectSearches**](ProjectSearchesV1Api.md#listProjectSearches) | **GET** /api/v1/{owner}/{name}/searches | List project searches
[**patchProjectSearch**](ProjectSearchesV1Api.md#patchProjectSearch) | **PATCH** /api/v1/{owner}/{project}/searches/{search.uuid} | Patch project search
[**promoteProjectSearch**](ProjectSearchesV1Api.md#promoteProjectSearch) | **POST** /api/v1/{owner}/{entity}/searches/{uuid}/promote | Promote project search
[**updateProjectSearch**](ProjectSearchesV1Api.md#updateProjectSearch) | **PUT** /api/v1/{owner}/{project}/searches/{search.uuid} | Update project search



## createProjectSearch

> V1Search createProjectSearch(owner, project, body)

Create project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let project = "project_example"; // String | Project under namespace
let body = new PolyaxonSdk.V1Search(); // V1Search | Search body
apiInstance.createProjectSearch(owner, project, body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **project** | **String**| Project under namespace | 
 **body** | [**V1Search**](V1Search.md)| Search body | 

### Return type

[**V1Search**](V1Search.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteProjectSearch

> deleteProjectSearch(owner, entity, uuid)

Delete project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let entity = "entity_example"; // String | Entity: project name, hub name, registry name, ...
let uuid = "uuid_example"; // String | Uuid identifier of the sub-entity
apiInstance.deleteProjectSearch(owner, entity, uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **entity** | **String**| Entity: project name, hub name, registry name, ... | 
 **uuid** | **String**| Uuid identifier of the sub-entity | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getProjectSearch

> V1Search getProjectSearch(owner, entity, uuid)

Get project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let entity = "entity_example"; // String | Entity: project name, hub name, registry name, ...
let uuid = "uuid_example"; // String | Uuid identifier of the sub-entity
apiInstance.getProjectSearch(owner, entity, uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **entity** | **String**| Entity: project name, hub name, registry name, ... | 
 **uuid** | **String**| Uuid identifier of the sub-entity | 

### Return type

[**V1Search**](V1Search.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listProjectSearchNames

> V1ListSearchesResponse listProjectSearchNames(owner, name, opts)

List project search names

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let name = "name_example"; // String | Entity managing the resource
let opts = {
  'offset': 56, // Number | Pagination offset.
  'limit': 56, // Number | Limit size.
  'sort': "sort_example", // String | Sort to order the search.
  'query': "query_example", // String | Query filter the search.
  'bookmarks': true, // Boolean | Filter by bookmarks.
  'mode': "mode_example", // String | Mode of the search.
  'no_page': true // Boolean | No pagination.
};
apiInstance.listProjectSearchNames(owner, name, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **name** | **String**| Entity managing the resource | 
 **offset** | **Number**| Pagination offset. | [optional] 
 **limit** | **Number**| Limit size. | [optional] 
 **sort** | **String**| Sort to order the search. | [optional] 
 **query** | **String**| Query filter the search. | [optional] 
 **bookmarks** | **Boolean**| Filter by bookmarks. | [optional] 
 **mode** | **String**| Mode of the search. | [optional] 
 **no_page** | **Boolean**| No pagination. | [optional] 

### Return type

[**V1ListSearchesResponse**](V1ListSearchesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listProjectSearches

> V1ListSearchesResponse listProjectSearches(owner, name, opts)

List project searches

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let name = "name_example"; // String | Entity managing the resource
let opts = {
  'offset': 56, // Number | Pagination offset.
  'limit': 56, // Number | Limit size.
  'sort': "sort_example", // String | Sort to order the search.
  'query': "query_example", // String | Query filter the search.
  'bookmarks': true, // Boolean | Filter by bookmarks.
  'mode': "mode_example", // String | Mode of the search.
  'no_page': true // Boolean | No pagination.
};
apiInstance.listProjectSearches(owner, name, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **name** | **String**| Entity managing the resource | 
 **offset** | **Number**| Pagination offset. | [optional] 
 **limit** | **Number**| Limit size. | [optional] 
 **sort** | **String**| Sort to order the search. | [optional] 
 **query** | **String**| Query filter the search. | [optional] 
 **bookmarks** | **Boolean**| Filter by bookmarks. | [optional] 
 **mode** | **String**| Mode of the search. | [optional] 
 **no_page** | **Boolean**| No pagination. | [optional] 

### Return type

[**V1ListSearchesResponse**](V1ListSearchesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## patchProjectSearch

> V1Search patchProjectSearch(owner, project, search_uuid, body)

Patch project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let project = "project_example"; // String | Project under namespace
let search_uuid = "search_uuid_example"; // String | UUID
let body = new PolyaxonSdk.V1Search(); // V1Search | Search body
apiInstance.patchProjectSearch(owner, project, search_uuid, body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **project** | **String**| Project under namespace | 
 **search_uuid** | **String**| UUID | 
 **body** | [**V1Search**](V1Search.md)| Search body | 

### Return type

[**V1Search**](V1Search.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## promoteProjectSearch

> promoteProjectSearch(owner, entity, uuid)

Promote project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let entity = "entity_example"; // String | Entity: project name, hub name, registry name, ...
let uuid = "uuid_example"; // String | Uuid identifier of the sub-entity
apiInstance.promoteProjectSearch(owner, entity, uuid, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **entity** | **String**| Entity: project name, hub name, registry name, ... | 
 **uuid** | **String**| Uuid identifier of the sub-entity | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateProjectSearch

> V1Search updateProjectSearch(owner, project, search_uuid, body)

Update project search

### Example

```javascript
import PolyaxonSdk from 'polyaxon-sdk';
let defaultClient = PolyaxonSdk.ApiClient.instance;
// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new PolyaxonSdk.ProjectSearchesV1Api();
let owner = "owner_example"; // String | Owner of the namespace
let project = "project_example"; // String | Project under namespace
let search_uuid = "search_uuid_example"; // String | UUID
let body = new PolyaxonSdk.V1Search(); // V1Search | Search body
apiInstance.updateProjectSearch(owner, project, search_uuid, body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **String**| Owner of the namespace | 
 **project** | **String**| Project under namespace | 
 **search_uuid** | **String**| UUID | 
 **body** | [**V1Search**](V1Search.md)| Search body | 

### Return type

[**V1Search**](V1Search.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

