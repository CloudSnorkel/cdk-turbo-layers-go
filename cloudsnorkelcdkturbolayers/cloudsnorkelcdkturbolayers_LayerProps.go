// Speed-up Lambda function deployment with dependency layers built in AWS
package cloudsnorkelcdkturbolayers


// Experimental.
type LayerProps struct {
	// Always rebuild the layer, even when the depdencies definition files haven't changed.
	// Experimental.
	AlwaysRebuild *bool `field:"optional" json:"alwaysRebuild" yaml:"alwaysRebuild"`
}

