package backend

import (
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func ListProjects(baseURL string, user m.AuthUser, regex string, verbose bool) []m.Project {
	url := fmt.Sprintf("%s/api/projects/index", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)

	respBody, _ := u.HTTPRequest(req, verbose)

	var (
		projects      []m.Project
		projectsArray []m.Project
	)
	json.Unmarshal(respBody, &projects)

	if regex == "" {
		for _, project := range projects {
			projectsArray = append(projectsArray, project)
		}

		if len(projectsArray) == 0 {
			log.Println("There are no project(s) in sonarqube")
		}
	} else {
		cRegex, err := regexp.Compile(regex)
		u.Error(err, "There was a error compiling the regex")

		for _, project := range projects {
			if cRegex.MatchString(project.Key) {
				projectsArray = append(projectsArray, project)
			}
		}

		if len(projectsArray) == 0 {
			log.Fatal("There are no project(s) that match the entered regex '%s'", regex)
		}
	}
	return projectsArray
}

func ProjectExists(baseURL string, user m.AuthUser, projectKey string, verbose bool) bool {
	var projectExists bool
	projects := ListProjects(baseURL, user, "", verbose)
	for _, project := range projects {
		if project.Key == projectKey {
			projectExists = true
			break
		} else {
			projectExists = false
		}
	}
	return projectExists
}

func CreateProject(baseURL string, user m.AuthUser, project m.Project, verbose bool) {
	if project.Key == "" || project.Name == "" {
		log.Fatal("projectKey and projectName are required parameters for creating a new project")
	}

	url := fmt.Sprintf("%s/api/projects/create", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("key", project.Key)
	query.Add("name", project.Name)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	if status == "200 OK" {
		log.Printf("Project with key '%s' is created.", project.Key)
	} else if status == "400 Bad Request" {
		log.Printf("Could not create Project, key already exists: %s", project.Key)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to create a project", user.Username)
	}
}

func DeleteProject(baseURL string, user m.AuthUser, project m.Project, verbose bool) {
	if project.Key == "" {
		log.Fatal("projectKey is a required parameter for deleting a project")
	}

	url := fmt.Sprintf("%s/api/projects/delete", baseURL)

	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("key", project.Key)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(req, verbose)

	if status == "204 No Content" {
		log.Printf("Project with key '%s' is deleted.", project.Key)
	} else if status == "404 Not Found" {
		log.Printf("Project with key '%s' not found", project.Key)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to delete a project", user.Username)
	}
}
