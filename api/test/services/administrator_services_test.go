package test

import (
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/view"
	"testing"
)

func TestRegisterAdministrator(t *testing.T) {

	a := &view.Administrator{Email: "admin8@example.com", Password: "admin", Username: "admin8"}

	a, err := services.RegisterAdministrator(*a, false)

	if err != nil {
		t.Fatal("Error registering administrator", err)
	}

	if !a.NeedsPasswordReset {
		t.Fatal("Administrator needs password reset")
	}

	if a.Email != "admin8@example.com" {
		t.Fatal("Email not set correctly")
	}

}
