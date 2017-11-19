package main

import (
	"flag"
	b "com/privatesquare/sonarqube-cli/backend"
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"log"
	"fmt"
)

func main() {

	//options
	createUser := flag.Bool("createUser", false, "Create a user. Required paramters: login, name, email")
	deactivateUser := flag.Bool("deactivateUser", false, "Deactivate a user. Required paramters: login")
	listProjects := flag.Bool("listProjects", false, "Lists the projects in sonarqube. Optional paramter: regex")
	createProject := flag.Bool("createProject", false, "Creates a project in sonarqube. Required paramters: projectKey, projectName")
	deleteProject := flag.Bool("deleteProject", false, "Deletes a project from sonarqube. Required paramter: projectKey")
	listViews := flag.Bool("listViews", false, "Lists the views in sonarqube")
	createView := flag.Bool("createView", false, "Creates a view in sonarqube. Required paramters: viewKey, viewName")
	deleteView := flag.Bool("deleteView", false, "Deletes a view from sonarqube. Required paramter: viewKey")
	addLocalSubview := flag.Bool("addLocalSubview", false, "Add a subview as a local reference. Required viewKey: refViewKey")

	//paramters
	sonarURL := flag.String("sonarUrl", "http://localhost:9000", "SonarQube URL (Required)")
	username := flag.String("username", "admin", "SonarQube username (Required)")
	password := flag.String("password", "admin", "SonarQube username's password (Required)")

	login := flag.String("login", "something", "Login ID of the user")
	name := flag.String("name", "something", "Name of the user")
	email := flag.String("email", "something@something.com", "Email ID of the user")

	regex := flag.String("regex", "", "Regular expression to filter projects")
	projectKey := flag.String("projectKey", "", "Project Key")
	projectName := flag.String("projectName", "", "ProjectName")

	viewKey := flag.String("viewKey", "", "View key")
	viewName := flag.String("viewName", "", "View name")
	viewDescription := flag.String("viewDesc", "", "View description")
	refViewKey := flag.String("refViewKey", "", "Local reference view key")

	verbose := flag.Bool("verbose", false, "Set the flag if you want verbose output")
	flag.Parse()

	user := m.AuthUser{Username: *username, Password: *password}
	sonarUser := m.SonarUser{Login: *login, Name: *name, Email: *email, Password: "defaultPass"}
	project := m.Project{Key: *projectKey, Name: *projectName}
	view := m.View{Key: *viewKey, Name: *viewName, Description: *viewDescription, RefKey: *refViewKey}

	b.ChechAuthentication(*sonarURL, user, *verbose)

	fmt.Println(b.ViewExists(*sonarURL, user, *viewKey, *verbose))

	if *createUser {
		b.CreateUser(*sonarURL, user, sonarUser, *verbose)
	} else if *deactivateUser {
		b.DeactivateUser(*sonarURL, user, sonarUser, *verbose)
	} else if *listProjects {
		projects := b.ListProjects(*sonarURL, user, *regex, *verbose)
		u.PrintProjectsArray(projects, *regex)
	} else if *createProject {
		b.CreateProject(*sonarURL, user, project, *verbose)
	} else if *deleteProject {
		b.DeleteProject(*sonarURL, user, project, *verbose)
	} else if *listViews {
		views := b.ListViews(*sonarURL, user, *verbose)
		u.PrintViewsArray(views)
	} else if *createView {
		b.CreateView(*sonarURL, user, view, *verbose)
	} else if *deleteView {
		b.DeleteView(*sonarURL, user, view, *verbose)
	} else if *addLocalSubview {
		b.AddLocalSubview(*sonarURL, user, view, *verbose)
	} else {
		flag.Usage()
		log.Fatal("Select a valid action flag")
	}


}