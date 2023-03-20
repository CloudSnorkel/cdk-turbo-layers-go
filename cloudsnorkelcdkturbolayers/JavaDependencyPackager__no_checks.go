//go:build no_runtime_type_checking

// Speed-up Lambda function deployment with dependency layers built in AWS
package cloudsnorkelcdkturbolayers

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaDependencyPackager) validateLayerFromMavenParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func validateJavaDependencyPackager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJavaDependencyPackagerParameters(scope constructs.Construct, id *string, props *DependencyPackagerProps) error {
	return nil
}

