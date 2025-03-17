// Package Declartion
package devops

// Imports
import (
	"fmt"
	urls "net/url"
)

// Get Base Endpoint
func GetBaseEndpoint(org, project, team string) string {

	// Get URL Containing Organization
	endpoint := fmt.Sprintf("https://dev.azure.com/%s", org)

	// Add Project
	if project != "" {
		endpoint += fmt.Sprintf("/%s", project)
	}

	// Add Team
	if team != "" {
		endpoint += fmt.Sprintf("/%s", team)
	}

	// Return Result
	return endpoint

}

// Add Query Parameters
func AddQueryParams(url string, args map[string]string) string {

	// Add Base Parameters
	urlNew := fmt.Sprintf("%s?api-version=7.1", url)

	// Add Custom Arguments
	for key, value := range args {
		urlNew += fmt.Sprintf("&%s=%s", urls.QueryEscape(key), urls.QueryEscape(value))
	}

	// Return Result
	return urlNew
}

// Get WIQL Endpoint
func GetWiqlEndpoint(org, project, team, id string) string {

	// Get Base URL
	urlBase := GetBaseEndpoint(org, project, team)

	// Add WIQL Specifics
	url := fmt.Sprintf("%s/_apis/wit/wiql", urlBase)
	if id != "" {
		url += fmt.Sprintf("/%s", id)
	}

	// Add Query Parameters
	url = AddQueryParams(url, nil)

	// Return URL
	return url
}
