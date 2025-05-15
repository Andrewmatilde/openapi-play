# UserApi.DefaultApi

All URIs are relative to *http://play:2000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getUser**](DefaultApi.md#getUser) | **POST** /user | 用户操作



## getUser

> User getUser(userNameActionRequest)

用户操作

### Example

```javascript
import UserApi from 'user_api';

let apiInstance = new UserApi.DefaultApi();
let userNameActionRequest = new UserApi.UserNameActionRequest(); // UserNameActionRequest | 
apiInstance.getUser(userNameActionRequest, (error, data, response) => {
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
 **userNameActionRequest** | [**UserNameActionRequest**](UserNameActionRequest.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

