package backend

import (
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"encoding/json"
	"fmt"
	"log"
)

func ListViews(baseURL string, user m.AuthUser, verbose bool) []m.View {
	url := fmt.Sprintf("%s/api/views/list", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)

	respBody, _ := u.HTTPRequest(req, verbose)

	var (
		views      m.Views
		viewsArray []m.View
	)

	json.Unmarshal(respBody, &views)

	for _, view := range views.Views {
		viewsArray = append(viewsArray, view)
	}

	if len(viewsArray) == 0 {
		log.Fatal("There are no view(s) in sonarqube")
	}

	return viewsArray
}

func ViewExists(baseURL string, user m.AuthUser, viewKey string, verbose bool) bool {

	var viewExists bool
	views := ListViews(baseURL, user, verbose)

	for _, view := range views {
		if view.Key == viewKey {
			viewExists = true
			break
		} else {
			viewExists = false
		}
	}
	return viewExists
}

func CreateView(baseURL string, user m.AuthUser, view m.View, verbose bool) {
	if view.Key == "" || view.Name == "" {
		log.Fatal("viewKey and viewName are required parameters for creating a new view")
	}

	if !ViewExists(baseURL, user, view.Key, verbose) {

		status := createView(baseURL, user, view, verbose)

		switch status {
		case "200 OK":
			log.Printf("View '%s' is created", view.Name)
		case "403 Forbidden":
			log.Printf("User '%s' is not Authorized to create a view", user.Username)
		default:
			panic(fmt.Sprintf("ERROR: call status=%v\n", status))
		}
	} else {
		log.Printf("A view with key %s already exists", view.Key)
	}
}

func createView(baseURL string, user m.AuthUser, view m.View, verbose bool) string {

	url := fmt.Sprintf("%s/api/views/create", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("key", view.Key)
	query.Add("name", view.Name)
	if view.Description != "" {
		query.Add("description", view.Description)
	}
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	return status
}

func DeleteView(baseURL string, user m.AuthUser, view m.View, verbose bool) {
	if view.Key == "" {
		log.Fatal("viewKey is a required parameter for creating a new view")
	}

	if ViewExists(baseURL, user, view.Key, verbose) {

		status := deleteView(baseURL, user, view, verbose)

		switch status {
		case "204 No Content":
			log.Printf("View with key '%s' is deleted", view.Key)
		case "403 Forbidden":
			log.Printf("User '%s' is not Authorized to delete a view", user.Username)
		default:
			panic(fmt.Sprintf("ERROR: call status=%v\n", status))
		}
	} else {
		log.Printf("A view with key %s does not exists", view.Key)
	}
}

func deleteView(baseURL string, user m.AuthUser, view m.View, verbose bool) string {

	url := fmt.Sprintf("%s/api/views/delete", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("key", view.Key)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	return status
}

func AddLocalSubview(baseURL string, user m.AuthUser, view m.View, verbose bool) {
	if view.Key == "" || view.RefKey == "" {
		log.Fatal("viewKey and refViewKey are required parameters for adding a view as a local reference")
	}

	url := fmt.Sprintf("%s/api/views/add_local_view", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("key", view.Key)
	query.Add("ref_key", view.RefKey)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	if status == "200 OK" {
		log.Printf("View with key '%s' is added as a local reference to view with key '%s'", view.RefKey, view.Key)
	} else if status == "400 Bad Request" {
		log.Printf("View with key '%s' is already referenced to view with key '%s'", view.RefKey, view.Key)
	} else if status == "403 Forbidden" {
		log.Printf("User '%s' is not Authorized to perform this operation", user.Username)
	}
}
