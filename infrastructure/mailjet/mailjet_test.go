package mailjet

import "testing"

func TestInitialize(t *testing.T) {
	publicKey := "test_public_key"
	secretKey := "test_private_key"

	emailer := EmailNotifier{}
	emailer.Initialize(publicKey, secretKey)

	if publicKey != emailer.Client.APIKeyPublic() {
		t.Errorf("expected %s but got %s", publicKey, emailer.Client.APIKeyPublic())
	}

	if secretKey != emailer.Client.APIKeyPrivate() {
		t.Errorf("expected %s but got %s", secretKey, emailer.Client.APIKeyPrivate())
	}
}

func TestSendSuccess(t *testing.T) {
	publicKey := "2c3500e6b2a1006951eea54dfc4a75fb"
	secretKey := "6dcca9ea4e6793652236f15eb8929bd1"

	emailer := EmailNotifier{}
	emailer.Initialize(publicKey, secretKey)

	fromEmail := "mock-sender@dev.com"
	toEmail := "mock-recipient@dev.com"
	err := emailer.Send("!!daily kural!!", fromEmail, toEmail)
	if err != nil {
		t.Errorf("Expected no error")
	}
}

func TestSendFailure(t *testing.T) {
	publicKey := "invalid-key"
	secretKey := "invalid-key"

	emailer := EmailNotifier{}
	emailer.Initialize(publicKey, secretKey)

	fromEmail := "mock-sender@dev.com"
	toEmail := "mock-recipient@dev.com"
	err := emailer.Send("!!daily kural!!", fromEmail, toEmail)
	if err == nil {
		t.Errorf("Expected api key authentication error")
	}
}
