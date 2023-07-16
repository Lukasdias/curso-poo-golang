package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiResponse struct {
	Gender string
	Name struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Dob      struct {
		Age  int       `json:"age"`
	} `json:"dob"`
	Phone string `json:"phone"`
}

const API_URL = "https://randomuser.me/api/?results=20&nat=br"

func GetRandomPeople() []ApiResponse {
	response, err := http.Get(API_URL)
	if err != nil {
		fmt.Printf("Failed to send GET request: %s", err.Error())
	}
	defer response.Body.Close()
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s", err.Error())
	}

	var result struct {
		Results []ApiResponse `json:"results"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Failed to parse JSON response: %s", err.Error())
	}

	return result.Results
}