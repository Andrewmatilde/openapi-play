# K8SResourceApi.DefaultApi

All URIs are relative to *http://127.0.0.1:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getResource**](DefaultApi.md#getResource) | **POST** /resource | 根据 GVK + namespace + name 获取对应的 K8s 资源



## getResource

> GetResource200Response getResource(resourceRequest)

根据 GVK + namespace + name 获取对应的 K8s 资源

### Example

```javascript
import K8SResourceApi from 'k8_s_resource_api';

let apiInstance = new K8SResourceApi.DefaultApi();
let resourceRequest = new K8SResourceApi.ResourceRequest(); // ResourceRequest | 
apiInstance.getResource(resourceRequest, (error, data, response) => {
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
 **resourceRequest** | [**ResourceRequest**](ResourceRequest.md)|  | 

### Return type

[**GetResource200Response**](GetResource200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

