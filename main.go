package main

import (
	"fmt"
	"time"

	"local.dev.com/config"
	"local.dev.com/services/kural"
	"local.dev.com/utils"
)

var appSettings *config.Config

func init() {
	appSettings = &config.Config{}
	appSettings.GetAll()

	if appSettings.MJ_PUBLIC_KEY == "" || appSettings.MJ_SECRET_KEY == "" {
		panic("mail jet keys are empty")
	}

	if appSettings.MJ_MAIL_SENDER == "" || len(appSettings.RECIPIENTS) == 0 {
		panic("mail jet sender or recipients are empty")
	}
}

func main() {
	kural, err := kural.GetDailyKural(appSettings, utils.RandomNumber(), "tamil")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Kural Number: %d, Ran successfully on %s", kural.Number, time.Now())
	}
}

// TODO: Apply Visitor Pattern for translator
// TODO: Add Notifier
