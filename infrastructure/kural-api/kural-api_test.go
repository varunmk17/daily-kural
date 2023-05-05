package kuralapi

import (
	"testing"
)

func TestGetKuralReturnsKural(t *testing.T) {
	got, _ := GetKuralByNumber(1)

	if got == nil {
		t.Errorf("expected to be not empty")
	}
}

func TestGetKuralReturnsError(t *testing.T) {
	_, err := GetKuralByNumber(3000)

	if err == nil {
		t.Errorf("expected an error")
	}
}
