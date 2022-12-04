package main

import (
	"encoding/json"
	"fmt"
	githubhttpfetcher "github.com/bachrc/repo-stats/internal/github-http-fetcher"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Repositories []Repository

type Repository struct {
	Full_Name     string
	Languages_url string
}

func main() {
	fetch_license_data()
}

func fetch_license_data() {
	file, err := ioutil.ReadFile("test/data/github-data/repositories.json")
	if err != nil {
		fmt.Println(err)
	}
	var repositories Repositories

	_ = json.Unmarshal(file, &repositories)

	fmt.Println(repositories)

	for i := range repositories {
		writeLicenseData(repositories[i])
	}

	//stubPath := "../../test/data/github-data" + requestedGithubPath + ".json"
}

func writeLicenseData(repository Repository) {
	fetcher := githubhttpfetcher.New(os.Getenv("GITHUBACCESSTOKEN"))

	licenseUrl := "https://api.github.com/repos/" + repository.Full_Name + "/license"

	request, _ := http.NewRequest(http.MethodGet, licenseUrl, nil)

	urlLicense, _ := url.Parse(licenseUrl)

	response, _ := fetcher.Client.Do(request)

	content, _ := ioutil.ReadAll(response.Body)
	_ = ioutil.WriteFile("test/data/github-data"+urlLicense.Path+".json", content, os.ModePerm)
}
