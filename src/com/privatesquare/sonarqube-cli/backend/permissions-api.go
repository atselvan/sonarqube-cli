package backend

import (
	"fmt"
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"log"
	"os"
)

func GrantDeveloperRole(baseURL string, user m.AuthUser, permission m.Permission, verbose bool){
	permission.Permission = "user"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "codeviewer"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "scan"
	grantProjectPermission(baseURL, user, permission, verbose)
}

func GrantIssueAdminRole(baseURL string, user m.AuthUser, permission m.Permission, verbose bool){
	permission.Permission = "user"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "codeviewer"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "issueadmin"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "scan"
	grantProjectPermission(baseURL, user, permission, verbose)
}

func GrantAdminRole(baseURL string, user m.AuthUser, permission m.Permission, verbose bool){
	permission.Permission = "admin"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "user"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "codeviewer"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "issueadmin"
	grantProjectPermission(baseURL, user, permission, verbose)
	permission.Permission = "scan"
	grantProjectPermission(baseURL, user, permission, verbose)
}

func grantProjectPermission(baseURL string, user m.AuthUser, permission m.Permission, verbose bool){

	if permission.ProjectKey == "" && permission.ViewKey == "" {
		log.Fatal("Providing a projectKey or a viewKey is required for granting permissions")
	} else if permission.Login == "" {
		log.Fatal("login is a required parameter for granting permission")
	}

	url := fmt.Sprintf("%s/api/permissions/add_user", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("login", permission.Login)
	query.Add("permission", permission.Permission)
	if permission.ProjectKey != "" {
		query.Add("projectKey", permission.ProjectKey)
	}else {
		query.Add("projectKey", permission.ViewKey)
	}
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	switch status {
	case "204 No Content":
		if permission.ProjectKey != "" {
			log.Printf("Permission '%s' is granted to user '%s' on project with key '%s'", permission.Permission, permission.Login, permission.ProjectKey)
		}else {
			log.Printf("Permission '%s' is granted to user '%s' on view with key '%s'", permission.Permission, permission.Login, permission.ViewKey)
		}
	case "404 Not Found":
		if permission.ProjectKey != "" {
			log.Printf("Project with key '%s' does not exist", permission.ProjectKey)
		}else {
			log.Printf("View with key '%s' does not exist", permission.ViewKey)
			os.Exit(1)
		}
	case "403 Forbidden":
		log.Printf("User '%s' is not Authorized to grant permissions", user.Username)
		os.Exit(1)
	case "400 Bad Request":
		log.Printf("User with login '%s' does not exist", permission.Login)
		os.Exit(1)
	default:
		panic(fmt.Sprintf("ERROR: call status=%v\n", status))
	}
}
