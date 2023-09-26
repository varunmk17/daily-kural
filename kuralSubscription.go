package main

import (
	"fmt"

	"local.dev.com/services/kural"

	emailer "local.dev.com/infrastructure/mailjet"
)

// Subscription
type KuralSubscription struct {
	subscribers []KuralSubscriber
}

func (k *KuralSubscription) Subscribe(subscriber KuralSubscriber) {
	k.subscribers = append(k.subscribers, subscriber)
}

func (k *KuralSubscription) UnSubscribe(subscriber KuralSubscriber) {
	// delete(k.subscribers, subscriber)
}

func (k *KuralSubscription) Notify(kural *kural.Kural) {
	mailer := emailer.EmailNotifier{}
	mailer.Initialize(appSettings.MJ_PUBLIC_KEY, appSettings.MJ_SECRET_KEY)

	for _, subscriber := range k.subscribers {
		subscriber.GetNotification(kural, mailer)
	}
}

// Subscriber
type KuralSubscriber struct {
	email string
}

func (s KuralSubscriber) GetNotification(dailyKural *kural.Kural, mailer emailer.EmailNotifier) {
	// TODO: Create Email template type driven by language
	message := fmt.Sprintf("%s: <br/> %s <br/> <br/> %s: <br/>%s", dailyKural.Headers.HeaderKural, dailyKural.Kural, dailyKural.Headers.HeaderExplanation, dailyKural.Explanation)
	mailer.Send(message, appSettings.MJ_MAIL_SENDER, s.email)
}
