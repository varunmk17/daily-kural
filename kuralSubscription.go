package main

import (
	"fmt"

	"bytes"
	"html/template"
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
	htmlMessage, err := renderEmail(dailyKural)
	if err != nil {
		htmlMessage = fmt.Sprintf("%s: <br/> %s <br/> <br/> %s: <br/>", dailyKural.Headers.HeaderKural, dailyKural.Kural, dailyKural.Headers.HeaderExplanation)

		for _, urai := range dailyKural.Urai {
			htmlMessage = fmt.Sprintf("%s <br/> %s - %s <br/>", htmlMessage, urai.Explanation, urai.Author)
		}
	}

	mailer.Send(htmlMessage, appSettings.MJ_MAIL_SENDER, s.email)
}

func renderEmail(k *kural.Kural) (string, error) {
	tmpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, k); err != nil {
		return "", err
	}

	return buf.String(), nil
}
