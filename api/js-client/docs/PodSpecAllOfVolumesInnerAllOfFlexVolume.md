# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfFlexVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**driver** | **String** | driver is the name of the driver to use for this volume. | [default to &#39;&#39;]
**fsType** | **String** | fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. The default filesystem depends on FlexVolume script. | [optional] 
**options** | **{String: String}** | options is Optional: this field holds extra command options if any. | [optional] 
**readOnly** | **Boolean** | readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**secretRef** | [**PodSpecAllOfVolumesInnerAllOfFlexVolumeAllOfSecretRef**](PodSpecAllOfVolumesInnerAllOfFlexVolumeAllOfSecretRef.md) |  | [optional] 


