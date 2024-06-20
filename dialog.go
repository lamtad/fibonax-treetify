package main

import (
	"github.com/fibonax/treetify/events"
	"github.com/fibonax/treetify/steps"
	"github.com/michlabs/fbbot"
)

func NewDialog(apikey, imagesDir string) *fbbot.Dialog {
	var greetingstep steps.Greeting
	var flowerorleafstep steps.FlowerOrLeaf
	var treetifystep steps.Treetify = steps.NewTreetify(apikey, imagesDir)
	var errorstep steps.Error
	var goodbyestep steps.Goodbye
	d := fbbot.NewDialog()
	d.SetBeginStep(greetingstep)
	d.SetEndStep(goodbyestep)
	// d.AddTransition(DoneEvent, greetingstep, flowerorleafstep)
	d.AddTransition(events.GoFlowerOrLeaf, flowerorleafstep)
	d.AddTransition(events.Done, flowerorleafstep, treetifystep)
	d.AddTransition(events.Done, treetifystep, goodbyestep)
	d.AddTransition(events.Err, errorstep)
	d.AddTransition(events.Done, errorstep, flowerorleafstep)
	return d
}
