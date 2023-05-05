package mailjet

import (
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type EmailNotifier struct {
	Client *mailjet.Client
}

func (mailer *EmailNotifier) Initialize(publicKey, secretKey string) {
	mailer.Client = mailjet.NewMailjetClient(publicKey, secretKey)
}

func (mailer *EmailNotifier) Send(message string, sender string, recipients []string) error {
	messagesInfo := []mailjet.InfoMessagesV31{}

	for i := 0; i < len(recipients); i++ {
		recipient := recipients[i]

		message := mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: sender,
				Name:  "Learning Workspace",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipient,
				},
			},
			Subject:  "Daily Kural",
			TextPart: message,
			HTMLPart: fmt.Sprintf("<h3>%s</h3><br />", message),
		}

		messagesInfo = append(messagesInfo, message)
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailer.Client.SendMailV31(&messages)
	if err != nil {
		return err
	}

	return nil
}
