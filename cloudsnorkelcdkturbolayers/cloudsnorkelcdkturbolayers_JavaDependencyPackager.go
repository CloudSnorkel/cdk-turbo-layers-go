// Speed-up Lambda function deployment with dependency layers built in AWS
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

// Packager for creating Lambda layers for Java dependencies in AWS.
//
// Nothing is done locally so this doesn't require Docker and doesn't upload huge files to S3.
// Experimental.
type JavaDependencyPackager interface {
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
	// Create a layer for dependencies defined in pom.xml installed with Maven.
	// Experimental.
	LayerFromMaven(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JavaDependencyPackager
type jsiiProxy_JavaDependencyPackager struct {
	internal.Type__constructsConstruct
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_JavaDependencyPackager) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaDependencyPackager) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaDependencyPackager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewJavaDependencyPackager(scope constructs.Construct, id *string, props *DependencyPackagerProps) JavaDependencyPackager {
	_init_.Initialize()

	if err := validateNewJavaDependencyPackagerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaDependencyPackager{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.JavaDependencyPackager",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJavaDependencyPackager_Override(j JavaDependencyPackager, scope constructs.Construct, id *string, props *DependencyPackagerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.JavaDependencyPackager",
		[]interface{}{scope, id, props},
		j,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func JavaDependencyPackager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJavaDependencyPackager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-turbo-layers.JavaDependencyPackager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaDependencyPackager) LayerFromMaven(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := j.validateLayerFromMavenParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		j,
		"layerFromMaven",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaDependencyPackager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

