package backend

import (
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"encoding/json"
	"fmt"
	"log"
)

func ChechAuthentication(baseURL string, user m.AuthUser, verbose bool) {
	if user.Username == "" || user.Password == "" {
		log.Fatal("username and password are required parameters")
	}

	url := fmt.Sprintf("%s/api/authentication/validate", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	respBody, _ := u.HTTPRequest(req, verbose)

	json.Unmarshal(respBody, &user)

	if user.Valid {
		log.Println("Authentication successful")
	} else {
		log.Fatal("Authentication failed. Please check your username and password")
	}
}

func CreateUser(baseURL string, user m.AuthUser, sonarUser m.SonarUser, verbose bool) {
	if sonarUser.Login == "" || sonarUser.Name == "" || sonarUser.Email == "" {
		log.Fatal("login, name and email are required parameters for creating a user")
	}

	url := fmt.Sprintf("%s/api/users/create", baseURL)

	reqBody, err := json.Marshal(sonarUser)
	u.Error(err, "")

	req := u.CreateBaseRequest("POST", url, reqBody, user, verbose)
	_, status := u.HTTPRequest(req, verbose)

	if status == "200 OK" {
		log.Printf("User with login='%s' is created", sonarUser.Login)
	} else if status == "400 BAD REQUEST" {
		log.Printf("An active user with login='%s' already exists", sonarUser.Login)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to create a new user", user.Username)
	}
}

func DeactivateUser(baseURL string, user m.AuthUser, sonarUser m.SonarUser, verbose bool) {
	if sonarUser.Login == "" {
		log.Fatal("login is a required parameter for deactivating a user")
	}

	url := fmt.Sprintf("%s/api/users/deactivate", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("login", sonarUser.Login)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	if status == "200 OK" {
		log.Printf("User with login='%s' is deactivated", sonarUser.Login)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to delete a user", user.Username)
	}
}
