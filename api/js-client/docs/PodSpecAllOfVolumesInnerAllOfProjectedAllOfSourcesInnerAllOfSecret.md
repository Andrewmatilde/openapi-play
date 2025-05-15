# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfProjectedAllOfSourcesInnerAllOfSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**[PodSpecAllOfVolumesInnerAllOfConfigMapAllOfItemsInner]**](PodSpecAllOfVolumesInnerAllOfConfigMapAllOfItemsInner.md) | items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the &#39;..&#39; path or start with &#39;..&#39;. | [optional] 
**name** | **String** | Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] [default to &#39;&#39;]
**optional** | **Boolean** | optional field specify whether the Secret or its key must be defined | [optional] 


