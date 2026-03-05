package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HISPatientResponse struct {
	FirstNameTH  string `json:"first_name_th"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th"`
	FirstNameEN  string `json:"first_name_en"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en"`
	DateOfBirth  string `json:"date_of_birth"`
	PatientHN    string `json:"patient_hn"`
	NationalID   string `json:"national_id"`
	PassportID   string `json:"passport_id"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
}

type HISClient struct {
	BaseURL string
}

func NewHISClient(baseURL string) *HISClient {
	return &HISClient{BaseURL: baseURL}
}

func (c *HISClient) SearchPatient(id, baseUrl string) (*HISPatientResponse, error) {
	base := strings.TrimSpace(baseUrl)
	if base == "" {
		base = c.BaseURL
	}

	url := fmt.Sprintf("%s/patient/search/%s", base, id)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result HISPatientResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
