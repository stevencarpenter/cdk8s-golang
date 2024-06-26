package k8s

// ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
type ContainerResourceMetricSourceV2 struct {
	// container is the name of the container in the pods of the scaling target.
	Container *string `field:"required" json:"container" yaml:"container"`
	// name is the name of the resource in question.
	Name *string `field:"required" json:"name" yaml:"name"`
	// target specifies the target value for the given metric.
	Target *MetricTargetV2 `field:"required" json:"target" yaml:"target"`
}
