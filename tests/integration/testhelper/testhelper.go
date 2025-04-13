package testhelper

import (
	"ams-service/internal/core/services"
	"ams-service/internal/ports/primary"
	"database/sql"
	"os"
	"testing"
)

// SetupTestEnv sets up the test environment
func SetupTestEnv(t *testing.T) {
	// Set test environment variables
	os.Setenv("ENVIRONMENT", "test")

	// Add any other test environment setup here
}

// TeardownTestEnv cleans up the test environment
func TeardownTestEnv(t *testing.T) {
	// Clean up any test resources
	os.Unsetenv("ENVIRONMENT")

	// Add any other cleanup here
}

// SkipIfShort skips the test if the short flag is set
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
}

func InitDB(t *testing.T) *sql.DB {
	// Initialize database connection with hardcoded values
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=12345678 dbname=ams-local sslmode=disable")
	if err != nil {
		t.Fatal("Failed to connect to database:", err)
	}

	return db
}

func SetupTokenService() primary.TokenService {
	return services.NewTokenService("verySecureSystem")
}
