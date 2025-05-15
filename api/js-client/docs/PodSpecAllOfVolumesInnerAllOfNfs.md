# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfNfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**path** | **String** | path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | [default to &#39;&#39;]
**readOnly** | **Boolean** | readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | [optional] 
**server** | **String** | server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | [default to &#39;&#39;]


