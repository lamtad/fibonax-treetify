package steps

import (
	"fmt"
	"strings"

	"github.com/fibonax/plantnet"
	"github.com/fibonax/treetify/events"
	"github.com/fibonax/treetify/lobe"
	"github.com/fibonax/treetify/utils"
	"github.com/fibonax/wikipedia"
	"github.com/michlabs/fbbot"
)

// NewTreetify creates Treetify with given apikey and imagesDir
func NewTreetify(apikey, imagesDir string) Treetify {
	var tree Treetify
	tree.engine = plantnet.New(apikey)
	tree.imagesDir = imagesDir
	return tree
}

// Treetify identifies a tree by a given image
type Treetify struct {
	fbbot.BaseStep
	engine    *plantnet.Plantnet
	imagesDir string
}

func (s Treetify) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, T("sendme"))
	return fbbot.NilEvent
}

func (s Treetify) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.TypingOn(msg.Sender)

	var imgs []plantnet.Image
	for i, image := range msg.Images {
		var img plantnet.Image
		img.URL = image.URL
		img.Organ = bot.STMemory.For(msg.Sender.ID).Get("flower_or_leaf")
		location, err := utils.Download(img.URL, s.imagesDir, fmt.Sprintf("%s_%d.jpg", msg.ID, i))
		if err != nil {
			bot.Logger.Errorf("cannot download %s to %s: %v", img.URL, s.imagesDir, err)
		} else {
			img.Location = location
			imgs = append(imgs, img)
		}
	}
	if len(imgs) == 0 {
		bot.SendText(msg.Sender, T("no_image"))
		return fbbot.NilEvent
	}

	var sciname string
	switch bot.STMemory.For(msg.Sender.ID).Get("flower_or_leaf") {
	case "fruit":
		sciname = lobe.IdentifyByImages(imgs[0]).Name
	default:
		r, err := s.engine.IdentifyByImages(imgs...)
		if err != nil {
			bot.Logger.Error(err)
			return events.Err
		}
		best := r.Results[0]
		sciname = best.Species.ScientificNameWithoutAuthor
	}

	summary, err := wikipedia.GetSummary(sciname)

	if err == wikipedia.ErrNotFound {
		bot.SendText(msg.Sender, T("sci_name"))
		bot.SendText(msg.Sender, sciname)
		bot.SendText(msg.Sender, T("detail"))
		bot.SendText(msg.Sender, fmt.Sprintf("https://vi.wikipedia.org/wiki/%s", sciname))
		bot.SendText(msg.Sender, fmt.Sprintf("https://www.google.com/search?hl=vi&q=%s", strings.ReplaceAll(sciname, " ", "+")))
		return events.Done

		// bot.Logger.Infof("not found %s", sciname)
		// bot.SendText(msg.Sender, T("sci_name"))
		// bot.SendText(msg.Sender, sciname)
		// if len(best.Species.CommonNames) > 0 {
		// 	bot.SendText(msg.Sender, T("common_name"))
		// 	for i, name := range best.Species.CommonNames {
		// 		bot.SendText(msg.Sender, fmt.Sprintf("%d. %s", i+1, name))
		// 	}
		// }
		// bot.SendText(msg.Sender, T("detail"))
		// bot.SendText(msg.Sender, fmt.Sprintf("https://www.google.com/search?hl=vi&q=%s", strings.ReplaceAll(best.Species.ScientificNameWithoutAuthor, " ", "+")))
		// return events.Done
	}

	if err != nil {
		bot.Logger.Error(err)
		return events.Err
	}

	bot.SendText(msg.Sender, T("name")+" "+summary.Title)
	bot.SendText(msg.Sender, summary.Description)
	bot.SendText(msg.Sender, summary.Extract)
	bot.SendText(msg.Sender, T("detail"))
	bot.SendText(msg.Sender, fmt.Sprintf("https://vi.wikipedia.org/wiki/%s", sciname))
	bot.SendText(msg.Sender, fmt.Sprintf("https://www.google.com/search?hl=vi&q=%s", strings.ReplaceAll(summary.Title, " ", "+")))
	return events.Done
}
