package steps

import (
	"github.com/fibonax/treetify/events"
	"github.com/michlabs/fbbot"
)

type FlowerOrLeaf struct {
	fbbot.BaseStep
}

func (s FlowerOrLeaf) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	m := fbbot.NewButtonMessage()
	m.Text = T("flower_or_leaf")
	m.AddPostbackButton("Hoa", "flower")
	m.AddPostbackButton("Lá", "leaf")
	m.AddPostbackButton("Quả", "fruit")

	bot.Send(msg.Sender, m)
	return fbbot.NilEvent
}

func (s FlowerOrLeaf) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.Logger.Infof("%s selected %s", msg.Sender.FirstName(), msg.Text)
	bot.STMemory.For(msg.Sender.ID).Set("flower_or_leaf", msg.Text)
	return events.Done
}
