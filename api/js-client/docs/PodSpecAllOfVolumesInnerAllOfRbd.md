# K8SResourceApi.PodSpecAllOfVolumesInnerAllOfRbd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fsType** | **String** | fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | [optional] 
**image** | **String** | image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [default to &#39;&#39;]
**keyring** | **String** | keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to &#39;/etc/ceph/keyring&#39;]
**monitors** | **[String]** | monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | 
**pool** | **String** | pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to &#39;rbd&#39;]
**readOnly** | **Boolean** | readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] 
**secretRef** | [**PodSpecAllOfVolumesInnerAllOfRbdAllOfSecretRef**](PodSpecAllOfVolumesInnerAllOfRbdAllOfSecretRef.md) |  | [optional] 
**user** | **String** | user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | [optional] [default to &#39;admin&#39;]


