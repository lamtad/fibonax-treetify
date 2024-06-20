package steps

import (
	"github.com/fibonax/treetify/events"
	"github.com/michlabs/fbbot"
)

type Greeting struct {
	fbbot.BaseStep
}

func (s Greeting) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, T("greeting"))
	bot.SendText(msg.Sender, T("introduction"))
	return events.GoFlowerOrLeaf
}

func (s Greeting) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.Logger.Info("greeting process")
	return events.Done
}
