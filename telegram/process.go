package telegram

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"time"
	"token-bot/config"
)

func Process() error {
	for {
		config.Load()
		latestId := config.App.LastUpdateID
		updates, err := GetUpdates(latestId)
		if err != nil {
			logrus.Errorf("update process error: %s", err)
		} else {
			for _, update := range updates {
				logrus.Debugf("update received: %s", spew.Sdump(update))
				err := ProcessUpdate(&update)
				if err != nil {
					logrus.Errorf("failed to process update: %s", err)
				}
				config.App.LastUpdateID = update.UpdateID + 1
				config.Save(config.App)
			}
		}
		time.Sleep(2 * time.Second)
	}
}
