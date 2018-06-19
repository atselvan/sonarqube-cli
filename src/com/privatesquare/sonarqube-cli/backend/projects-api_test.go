package backend_test

import (
	"com/privatesquare/sonarqube-cli/backend"
	"com/privatesquare/sonarqube-cli/model"
	"testing"
)

const (
	baseURL  = "http://localhost:9000"
	username = "admin"
	password = "admin"
	verbose  = false
)
var(
	authUser = model.AuthUser{Username: username, Password: password}
	projectName = "test"
	projectKey = "com.test.test:test"
	project = model.Project{Name: projectName, Key: projectKey}
)

func TestProjectDoesNotExists(t *testing.T) {
	isExists := backend.ProjectExists(baseURL, authUser, projectKey, verbose)
	if isExists {
		t.Error("Expected false but got true.")
	}
}

func TestCreateProject200(t *testing.T) {
	status := backend.CreateProject(baseURL, authUser, project, verbose)
	if status != "200 " {
		t.Error("Create project test for status 200 failed")
	}
}

func TestCreateProject400(t *testing.T) {
	status := backend.CreateProject(baseURL, authUser, project, verbose)
	if status != "400 " {
		t.Error("Create project test for status 400 failed")
	}
}

func TestProjectExists(t *testing.T) {
	isExists := backend.ProjectExists(baseURL, authUser, projectKey, verbose)
	if !isExists {
		t.Error("Expected true but got false.")
	}
}

func TestDeleteProject204(t *testing.T) {
	status := backend.DeleteProject(baseURL, authUser, project, verbose)
	if status != "204 " {
		t.Error("Delete project test for status 204 failed")
	}
}

func TestDeleteProject404(t *testing.T) {
	status := backend.DeleteProject(baseURL, authUser, project, verbose)
	if status != "404 " {
		t.Error("Delete project test for status 404 failed")
	}
}