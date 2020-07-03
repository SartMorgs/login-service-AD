package main

import(
	//Go Packages
	"net/http"
	"testing"
	"encoding/json"
	"bytes"

	// External Packages

	// Internal Packages
)


func TestLoginAccessRouter(t *testing.T){
	requestBody, err := json.Marshal(map[string]string{
		"username": "teste",
		"password": "teste",
	})

	if err != nil{
		t.Errorf("Request Body Error")
	}

	resp, err := http.Post("http://localhost:8000/api/v1/access/", "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()

	if err != nil{
		t.Errorf("Expected nil, received %s", err. Error())
	}

	if resp.StatusCode != http.StatusOK{
		t.Errorf("Expected %d, received %d", http.StatusOK, resp.StatusCode)
	}	
}

func TestLogoutRouter(t *testing.T){
	requestBody, err := json.Marshal(map[string]string{
		"username": "username01",
	})

	if err != nil{
		t.Errorf("Request Body Error")
	}

	resp, err := http.Post("http://localhost:8000/api/v1/logout/", "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()

	if err != nil{
		t.Errorf("Expected nil, received %s", err. Error())
	}

	if resp.StatusCode != http.StatusOK{
		t.Errorf("Expected %d, received %d", http.StatusOK, resp.StatusCode)
	}
}

func TestOtherRouter(t *testing.T){
	resp, err := http.Get("http://localhost:8000/api/v1/other/")
	if err != nil{
		t.Errorf("Request problems")
	}

	defer resp.Body.Close()

	if err != nil{
		t.Errorf("Expected nil, received %s", err. Error())
	}

	if resp.StatusCode != http.StatusNotFound{
		t.Errorf("Expected %d, received %d", http.StatusNotFound, resp.StatusCode)
	}
}