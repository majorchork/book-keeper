package models

import "gorm.io/gorm"

type Computer struct {
	gorm.Model
	MacAddress   string `gorm:"unique;not null" json:"mac_address"`
	Name         string `gorm:"not null" json:"name"`
	IpAddress    string `gorm:"not null" json:"ip_address"`
	EmployeeAbbr string `json:"employee_abbr"`
	Description  string `json:"description"`
}
