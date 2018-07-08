package backend

import (
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckAuthentication(baseURL string, user m.AuthUser, verbose bool) {
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

func getUsers(baseURL string, user m.AuthUser, verbose bool) []m.UserDetails {
	url := fmt.Sprintf("%s/api/users/search", baseURL)
	var (
		usersApiResp m.UsersAPIResp
		users        []m.UserDetails
		isLast       = false
	)
	paging := m.Paging{PageIndex: 1, PageSize: 500}
	for isLast == false {
		req := u.CreateBaseRequest("GET", url, nil, user, verbose)
		query := req.URL.Query()
		query.Add("p", fmt.Sprintf("%d", paging.PageIndex))
		query.Add("ps", fmt.Sprintf("%d", paging.PageSize))
		req.URL.RawQuery = query.Encode()
		respBody, _ := u.HTTPRequest(req, verbose)
		json.Unmarshal(respBody, &usersApiResp)
		for _, user := range usersApiResp.Users {
			users = append(users, user)
		}
		if usersApiResp.Paging.Total-usersApiResp.Paging.PageIndex*usersApiResp.Paging.PageSize > 0 {
			paging.PageIndex++
		} else {
			isLast = true
		}
	}
	return users
}

func userExists(baseURL, userId string, user m.AuthUser, verbose bool) bool {
	if userId == "" {
		log.Fatal("userId is a required parameter for checking if a user exists")
	}
	var isExist bool
	users := getUsers(baseURL, user, verbose)
	for _, user := range users {
		if userId == user.Login {
			isExist = true
			break
		} else {
			isExist = false
		}
	}
	return isExist
}

func getUserDetails(baseURL, userId string, user m.AuthUser, verbose bool) m.UserDetails {
	if userId == "" {
		log.Fatal("userId is a required parameter for getting user details")
	}
	users := getUsers(baseURL, user, verbose)
	var userDetails m.UserDetails
	if userExists(baseURL, userId, user, verbose) {
		for _, user := range users {
			if userId == user.Login {
				userDetails = user
				break
			}
		}
	} else {
		log.Printf("User %s does not exist\n", userId)
		os.Exit(1)
	}
	return userDetails
}

func GetUsersList(baseURL string, user m.AuthUser, verbose bool) []string {
	var usersList []string
	users := getUsers(baseURL, user, verbose)
	for _, user := range users {
		usersList = append(usersList, user.Login)
	}
	return usersList
}

func PrintUserDetails(baseURL, userId string, user m.AuthUser, verbose bool) {
	if userId == "" {
		log.Fatal("userId is a required parameter for printing user details")
	}
	userDetails := getUserDetails(baseURL, userId, user, verbose)
	fmt.Printf("User ID				: %s\n", userDetails.Login)
	fmt.Printf("Name				: %s\n", userDetails.Name)
	fmt.Printf("Active				: %v\n", userDetails.Active)
	fmt.Printf("Email				: %s\n", userDetails.Email)
	fmt.Printf("Groups				: %s\n", userDetails.Groups)
	fmt.Printf("Tokens Count			: %d\n", userDetails.TokensCount)
	fmt.Printf("IsLocal				: %v\n", userDetails.Local)
	fmt.Printf("Identity Provider		: %s\n", userDetails.ExternalProvider)
}

func CheckAndCreateUser(baseURL, userPassword string, user m.AuthUser, userDetails m.UserDetails, verbose bool) {
	if !userExists(baseURL, userDetails.Login, user, verbose) {
		createUser(baseURL, userPassword, user, userDetails, verbose)
	} else {
		log.Printf("User %s already exists\n", userDetails.Login)
		os.Exit(1)
	}
}

func createUser(baseURL, userPassword string, user m.AuthUser, userDetails m.UserDetails, verbose bool) {
	if userDetails.Login == "" || userDetails.Name == "" || userDetails.Email == "" || userPassword == "" {
		log.Fatal("userId, name, email and userPassword are required parameters for creating a user")
	}
	url := fmt.Sprintf("%s/api/users/create", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)
	query := req.URL.Query()
	query.Add("login", strings.ToUpper(userDetails.Login))
	query.Add("name", strings.Title(userDetails.Name))
	query.Add("email", userDetails.Email)
	query.Add("password", userPassword)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	status = strings.Trim(status, " ")

	if status == "200" {
		log.Printf("User with login='%s' is created", userDetails.Login)
	} else if status == "401" {
		log.Printf("User '%s' is not authorized to create a new user", user.Username)
	} else {
		log.Printf("There was a problem creating the user. Use -verbose flag for more details", userDetails.Login)
	}
}

func CheckAndDeactivateUser(baseURL, userId string, user m.AuthUser, verbose bool) {
	if userExists(baseURL, userId, user, verbose) {
		deactivateUser(baseURL, userId, user, verbose)
	} else {
		log.Printf("User %s does not exist\n", userId)
		os.Exit(1)
	}
}

func deactivateUser(baseURL, userId string, user m.AuthUser, verbose bool) {
	if userId == "" {
		log.Fatal("userId is a required parameter for deactivating a user")
	}
	url := fmt.Sprintf("%s/api/users/deactivate", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)
	query := req.URL.Query()
	query.Add("login", userId)
	req.URL.RawQuery = query.Encode()
	_, status := u.HTTPRequest(req, verbose)

	status = strings.Trim(status, " ")

	if status == "200" {
		log.Printf("User %s is deactivated", userId)
	} else if status == "401" {
		log.Printf("User '%s' is not authorized to delete a user", user.Username)
	} else {
		log.Printf("There was a problem deactivating the user. Use -verbose flag for more details", userId)
	}
}
