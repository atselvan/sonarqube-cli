package utils

import (
	m "com/privatesquare/sonarqube-cli/model"
	"fmt"
	"log"
)

// PrintStringArray prints a string Array
func PrintStringArray(array []string) {
	for _, element := range array {
		fmt.Println(element)
	}

}

func PrintProjectsArray(projectsArray []m.Project, regex string) {
	for _, array := range projectsArray {
		fmt.Println(array)
	}
	if regex != "" {
		log.Printf("Number of projects matching the regex '%s' is %d", regex, len(projectsArray))
	} else {
		log.Printf("Number of projects in sonarqube = %d", len(projectsArray))
	}
}

func PrintViewsArray(viewsArray []m.View) {
	for _, array := range viewsArray {
		fmt.Println(array)
	}
	log.Printf("Number of views in sonarqube = %d", len(viewsArray))
}