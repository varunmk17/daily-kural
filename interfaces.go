package main

import "local.dev.com/services/kural"

type Subscription interface {
	Subscribe(subscriber Subscriber)
	Unsubscribe(subscriber Subscriber)
	Notify(kural *kural.Kural)
}

type Subscriber interface {
	GetNotification(kural *kural.Kural)
}
