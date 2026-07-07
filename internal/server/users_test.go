package server

import (
	"testing"
	"time"
)

func TestGetTokenExpiration_DefaultsToOneHour(t *testing.T) {
	got := getTokenExpiration(nil)
	if got != time.Hour {
		t.Fatalf("expected 1 hour, got %v", got)
	}
}

func TestGetTokenExpiration_CapsAtOneHour(t *testing.T) {
	expiresInSeconds := 7200
	got := getTokenExpiration(&expiresInSeconds)
	if got != time.Hour {
		t.Fatalf("expected 1 hour cap, got %v", got)
	}
}

func TestGetTokenExpiration_UsesRequestedValue(t *testing.T) {
	expiresInSeconds := 120
	got := getTokenExpiration(&expiresInSeconds)
	if got != 120*time.Second {
		t.Fatalf("expected 120 seconds, got %v", got)
	}
}
