package steps

import (
	"github.com/michlabs/bottext"
)

// T is translator for application, a global object
var T bottext.BotTextFunc

func Init(langFile string) {
	bottext.MustLoad(langFile)
	T = bottext.New("vi")
}
