package k8s

// StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.
type StatusCause struct {
	// The field of the resource that has caused this error, as named by its JSON serialization.
	//
	// May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.
	//
	// Examples:
	// "name" - the field "name" on the current resource
	// "items[0].name" - the field "name" on the first array entry in "items"
	Field *string `field:"optional" json:"field" yaml:"field"`
	// A human-readable description of the cause of the error.
	//
	// This field may be presented as-is to a reader.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// A machine-readable description of the cause of the error.
	//
	// If this value is empty there is no information available.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}
