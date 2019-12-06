package dockerhub

import (
	"fmt"
	"github.com/Matt-Gleich/logErr"
	"io/ioutil"
	"net/http"
)

// ListOfContainers ... List all all a user's containers
func ListOfContainers(uName string) []string {
	// Making request
	resp, err := http.Get("https://hub.docker.com/v2/repositories/" + uName + "/")
	logErr.Log(err, "Failed to get for the user: " + uName)

	fmt.Println(resp)

	// Reading request response
	body, err := ioutil.ReadAll(resp.Body)
	logErr.Log(err, "Failed to get unpack http request to docker hub")


	type results struct {}
	results := body["results"]
	if results == [] {
		logErr.Log(true, "No repos found for the given user")
	}
	repos := []string
	for i := 0; i != len(results); i++ {
		repos = append(repos, uName + "/" + results[i]["name"])
	}
	return repos
}



