package kural

import (
	"testing"

	"local.dev.com/config"
)

func TestGetKuralInEnglishReturnsValid(t *testing.T) {
	number := 42
	language := "english"
	appSettings := &config.Config{MJ_PUBLIC_KEY: "mock-key", MJ_SECRET_KEY: "mock-secret", MJ_MAIL_SENDER: "mock-sender@dev.com", RECIPIENTS: []string{"mock-recipient@dev.com"}}

	result, err := GetDailyKural(appSettings, number, language)

	// Check that the returned error is nil
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check that the returned Kural has the expected fields
	if result.Number != 42 {
		t.Errorf("Expected Kural number to be 42, but got %v", result.Number)
	}

	if result.Language != "english" {
		t.Errorf("Expected Kural language to be 'english', but got %v", result.Language)
	}

}

func TestGetKuralInTamilReturnsValid(t *testing.T) {
	number := 1233
	language := "tamil"
	appSettings := &config.Config{MJ_PUBLIC_KEY: "mock-key", MJ_SECRET_KEY: "mock-secret", MJ_MAIL_SENDER: "mock-sender@dev.com", RECIPIENTS: []string{"mock-recipient@dev.com"}}

	result, err := GetDailyKural(appSettings, number, language)

	// Check that the returned error is nil
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check that the returned Kural has the expected fields
	if result.Number != 1233 {
		t.Errorf("Expected Kural number to be 42, but got %v", result.Number)
	}

	if result.Language != "tamil" {
		t.Errorf("Expected Kural language to be 'tamil', but got %v", result.Language)
	}

}

func TestGetKuralReturnsBadRequest(t *testing.T) {
	number := 0
	language := "tamil"
	appSettings := &config.Config{MJ_PUBLIC_KEY: "mock-key", MJ_SECRET_KEY: "mock-secret", MJ_MAIL_SENDER: "mock-sender@dev.com", RECIPIENTS: []string{"mock-recipient@dev.com"}}

	result, err := GetDailyKural(appSettings, number, language)

	// Check that the returned error is not nil
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	// Check that the error message is correct
	expectedErrorMsg := "kural number should be between 1 and 1330"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%v', but got '%v'", expectedErrorMsg, err.Error())
	}
	// Check that the returned Kural is nil
	if result != nil {
		t.Errorf("Expected Kural to be nil, but got %v", result)
	}

}

func TestGetKuralReturnsBadRequestForInvalidNumber(t *testing.T) {
	number := 1331
	language := "english"
	appSettings := &config.Config{MJ_PUBLIC_KEY: "mock-key", MJ_SECRET_KEY: "mock-secret", MJ_MAIL_SENDER: "mock-sender@dev.com", RECIPIENTS: []string{"mock-recipient@dev.com"}}

	result, err := GetDailyKural(appSettings, number, language)

	// Check that the returned error is not nil
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	// Check that the error message is correct
	expectedErrorMsg := "kural number should be between 1 and 1330"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%v', but got '%v'", expectedErrorMsg, err.Error())
	}
	// Check that the returned Kural is nil
	if result != nil {
		t.Errorf("Expected Kural to be nil, but got %v", result)
	}
}

func TestGetKuralReturnsBadRequestForUnsupportedLanguage(t *testing.T) {
	number := 100
	language := "french"
	appSettings := &config.Config{MJ_PUBLIC_KEY: "mock-key", MJ_SECRET_KEY: "mock-secret", MJ_MAIL_SENDER: "mock-sender@dev.com", RECIPIENTS: []string{"mock-recipient@dev.com"}}

	result, err := GetDailyKural(appSettings, number, language)

	// Check that the returned error is not nil
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	// Check that the error message is correct
	expectedErrorMsg := "language french is currently not supported"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%v', but got '%v'", expectedErrorMsg, err.Error())
	}
	// Check that the returned Kural is nil
	if result != nil {
		t.Errorf("Expected Kural to be nil, but got %v", result)
	}

}
