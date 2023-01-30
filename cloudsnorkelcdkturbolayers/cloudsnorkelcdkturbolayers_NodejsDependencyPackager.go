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

// Packager for creating Lambda layers for Node.js dependencies in AWS. Nothing is done locally so this doesn't require Docker, doesn't download any packages and doesn't upload huge files to S3.
// Experimental.
type NodejsDependencyPackager interface {
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
	// Create a layer for dependencies passed as an argument and installed with npm.
	// Experimental.
	LayerFromInline(id *string, libraries *[]*string, props *LayerProps) awslambda.LayerVersion
	// Create a layer for dependencies defined in package.json and (optionally) package-lock.json and installed with npm.
	// Experimental.
	LayerFromPackageJson(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Create a layer for dependencies defined in package.json and yarn.lock and installed with yarn.
	// Experimental.
	LayerFromYarn(id *string, path *string, props *LayerProps) awslambda.LayerVersion
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NodejsDependencyPackager
type jsiiProxy_NodejsDependencyPackager struct {
	internal.Type__constructsConstruct
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_NodejsDependencyPackager) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsDependencyPackager) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsDependencyPackager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodejsDependencyPackager(scope constructs.Construct, id *string, props *DependencyPackagerProps) NodejsDependencyPackager {
	_init_.Initialize()

	if err := validateNewNodejsDependencyPackagerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodejsDependencyPackager{}

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.NodejsDependencyPackager",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNodejsDependencyPackager_Override(n NodejsDependencyPackager, scope constructs.Construct, id *string, props *DependencyPackagerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cloudsnorkel/cdk-turbo-layers.NodejsDependencyPackager",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func NodejsDependencyPackager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodejsDependencyPackager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cloudsnorkel/cdk-turbo-layers.NodejsDependencyPackager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodejsDependencyPackager) LayerFromInline(id *string, libraries *[]*string, props *LayerProps) awslambda.LayerVersion {
	if err := n.validateLayerFromInlineParameters(id, libraries, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		n,
		"layerFromInline",
		[]interface{}{id, libraries, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodejsDependencyPackager) LayerFromPackageJson(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := n.validateLayerFromPackageJsonParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		n,
		"layerFromPackageJson",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodejsDependencyPackager) LayerFromYarn(id *string, path *string, props *LayerProps) awslambda.LayerVersion {
	if err := n.validateLayerFromYarnParameters(id, path, props); err != nil {
		panic(err)
	}
	var returns awslambda.LayerVersion

	_jsii_.Invoke(
		n,
		"layerFromYarn",
		[]interface{}{id, path, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodejsDependencyPackager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

