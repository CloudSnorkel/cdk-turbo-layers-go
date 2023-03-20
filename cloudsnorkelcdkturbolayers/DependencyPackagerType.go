// Speed-up Lambda function deployment with dependency layers built in AWS
package cloudsnorkelcdkturbolayers


// Type of dependency packager.
//
// This affects timeouts and capabilities of the packager.
// Experimental.
type DependencyPackagerType string

const (
	// Use Lambda function to package dependencies.
	//
	// It is much faster than the alternative, but limited to 15 minutes and can't build native extensions.
	// Experimental.
	DependencyPackagerType_LAMBDA DependencyPackagerType = "LAMBDA"
	// Use CodeBuild to package dependencies.
	//
	// It is capable of everything your local machine can do, but takes a little longer to startup.
	// Experimental.
	DependencyPackagerType_CODEBUILD DependencyPackagerType = "CODEBUILD"
)

