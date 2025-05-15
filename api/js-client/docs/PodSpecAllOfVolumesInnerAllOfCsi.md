# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfCsi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**driver** | **String** | driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster. | [default to &#39;&#39;]
**fsType** | **String** | fsType to mount. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply. | [optional] 
**nodePublishSecretRef** | [**PodSpecAllOfVolumesInnerAllOfCsiAllOfNodePublishSecretRef**](PodSpecAllOfVolumesInnerAllOfCsiAllOfNodePublishSecretRef.md) |  | [optional] 
**readOnly** | **Boolean** | readOnly specifies a read-only configuration for the volume. Defaults to false (read/write). | [optional] 
**volumeAttributes** | **{String: String}** | volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver&#39;s documentation for supported values. | [optional] 


