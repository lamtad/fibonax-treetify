package steps

import (
	"github.com/fibonax/treetify/events"
	"github.com/michlabs/fbbot"
)

type Error struct {
	fbbot.BaseStep
}

func (s Error) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, T("error"))
	return events.Done
}
