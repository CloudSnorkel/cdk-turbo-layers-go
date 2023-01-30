// Speed-up Lambda function deployment with dependency layers built in AWS
package cloudsnorkelcdkturbolayers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Experimental.
type DependencyPackagerProps struct {
	// Target Lambda architecture.
	//
	// Packages will be installed for this architecture so make sure it fits your Lambda functions.
	// Experimental.
	Architecture awslambda.Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Removal policy for logs of image builds.
	//
	// If deployment fails on the custom resource, try setting this to `RemovalPolicy.RETAIN`. This way the CodeBuild logs can still be viewed, and you can see why the build failed.
	//
	// We try to not leave anything behind when removed. But sometimes a log staying behind is useful.
	// Experimental.
	LogRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"logRemovalPolicy" yaml:"logRemovalPolicy"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Additional commands to run before installing packages.
	//
	// Use this to authenticate your package repositories like CodeArtifact.
	// Experimental.
	PreinstallCommands *[]*string `field:"optional" json:"preinstallCommands" yaml:"preinstallCommands"`
	// Target Lambda runtime.
	//
	// Packages will be installed for this runtime so make sure it fits your Lambda functions.
	// Experimental.
	Runtime awslambda.Runtime `field:"optional" json:"runtime" yaml:"runtime"`
	// VPC subnets used for packager.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Type of dependency packager.
	//
	// Use Lambda for speed and CodeBuild for complex dependencies that require building native extensions.
	// Experimental.
	Type DependencyPackagerType `field:"optional" json:"type" yaml:"type"`
	// VPC used for packager.
	//
	// Use this if your package repositories are only available from within a VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

