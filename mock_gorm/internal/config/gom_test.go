package config_test

import (
	. "golang_cheatsheet/mock_gorm/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestNewDatabase(t *testing.T) {
	mockViper := &MockViper{ // Asumsikan MockViper sudah diimplementasikan
		Config: map[string]string{
			"db_host": "localhost",
			"db_port": "5432",
			"db_user": "postgres",
			"db_pass": "postgres",
			"db_name": "postgres",
		},
	}
	mockOpener := new(MockOpener)
	mockDB := &gorm.DB{} // Buat atau konfigurasikan mock *gorm.DB sesuai kebutuhan

	// Atur perilaku mock
	mockOpener.On("Open", mock.Anything).Return(mockDB, nil)

	// Jalankan fungsi yang ingin diuji
	db, err := NewDatabase(mockViper, mockOpener)

	// Lakukan assertion
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Pastikan metode Open dipanggil dengan parameter yang benar
	mockOpener.AssertCalled(t, "Open", mock.Anything)
}

func TestOpen(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	opener := &GormOpener{}
	db, err := opener.Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
