package test

import (
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/view"
	"testing"
)

func TestLoginInterviewer_SuccessfulLogin(t *testing.T) {

	interviewer := view.Interviewer{
		Email:    "test@example.com",
		Password: "password",
	}

	token, resetToken, err := services.LoginInterviewer(interviewer)

	if err != nil {
		t.Fatalf("Error inesperado durante el inicio de sesión: %v", err)
	}
	if token == nil {
		t.Fatal("Se esperaba un token de acceso no nulo")
	}
	if resetToken != nil {
		t.Fatal("Se esperaba un token de restablecimiento de contraseña nulo")
	}
}

func TestRegisterInterviewer(t *testing.T) {

	a := &view.Interviewer{Email: "interviewer20@example.com", Password: "interviewer20", Username: "itw20", FirstName: "John",
		LastName: "Doe"}

	a, err := services.RegisterInterviewer(*a)

	if err != nil {
		t.Fatal("Error registering interviewer", err)
	}

	if !a.NeedsPasswordReset {
		t.Fatal("Interviewer needs password reset")
	}

	if a.Email != "interviewer20@example.com" {
		t.Fatal("Email not set correctly")
	}

}
