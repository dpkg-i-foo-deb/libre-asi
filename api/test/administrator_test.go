package test

import (
	"libre-asi-api/pkg/models"
	"testing"

	"gorm.io/gorm"
)

func TestCreateAdministrator(t *testing.T) {
	// create a user for the administrator
	user := &models.User{
		Email:    "admin@example.com",
		Username: "admin",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}

	// create the administrator
	admin := &models.Administrator{
		UserID: user.ID,
	}
	if err := DB.Create(admin).Error; err != nil {
		t.Fatal(err)
	}

	// check if the administrator was created successfully
	var result models.Administrator
	if err := DB.First(&result, admin.ID).Error; err != nil {
		t.Fatal(err)
	}
	if result.UserID != admin.UserID {
		t.Errorf("expected userID=%d, got %d", admin.UserID, result.UserID)
	}
}

func TestUpdateAdministrator(t *testing.T) {
	// create a user for the administrator
	user := &models.User{
		Email:    "admin2@example.com",
		Username: "admin2",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}

	// create the administrator
	admin := &models.Administrator{
		UserID: user.ID,
	}
	if err := DB.Create(admin).Error; err != nil {
		t.Fatal(err)
	}

	// update the administrator
	newUserID := uint(123)
	admin.UserID = newUserID
	if err := DB.Save(admin).Error; err != nil {
		t.Fatal(err)
	}

	// check if the administrator was updated successfully
	var result models.Administrator
	if err := DB.First(&result, admin.ID).Error; err != nil {
		t.Fatal(err)
	}
	if result.UserID != newUserID {
		t.Errorf("expected userID=%d, got %d", newUserID, result.UserID)
	}
}

func TestDeleteAdministrator(t *testing.T) {
	// create a user for the administrator
	user := &models.User{
		Email:    "admin3@example.com",
		Username: "admin3",
		Password: "password",
	}
	if err := DB.Create(user).Error; err != nil {
		t.Fatal(err)
	}

	// create the administrator
	admin := &models.Administrator{
		UserID: user.ID,
	}
	if err := DB.Create(admin).Error; err != nil {
		t.Fatal(err)
	}

	// delete the administrator
	if err := DB.Delete(admin).Error; err != nil {
		t.Fatal(err)
	}

	// check if the administrator was deleted successfully
	var result models.Administrator
	if err := DB.First(&result, admin.ID).Error; err != gorm.ErrRecordNotFound {
		t.Errorf("expected record not found error, got %v", err)
	}

	// check if the specific administrator was deleted
	var count int64
	if err := DB.Model(&models.Administrator{}).Where("id = ?", admin.ID).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("expected 0 records, got %d", count)
	}
}
