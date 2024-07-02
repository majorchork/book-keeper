package repository

import (
	"github.com/majorchork/book-keeper/internal/models"
)

func (p *Postgres) CreateComputer(computer *models.Computer) error {
	return p.DB.Create(computer).Error
}

func (p *Postgres) GetAllComputers() ([]models.Computer, error) {
	var computers []models.Computer
	err := p.DB.Find(&computers).Error
	return computers, err
}

func (p *Postgres) GetComputersByEmployee(employeeAbbr string) ([]models.Computer, error) {
	var computers []models.Computer
	err := p.DB.Where("employee_abbr = ?", employeeAbbr).Find(&computers).Error
	return computers, err
}

func (p *Postgres) GetComputerByID(id uint) (*models.Computer, error) {
	var computer models.Computer
	err := p.DB.First(&computer, id).Error
	return &computer, err
}

func (p *Postgres) UpdateComputer(computer *models.Computer) error {
	return p.DB.Save(computer).Error
}

func (p *Postgres) DeleteComputer(id uint) error {
	return p.DB.Delete(&models.Computer{}, id).Error
}
