# Sonarqube CLI

```sh
Usage of ./sonarqube-cli:
  -addLocalSubview
    	Add a subview as a local reference. Required viewKey: refViewKey
  -createProject
    	Creates a project in sonarqube. Required paramters: projectKey, projectName
  -createUser
    	Create a user. Required paramters: login, name, email
  -createView
    	Creates a view in sonarqube. Required paramters: viewKey, viewName
  -deactivateUser
    	Deactivate a user. Required paramters: login
  -deleteProject
    	Deletes a project from sonarqube. Required paramter: projectKey
  -deleteView
    	Deletes a view from sonarqube. Required paramter: viewKey
  -email string
    	Email ID of the user (default "something@something.com")
  -listProjects
    	Lists the projects in sonarqube. Optional paramter: regex
  -listViews
    	Lists the views in sonarqube
  -login string
    	Login ID of the user (default "something")
  -name string
    	Name of the user (default "something")
  -password string
    	SonarQube username's password (Required) (default "admin")
  -projectKey string
    	Project Key
  -projectName string
    	ProjectName
  -refViewKey string
    	Local reference view key
  -regex string
    	Regular expression to filter projects
  -sonarUrl string
    	SonarQube URL (Required) (default "http://localhost:9000")
  -username string
    	SonarQube username (Required) (default "admin")
  -verbose
    	Set the flag if you want verbose output
  -viewDesc string
    	View description
  -viewKey string
    	View key
  -viewName string
    	View name
```
