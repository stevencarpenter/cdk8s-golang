//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIngressList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeIngressList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIngressList_ManifestParameters(props *KubeIngressListProps) error {
	return nil
}

func validateKubeIngressList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIngressListParameters(scope constructs.Construct, id *string, props *KubeIngressListProps) error {
	return nil
}
