package steps

import (
	"github.com/fibonax/treetify/events"
	"github.com/michlabs/fbbot"
)

type Goodbye struct {
	fbbot.BaseStep
}

func (s Goodbye) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, T("goodbye"))
	return events.Done
}
