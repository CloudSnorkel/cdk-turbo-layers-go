//go:build no_runtime_type_checking

// Speed-up Lambda function deployment with dependency layers built in AWS
package cloudsnorkelcdkturbolayers

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodejsDependencyPackager) validateLayerFromInlineParameters(id *string, libraries *[]*string, props *LayerProps) error {
	return nil
}

func (n *jsiiProxy_NodejsDependencyPackager) validateLayerFromPackageJsonParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func (n *jsiiProxy_NodejsDependencyPackager) validateLayerFromYarnParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func validateNodejsDependencyPackager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNodejsDependencyPackagerParameters(scope constructs.Construct, id *string, props *DependencyPackagerProps) error {
	return nil
}

