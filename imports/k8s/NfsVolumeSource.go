package k8s

// Represents an NFS mount that lasts the lifetime of a pod.
//
// NFS volumes do not support ownership management or SELinux relabeling.
type NfsVolumeSource struct {
	// path that is exported by the NFS server.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Path *string `field:"required" json:"path" yaml:"path"`
	// server is the hostname or IP address of the NFS server.
	//
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Server *string `field:"required" json:"server" yaml:"server"`
	// readOnly here will force the NFS export to be mounted with read-only permissions.
	//
	// Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// Default: false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}
