package cloudsnorkelcdkturbolayers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-turbo-layers.DependencyPackagerProps",
		reflect.TypeOf((*DependencyPackagerProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cloudsnorkel/cdk-turbo-layers.DependencyPackagerType",
		reflect.TypeOf((*DependencyPackagerType)(nil)).Elem(),
		map[string]interface{}{
			"LAMBDA": DependencyPackagerType_LAMBDA,
			"CODEBUILD": DependencyPackagerType_CODEBUILD,
		},
	)
	_jsii_.RegisterStruct(
		"@cloudsnorkel/cdk-turbo-layers.LayerProps",
		reflect.TypeOf((*LayerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-turbo-layers.NodejsDependencyPackager",
		reflect.TypeOf((*NodejsDependencyPackager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromPackageJson", GoMethod: "LayerFromPackageJson"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromYarn", GoMethod: "LayerFromYarn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NodejsDependencyPackager{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-turbo-layers.PythonDependencyPackager",
		reflect.TypeOf((*PythonDependencyPackager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromInline", GoMethod: "LayerFromInline"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromPipenv", GoMethod: "LayerFromPipenv"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromPoetry", GoMethod: "LayerFromPoetry"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromRequirementsTxt", GoMethod: "LayerFromRequirementsTxt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PythonDependencyPackager{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cloudsnorkel/cdk-turbo-layers.RubyDependencyPackager",
		reflect.TypeOf((*RubyDependencyPackager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "layerFromBundler", GoMethod: "LayerFromBundler"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RubyDependencyPackager{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
}
