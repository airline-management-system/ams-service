package integration

import (
	"ams-service/internal/adapters/secondary/postgres"
	"ams-service/internal/core/entities"
	"ams-service/internal/core/services"
	"ams-service/internal/ports/secondary"
	"ams-service/tests/integration/testhelper"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmployeeServiceIntegration(t *testing.T) {
	// Skip if running in short mode
	testhelper.SkipIfShort(t)

	// Initialize database
	db := testhelper.InitDB(t)

	// Setup test dependencies
	repo := setupEmployeeRepository(db)
	tokenService := testhelper.SetupTokenService()

	// Create the service
	employeeService := services.NewEmployeeService(repo, tokenService)

	t.Run("should register and login employee successfully", func(t *testing.T) {
		// Test data
		employee := entities.Employee{
			EmployeeID:   "test1234",
			PasswordHash: "testPassword123",
			Name:         "Test",
			Surname:      "User",
			Email:        "test@example.com",
			Role:         "hr",
			Gender:       "male",
			BirthDate:    time.Now(),
			HireDate:     time.Now(),
			Position:     "HR Manager",
			Salary:       50000.00,
			Status:       "active",
		}

		// Register employee
		registerRequest := entities.RegisterEmployeeRequest{
			Employee: employee,
		}
		err := employeeService.RegisterEmployee(registerRequest)
		require.NoError(t, err, "Failed to register employee")

		// Login employee
		loggedInEmployee, token, err := employeeService.LoginEmployee(employee.EmployeeID, "testPassword123")
		require.NoError(t, err, "Failed to login employee")
		assert.NotEmpty(t, token, "Token should not be empty")
		assert.Equal(t, employee.EmployeeID, loggedInEmployee.EmployeeID, "Employee ID should match")
		assert.Equal(t, employee.Email, loggedInEmployee.Email, "Email should match")
	})

	t.Run("should get employee by ID", func(t *testing.T) {
		// Test data
		employeeID := "test1234"

		// Get employee
		request := entities.GetEmployeeByIdRequest{
			EmployeeID: employeeID,
		}
		employee, err := employeeService.GetEmployeeByID(request)
		require.NoError(t, err, "Failed to get employee by ID")
		assert.Equal(t, employeeID, employee.EmployeeID, "Employee ID should match")
	})

	t.Run("should fail login with wrong password", func(t *testing.T) {
		// Test data
		employeeID := "test1234"
		wrongPassword := "wrongPassword123"

		// Attempt login
		_, _, err := employeeService.LoginEmployee(employeeID, wrongPassword)
		assert.Error(t, err, "Should fail with wrong password")
	})

	// Cleanup
	t.Cleanup(func() {
		cleanupTestData(t, db)
	})
}

// Helper functions for test setup
func setupEmployeeRepository(db *sql.DB) secondary.EmployeeRepository {
	repo := postgres.NewEmployeeRepositoryImpl(db)
	return repo
}

func cleanupTestData(t *testing.T, db *sql.DB) {
	if _, err := db.Exec("DELETE FROM employees WHERE employee_id LIKE 'test%'"); err != nil {
		t.Log("Warning: Failed to clean up test data:", err)
	}
}
