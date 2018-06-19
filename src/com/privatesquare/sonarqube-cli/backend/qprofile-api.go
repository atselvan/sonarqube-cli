package backend

import (
	m "com/privatesquare/sonarqube-cli/model"
	u "com/privatesquare/sonarqube-cli/utils"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"io"
	"mime/multipart"
	"bytes"
	"net/http"
)

func BackupQualityProfile (baseURL string, user m.AuthUser, profile m.QualityProfile, verbose bool) {
	if profile.Name == "" || profile.Language == "" {
		log.Fatal("profileName and profileLang are required parameters for managing quality profiles")
	}
	url := fmt.Sprintf("%s/api/qualityprofiles/backup", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)
	filename := fmt.Sprintf("%s-quality-profile.xml", profile.Language)
	filePath := fmt.Sprintf("%s/%s", profile.FilePath, filename)

	_, err := os.Stat(profile.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s path does not exist", profile.FilePath)
			os.Exit(1)
		}
	}
	query := req.URL.Query()
	query.Add("profileName", profile.Name)
	query.Add("language", profile.Language)
	req.URL.RawQuery = query.Encode()

	respBody, status := u.HTTPRequest(req, verbose)

	formattedXML, err := u.FormatXML(respBody)
	u.Error(err, "There was a error formatting the XML file")

	writeErr := ioutil.WriteFile(filePath, []byte(formattedXML), 0666)
	u.Error(writeErr, "There was a error writing to the file")

	if status == "200 OK" {
		log.Printf("Backup done of profile '%s'", profile.Name)
	} else if status == "400 Bad Request" {
		log.Printf("Quality profile '%s' does not exist", profile.Name)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to manage quality profiles", user.Username)
	} else if status == "404 Not Found" {
		log.Printf("Unable to find a profile for language '%s' with name '%s'", profile.Language, profile.Name)
	}
}

func RestoreQualityProfile (baseURL string, user m.AuthUser, profile m.QualityProfile, verbose bool) {
	if profile.Name == "" || profile.Language == "" {
		log.Fatal("profileName and profileLang are required parameters for managing quality profiles")
	}
	url := fmt.Sprintf("%s/api/qualityprofiles/restore", baseURL)
	filename := fmt.Sprintf("%s-quality-profile.xml", profile.Language)
	filePath := fmt.Sprintf("%s/%s", profile.FilePath, filename)
	// Prepare a form that you will submit to that URL
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s file does not exist", filePath)
			os.Exit(1)
		}
	}
	fileWriter, err := writer.CreateFormFile("backup", filePath)
	u.Error(err, "There was a problem creating the form file")
	_, err = io.Copy(fileWriter, file)
	u.Error(err, "There was a problem copying the file")
	writer.Close()

	req, err := http.NewRequest("POST", url, &buffer)
	u.Error(err, "There was a error creating the request")
	req.SetBasicAuth(user.Username, user.Password)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if verbose {
		log.Println("Request Header: ", req)
	}
	_, status := u.HTTPRequest(req, verbose)

	if status == "200 OK" {
		log.Printf("Profile '%s' is restored", profile.Name)
	} else if status == "400 Bad Request" {
		log.Printf("There was a problem while restoring the profile %s. Set verbose for more information", profile.Name)
	} else if status == "401 Unauthorized" {
		log.Printf("User '%s' is not Authorized to manage quality profiles", user.Username)
	}
}