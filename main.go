package main

import (
	"fmt"
	"os"

	"github.com/fibonax/treetify/lobe"
	"github.com/fibonax/treetify/pushbullet"
	"github.com/fibonax/treetify/steps"
	"github.com/michlabs/fbbot"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("treetify is ready to Go")

	cfg, err := readConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}

	steps.Init(cfg.LanguageFile)

	if err := pushbullet.Init(cfg.PushbulletAPIKey, cfg.PushbulletDevice); err != nil {
		log.Fatalf("cannot init Pushbullet: %s", err)
	}
	lobe.Init(cfg.LobeEndpoint)

	if err := os.MkdirAll(cfg.ImagesDir, 0700); err != nil {
		log.Fatalf("cannot make dir %s: %v", cfg.ImagesDir, err)
	}
	// t := NewTreetify(cfg.PlantnetAPIKey, cfg.ImagesDir)
	t := NewDialog(cfg.PlantnetAPIKey, cfg.ImagesDir)

	bot := fbbot.New(cfg.Port, cfg.VerifyToken, cfg.AppSecret, cfg.PageAccessToken)
	bot.AddMessageHandler(t)
	bot.AddPostbackHandler(t)

	bot.Run()
}
