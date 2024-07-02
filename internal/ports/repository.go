package ports

import "github.com/majorchork/book-keeper/internal/models"

type Repository interface {
	CreateComputer(computer *models.Computer) error
	GetAllComputers() ([]models.Computer, error)
	GetComputersByEmployee(employeeAbbr string) ([]models.Computer, error)
	GetComputerByID(id uint) (*models.Computer, error)
	UpdateComputer(computer *models.Computer) error
	DeleteComputer(id uint) error
}

type ComputerService interface {
	CreateComputer(computer *models.Computer) error
	GetAllComputers() ([]models.Computer, error)
	GetComputersByEmployee(employeeAbbr string) ([]models.Computer, error)
	GetComputerByID(id uint) (*models.Computer, error)
	UpdateComputer(computer *models.Computer) error
	DeleteComputer(id uint) error
	AssignComputer(id uint, employeeAbbr string) error
}
