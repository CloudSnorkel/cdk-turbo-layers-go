package cloudsnorkelcdkturbolayers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/CloudSnorkel/cdk-turbo-layers-go/cloudsnorkelcdkturbolayers/internal"
)

// Packager for creating Lambda layers for Python dependencies in AWS.
//
// Nothing is done locally so this doesn't require Docker, doesn't download any packages and doesn't upload huge files to S3.
// Experimental.
type PythonDependencyPackager interface {
	constructs.Construct
	awsec2.IConnectable
	awsiam.IGrantable
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Create a layer for dependencies passed as an argument and installed with pip.
	// Experimental.
	LayerFromInline(id *string, requirements *[]*string, props *LayerProps) awslambda.LayerVersion
	// Create a layer for dependencies defined in Pipfile and (optionally) Pipfile.lock and installed with pipenv.
	// Experimental.
	LayerFromPipenv(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Create a layer for dependencies defined in pyproject.toml and (optionally) poetry.lock and installed with poetry.
	// Experimental.
	LayerFromPoetry(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Create a layer for dependencies defined in requirements.txt and installed with pip.
	// Experimental.
	LayerFromRequirementsTxt(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PythonDependencyPackager
type jsiiProxy_PythonDependencyPackager struct {
	internal.Type__constructsConstruct
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_PythonDependencyPackager) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonDependencyPackager) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonDependencyPackager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonDependencyPackager(scope constructs.Construct, id *string, props *DependencyPackagerProps) PythonDependencyPackager {
	_init_.Initialize()

	if err := validateNewPythonDependencyPackagerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PythonDependencyPackager{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.PythonDependencyPackager",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonDependencyPackager_Override(p PythonDependencyPackager, scope constructs.Construct, id *string, props *DependencyPackagerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.PythonDependencyPackager",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PythonDependencyPackager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePythonDependencyPackager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-turbo-layers.PythonDependencyPackager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonDependencyPackager) LayerFromInline(id *string, requirements *[]*string, props *LayerProps) awslambda.LayerVersion {
	if err := p.validateLayerFromInlineParameters(id, requirements, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		p,
		"layerFromInline",
		[]interface{}{id, requirements, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonDependencyPackager) LayerFromPipenv(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := p.validateLayerFromPipenvParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		p,
		"layerFromPipenv",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonDependencyPackager) LayerFromPoetry(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := p.validateLayerFromPoetryParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		p,
		"layerFromPoetry",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonDependencyPackager) LayerFromRequirementsTxt(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := p.validateLayerFromRequirementsTxtParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		p,
		"layerFromRequirementsTxt",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonDependencyPackager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

