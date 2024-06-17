package model

import (
	"gorm.io/gorm"

	"time"
)

type EmployeeClient struct {
    DB *gorm.DB
}

type Employee struct {
    gorm.Model
    FullName  string `json:"fullname"`
	Email string `json:"email" gorm:"uniqueIndex"`
    Password string
	CredentialID string `json:"credential_id"`
	Mobile string `json:"mobile"`
	BirthDate string `json:"birth_date"`
	Gender string `json:"gender"`
	Department string `json:"department"`
	Adress string `json:"adress"`
	AdmissionDate time.Time `json:"admission_date"`
	Reservations []Reservation `gorm:"foreignKey:EmployeeID"`
}

type EmployeeRepository interface {
	SaveEmployee(employee *Employee) error
	EmployeeFirst(query string, args ...interface{}) (*Employee, error)
	DeleteEmployee(employee *Employee) error
}

func (e EmployeeClient) SaveEmployee(employee *Employee) error {
	return e.DB.Save(employee).Error
}

func (e EmployeeClient) EmployeeFirst(query string, args ...interface{}) (*Employee, error) {
	var employee Employee
	if err := e.DB.Where(query, args...).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *EmployeeClient) DeleteEmployee(employee *Employee) error {
	return e.DB.Delete(employee).Error
}
