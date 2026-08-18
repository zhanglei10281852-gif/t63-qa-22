package auth

import (
	"context"
	"testing"
)

func TestAuthenticationStopsAfterRequestCancellation(t *testing.T) {
	service, setupCtx := testService(t)
	if err := service.Bootstrap(setupCtx, "admin", "correct-password", "Administrator"); err != nil {
		t.Fatal(err)
	}
	login, err := service.Login(setupCtx, "admin", "correct-password")
	if err != nil {
		t.Fatal(err)
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if principal, err := service.Authenticate(cancelled, login.Token); err == nil {
		t.Fatalf("cancelled authentication returned principal %+v", principal)
	}
}
