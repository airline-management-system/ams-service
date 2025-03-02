package entities

import (
	"time"
)

type Employee struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       string    `json:"employee_id" gorm:"unique;not null"`
	Name             string    `json:"name" gorm:"size:50;not null"`
	Surname          string    `json:"surname" gorm:"size:50;not null"`
	Email            string    `json:"email" gorm:"unique;size:100;not null"`
	Phone            string    `json:"phone" gorm:"size:15"`
	Address          string    `json:"address" gorm:"size:255"`
	Gender           string    `json:"gender" gorm:"type:enum('male', 'female', 'other');not null"`
	BirthDate        time.Time `json:"birth_date" gorm:"not null"`
	HireDate         time.Time `json:"hire_date" gorm:"not null"`
	Position         string    `json:"position" gorm:"size:100;not null"`
	Department       string    `json:"department" gorm:"size:100;not null;type:enum('Admin', 'Passenger Services', 'Human Resources','Ground Services',Flight Planner');not null')"`
	Salary           float64   `json:"salary" gorm:"type:decimal(10,2);not null"`
	Status           string    `json:"status" gorm:"type:enum('active', 'inactive', 'on_leave');default:'active';not null"`
	ManagerID        *uint     `json:"manager_id"`
	EmergencyContact string    `json:"emergency_contact" gorm:"size:100"`
	EmergencyPhone   string    `json:"emergency_phone" gorm:"size:15"`
	ProfileImageURL  string    `json:"profile_image_url" gorm:"size:255"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type GetEmployeeByIdRequest struct {
	ID uint `json:"id" binding:"required"`
}
