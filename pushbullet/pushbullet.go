package pushbullet

import (
	log "github.com/sirupsen/logrus"
	"github.com/xconstruct/go-pushbullet"
)

var dev *pushbullet.Device

func Init(apikey, device string) (err error) {
	pb := pushbullet.New(apikey)
	dev, err = pb.Device(device)
	if err != nil {
		return err
	}
	return nil
}

func Noti(title, message string) {
	if err := dev.PushNote(title, message); err != nil {
		log.Error(err)
	}
}
