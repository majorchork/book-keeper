package services

import (
	"fmt"
	"github.com/majorchork/book-keeper/internal/models"
	"github.com/majorchork/book-keeper/internal/ports"
	"github.com/majorchork/book-keeper/pkg"
)

type ComputerService struct {
	repo ports.Repository
}

func NewComputerService(repo ports.Repository) *ComputerService {
	return &ComputerService{repo: repo}
}

func (s *ComputerService) CreateComputer(computer *models.Computer) error {
	createErr := s.repo.CreateComputer(computer)
	if createErr != nil {
		return fmt.Errorf("could not create computer: %v", createErr)
	}

	notifyErr := s.CheckAndNotifyEmployeeComputerCount(computer.EmployeeAbbr)
	if notifyErr != nil {
		return notifyErr
	}

	return nil
}

func (s *ComputerService) GetAllComputers() ([]models.Computer, error) {
	return s.repo.GetAllComputers()
}

func (s *ComputerService) GetComputersByEmployee(employeeAbbr string) ([]models.Computer, error) {
	return s.repo.GetComputersByEmployee(employeeAbbr)
}

func (s *ComputerService) GetComputerByID(id uint) (*models.Computer, error) {
	return s.repo.GetComputerByID(id)
}

func (s *ComputerService) UpdateComputer(computer *models.Computer) error {
	return s.repo.UpdateComputer(computer)
}

func (s *ComputerService) DeleteComputer(id uint) error {
	return s.repo.DeleteComputer(id)
}

func (s *ComputerService) AssignComputer(id uint, employeeAbbr string) error {
	computer, err := s.repo.GetComputerByID(id)
	if err != nil {
		return fmt.Errorf("could not find computer by id: %v", err)
	}

	computer.EmployeeAbbr = employeeAbbr

	if updateErr := s.repo.UpdateComputer(computer); updateErr != nil {
		return fmt.Errorf("could not update computer: %v", updateErr)
	}

	notifyErr := s.CheckAndNotifyEmployeeComputerCount(computer.EmployeeAbbr)
	if notifyErr != nil {
		return fmt.Errorf("could not notify computer admin: %v", notifyErr)
	}

	return nil
}

func (s *ComputerService) CheckAndNotifyEmployeeComputerCount(employeeAbbr string) error {
	computers, err := s.repo.GetComputersByEmployee(employeeAbbr)
	if err != nil {
		return fmt.Errorf("could not get computers by employee: %v", err)
	}

	if len(computers) >= 3 {
		notification := pkg.Notification{
			Level:                "warning",
			EmployeeAbbreviation: employeeAbbr,
			Message:              "Employee has 3 or more computers assigned",
		}

		return pkg.Notify(notification)
	}

	return nil
}
