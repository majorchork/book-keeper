package api

import (
	"github.com/gin-gonic/gin"
	"github.com/majorchork/book-keeper/internal/models"
	"github.com/majorchork/book-keeper/internal/util"
	"net/http"
	"strconv"
)

func (u *HTTPHandler) CreateComputer(c *gin.Context) {
	var computer models.Computer
	if err := c.ShouldBindJSON(&computer); err != nil {
		util.Response(c, "bad request", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	if len(computer.EmployeeAbbr) != 3 && computer.EmployeeAbbr != "" {
		util.Response(c, "validation err", http.StatusBadRequest, nil, []string{"Employee abbreviation must be 3 characters"})
		return
	}

	if err := u.Service.CreateComputer(&computer); err != nil {
		util.Response(c, "failed to create computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	util.Response(c, "computer created", http.StatusCreated, computer, nil)
}

func (u *HTTPHandler) ViewAllComputers(c *gin.Context) {
	computers, err := u.Service.GetAllComputers()
	if err != nil {
		util.Response(c, "failed to get computers", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	util.Response(c, "success", http.StatusOK, computers, nil)
}

func (u *HTTPHandler) ViewEmployeeComputers(c *gin.Context) {
	employeeAbbr := c.Param("abbr")

	computers, err := u.Service.GetComputersByEmployee(employeeAbbr)
	if err != nil {
		util.Response(c, "failed to get computers", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	util.Response(c, "success", http.StatusOK, computers, nil)
}

func (u *HTTPHandler) ViewComputerInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Response(c, "invalid ID", http.StatusBadRequest, nil, []string{"Invalid ID"})
		return
	}

	computer, getErr := u.Service.GetComputerByID(uint(id))
	if getErr != nil {
		util.Response(c, "failed to get computer", http.StatusInternalServerError, nil, []string{getErr.Error()})
		return
	}

	util.Response(c, "success", http.StatusOK, computer, nil)
}

func (u *HTTPHandler) DeleteComputer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.Response(c, "invalid ID", http.StatusBadRequest, nil, []string{"Invalid ID"})
		return
	}

	if deleteErr := u.Service.DeleteComputer(uint(id)); deleteErr != nil {
		util.Response(c, "failed to delete computer", http.StatusInternalServerError, nil, []string{deleteErr.Error()})
		return
	}
	util.Response(c, "computer deleted", http.StatusOK, nil, nil)
}

func (u *HTTPHandler) AssignComputer(c *gin.Context) {
	var input struct {
		ComputerID   int    `json:"computer_id"`
		EmployeeAbbr string `json:"employee_abbr"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		util.Response(c, "invalid request", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	computer, err := u.Service.GetComputerByID(uint(input.ComputerID))
	if err != nil {
		util.Response(c, "failed to get computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	computer.EmployeeAbbr = input.EmployeeAbbr

	if err := u.Service.UpdateComputer(computer); err != nil {
		util.Response(c, "failed to assign computer", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	util.Response(c, "computer assigned", http.StatusOK, computer, nil)
}
