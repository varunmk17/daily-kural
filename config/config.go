package config

import (
	"os"
	"strings"
)

type Config struct {
	MJ_PUBLIC_KEY    string
	MJ_SECRET_KEY    string
	MJ_MAIL_SENDER   string
	MONGODB_USERNAME string
	MONGODB_PASSWORD string
	RECIPIENTS       []string
}

func (c *Config) GetAll() {
	c.MJ_PUBLIC_KEY = os.Getenv("MJ_PUBLIC_KEY")
	c.MJ_SECRET_KEY = os.Getenv("MJ_SECRET_KEY")
	c.MJ_MAIL_SENDER = os.Getenv("MJ_MAIL_SENDER")
	c.MONGODB_USERNAME = os.Getenv("MONGODB_USERNAME")
	c.MONGODB_PASSWORD = os.Getenv("MONGODB_PASSWORD")
	c.RECIPIENTS = strings.Split(os.Getenv("RECIPIENTS"), ",")
}
