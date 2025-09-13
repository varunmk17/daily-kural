package main

import (
	"fmt"

	"io"
	"log"
	"net/http"
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
	htmlMessage, err := renderRemoteTemplate(dailyKural)
	if err != nil {
		htmlMessage, err = renderLocalTemplate(dailyKural)
		if err != nil {
			htmlMessage = fmt.Sprintf("%s: <br/> %s <br/> <br/> %s: <br/>", dailyKural.Headers.HeaderKural, dailyKural.Kural, dailyKural.Headers.HeaderExplanation)

			for _, urai := range dailyKural.Urai {
				htmlMessage = fmt.Sprintf("%s <br/> %s - %s <br/>", htmlMessage, urai.Explanation, urai.Author)
			}
		}
	}
	mailer.Send(htmlMessage, appSettings.MJ_MAIL_SENDER, s.email)
}

func renderLocalTemplate(k *kural.Kural) (string, error) {
    tmpl, err := template.New("email_template.html").
        Funcs(template.FuncMap{
            "safeHTML": func(s string) template.HTML {
                return template.HTML(s) // keeps <br/> intact
            },
        }).
        ParseFiles("email_template.html")
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, k); err != nil {
        return "", err
    }

    return buf.String(), nil
}

func renderRemoteTemplate(k *kural.Kural) (string, error) {
	resp, err := http.Get("https://storage.googleapis.com/daily-kural/email_template.html")
	if err != nil {
		log.Fatalf("failed to fetch template: %v", err)
	}
	defer resp.Body.Close()
	tplBytes, _ := io.ReadAll(resp.Body)

	tmpl, err := template.New("email").Funcs(template.FuncMap{
            "safeHTML": func(s string) template.HTML {
                return template.HTML(s) // keeps <br/> intact
            },
        }).Parse(string(tplBytes))
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, k)
	htmlContent := buf.String()

	return htmlContent, nil
}

