package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Notification struct {
	Level                string `json:"level"`
	EmployeeAbbreviation string `json:"employeeAbbreviation"`
	Message              string `json:"message"`
}

func Notify(notification Notification) error {
	url := "http://localhost:8080/api/notify"
	payload, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
