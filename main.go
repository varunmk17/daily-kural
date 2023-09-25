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
	// kuralSubscription := KuralSubscription{}
	
	// for _, email := range appSettings.RECIPIENTS {
	// 	subscriber := KuralSubscriber{email}
	// 	kuralSubscription.Subscribe(subscriber)
	// }

	// kuralSubscription.Notify()

	kural, err := kural.GetDailyKural(appSettings, utils.RandomNumber(), "tamil")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Kural Number: %d, Ran successfully on %s", kural.Number, time.Now())
	}
}

type Subscription interface {
	Subscribe(subscriber  Subscriber)
	Unsubscribe (subscriber Subscriber)
	Notify()
}

type Subscriber interface {
	GetNotification()
}

type KuralSubscription struct {
	subscribers []Subscriber
}

func (k KuralSubscription) Subscribe (subscriber KuralSubscriber) {
	k.subscribers = append(k.subscribers, subscriber)
}

func (k KuralSubscription) UnSubscribe (subscriber KuralSubscriber) {
	// delete(k.subscribers, subscriber)
}

func (k KuralSubscription) Notify() {
	for _, subscriber := range k.subscribers {
		subscriber.GetNotification()
	}
}

type KuralSubscriber struct {
	email string
}

func (s KuralSubscriber) GetNotification() {
	fmt.Println("got notification")
}


// TODO: Apply Visitor Pattern for translator
// TODO: Add Notifier
