package main

import (
	"fmt"
	"strings"

	"github.com/fibonax/plantnet"
	"github.com/fibonax/treetify/utils"
	"github.com/fibonax/wikipedia"
	"github.com/michlabs/fbbot"
	log "github.com/sirupsen/logrus"
)

// NewTreetify creates Treetify with given apikey and imagesDir
func NewTreetify(apikey, imagesDir string) *Treetify {
	var tree Treetify
	tree.engine = plantnet.New(apikey)
	tree.imagesDir = imagesDir
	return &tree
}

// Treetify identifies a tree by a given image
type Treetify struct {
	engine    *plantnet.Plantnet
	imagesDir string
}

// HandleMessage handles received messsages
func (e Treetify) HandleMessage(bot *fbbot.Bot, msg *fbbot.Message) {
	bot.TypingOn(msg.Sender)

	var imgs []plantnet.Image
	for i, image := range msg.Images {
		var img plantnet.Image
		img.URL = image.URL
		img.Organ = plantnet.LeafOrgan
		location, err := utils.Download(img.URL, e.imagesDir, fmt.Sprintf("%s_%d.jpg", msg.ID, i))
		if err != nil {
			log.Errorf("cannot download %s to %s: %v", img.URL, e.imagesDir, err)
		} else {
			img.Location = location
			imgs = append(imgs, img)
		}
	}
	if len(imgs) > 0 {
		r, err := e.engine.IdentifyByImages(imgs...)
		if err != nil {
			bot.SendText(msg.Sender, "Lỗi rồi..."+err.Error())
		} else {
			log.Debugf("%+v", r)
			best := r.Results[0]
			sciname := best.Species.ScientificNameWithoutAuthor
			summary, err := wikipedia.GetSummary(sciname)
			if err == wikipedia.ErrNotFound {
				log.Infof("not found %s", sciname)
				bot.SendText(msg.Sender, "Tên khoa học là: "+sciname)
				if len(best.Species.CommonNames) > 0 {
					bot.SendText(msg.Sender, "Còn có các tên thường gọi là: ")
					for i, name := range best.Species.CommonNames {
						bot.SendText(msg.Sender, fmt.Sprintf("%d. %s", i+1, name))
					}
				}
				bot.SendText(msg.Sender, fmt.Sprintf("Chi tiết: https://www.google.com/search?hl=vi&q=%s", strings.ReplaceAll(best.Species.ScientificNameWithoutAuthor, " ", "+")))
				return
			}
			if err != nil {
				bot.SendText(msg.Sender, "Lỗi rồi..."+err.Error())
				return
			}
			bot.SendText(msg.Sender, "Đây là "+summary.Title)
			bot.SendText(msg.Sender, summary.Description)
			bot.SendText(msg.Sender, summary.Extract)
			bot.SendText(msg.Sender, fmt.Sprintf("Chi tiết: https://www.google.com/search?hl=vi&q=%s", strings.ReplaceAll(summary.Title, " ", "+")))
		}
	} else {
		bot.SendText(msg.Sender, "Chẳng thấy cái ảnh nào cả")
	}
}
