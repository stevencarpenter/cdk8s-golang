package k8s

// Represents a Rados Block Device mount that lasts the lifetime of a pod.
//
// RBD volumes support ownership management and SELinux relabeling.
type RbdVolumeSource struct {
	// image is the rados image name.
	//
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image *string `field:"required" json:"image" yaml:"image"`
	// monitors is a collection of Ceph monitors.
	//
	// More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// fsType is the filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// keyring is the path to key ring for RBDUser.
	//
	// Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Default: etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	//
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// pool is the rados pool name.
	//
	// Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Default: rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	//
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	//
	// Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Default: false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretRef is name of the authentication secret for RBDUser.
	//
	// If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Default: nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	//
	SecretRef *LocalObjectReference `field:"optional" json:"secretRef" yaml:"secretRef"`
	// user is the rados user name.
	//
	// Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	// Default: admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	//
	User *string `field:"optional" json:"user" yaml:"user"`
}
