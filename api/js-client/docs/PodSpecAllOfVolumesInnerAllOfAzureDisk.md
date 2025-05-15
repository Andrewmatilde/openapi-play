# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfAzureDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cachingMode** | **String** | cachingMode is the Host Caching mode: None, Read Only, Read Write. | [optional] [default to &#39;ReadWrite&#39;]
**diskName** | **String** | diskName is the Name of the data disk in the blob storage | [default to &#39;&#39;]
**diskURI** | **String** | diskURI is the URI of data disk in the blob storage | [default to &#39;&#39;]
**fsType** | **String** | fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. | [optional] [default to &#39;ext4&#39;]
**kind** | **String** | kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared | [optional] [default to &#39;Shared&#39;]
**readOnly** | **Boolean** | readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] [default to false]


