package cloudsnorkelcdkturbolayers


// Experimental.
type LayerProps struct {
	// Always rebuild the layer, even when the dependencies definition files haven't changed.
	// Default: false.
	//
	// Experimental.
	AlwaysRebuild *bool `field:"optional" json:"alwaysRebuild" yaml:"alwaysRebuild"`
}

