package events

import (
	"github.com/michlabs/fbbot"
)

// Done tells a step has done its job
var Done fbbot.Event = "done"

// Err tells that something went wrong
var Err fbbot.Event = "on_error"

// GoFlowerOrLeaf tells that bot ready go to the step FlowerOrLeaf
var GoFlowerOrLeaf fbbot.Event = "flower_or_leaf"
