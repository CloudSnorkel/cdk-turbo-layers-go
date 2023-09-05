//go:build no_runtime_type_checking

package cloudsnorkelcdkturbolayers

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PythonDependencyPackager) validateLayerFromInlineParameters(id *string, requirements *[]*string, props *LayerProps) error {
	return nil
}

func (p *jsiiProxy_PythonDependencyPackager) validateLayerFromPipenvParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func (p *jsiiProxy_PythonDependencyPackager) validateLayerFromPoetryParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func (p *jsiiProxy_PythonDependencyPackager) validateLayerFromRequirementsTxtParameters(id *string, path *string, props *LayerProps) error {
	return nil
}

func validatePythonDependencyPackager_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPythonDependencyPackagerParameters(scope constructs.Construct, id *string, props *DependencyPackagerProps) error {
	return nil
}

