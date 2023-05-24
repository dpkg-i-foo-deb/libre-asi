package test

import (
	"errors"
	"libre-asi-api/pkg/database"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/services"
	"libre-asi-api/pkg/view"
	"testing"

	"gorm.io/gorm"
)

// type loginCredentials struct {
// 	Email    string
// 	Password string
// }

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

func TestGetAdministrators(t *testing.T) {

	admin9 := &view.Administrator{Email: "admin9@example.com", Password: "admin9", Username: "admin9"}
	admin10 := &view.Administrator{Email: "admin10@example.com", Password: "admin10", Username: "admin10"}

	_, err := services.RegisterAdministrator(*admin9, false)
	if err != nil {
		t.Fatalf("Failed to register administrator: %v", err)
	}

	_, err = services.RegisterAdministrator(*admin10, false)
	if err != nil {
		t.Fatalf("Failed to register administrator: %v", err)
	}

	admins, err := services.GetAdministrators()

	if err != nil {
		t.Fatalf("Failed to get administrators: %v", err)
	}

	if len(admins) != 3 {
		t.Fatalf("Expected to get 2 administrators, got %d", len(admins))
	}
}

func TestUpdateAdministrator(t *testing.T) {

	administrators, err := services.GetAdministrators()
	if err != nil {
		t.Fatalf("Failed to get administrators: %v", err)
	}

	if len(administrators) == 0 {
		t.Fatal("No administrators found in the database")
	}

	admin := administrators[0]

	admin.Username = "admin_updated"

	updatedAdmin, err := services.UpdateAdministrator(admin)
	if err != nil {
		t.Fatalf("Failed to update administrator: %v", err)
	}

	if updatedAdmin.Username != "admin_updated" {
		t.Fatalf("Expected username to be 'admin_updated', got '%s'", updatedAdmin.Username)
	}

	admins, err := services.GetAdministrators()
	if err != nil {
		t.Fatalf("Failed to get administrators: %v", err)
	}

	if len(admins) != len(administrators) {
		t.Fatalf("Expected to get %d administrators, got %d", len(administrators), len(admins))
	}
}

func TestDeleteAdministrator(t *testing.T) {

	administrators, err := services.GetAdministrators()
	if err != nil {
		t.Fatalf("Failed to get administrators: %v", err)
	}

	if len(administrators) == 0 {
		t.Fatal("No administrators found in the database")
	}

	admin := administrators[0]

	err = services.DeleteAdministrator(int(admin.ID))
	if err != nil {
		t.Fatalf("Failed to delete administrator: %v", err)
	}
	var queriedAdmin models.Administrator
	if err := database.DB.Where("ID = ?", admin.ID).First(&queriedAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		t.Fatalf("Error retrieving deleted administrator: %v", err)
	}
}

// func TestSetAdministratorPassword(t *testing.T) {

// 	administrators, err := services.GetAdministrators()
// 	if err != nil {
// 		t.Fatalf("Failed to get administrators: %v", err)
// 	}

// 	var admin *view.Administrator
// 	if len(administrators) > 0 {
// 		admin = &administrators[0]
// 	} else {
// 		t.Fatal("No administrators found in the database")
// 	}

// 	credentials := models.PasswordChange{
// 		CurrentPassword: "admin",
// 		NewPassword:     "newpassword",
// 	}

// 	err = services.SetAdministratorPassword(admin.Email, credentials)
// 	if err != nil {
// 		t.Fatalf("Failed to set administrator password: %v", err)
// 	}

// 	err = services.SetAdministratorPassword(admin.Email, models.PasswordChange{
// 		CurrentPassword: credentials.NewPassword,
// 		NewPassword:     "admin",
// 	})
// 	if err != nil {
// 		t.Fatalf("Failed to restore administrator password: %v", err)
// 	}
// }
