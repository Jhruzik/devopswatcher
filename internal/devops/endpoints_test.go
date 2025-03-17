// Package Declaration
package devops

// Imports
import (
	"fmt"
	"testing"
)

// Test Base Endpoint Retrieval
func TestGetBaseEndpoint(t *testing.T) {

	// Prepare Arguments
	org := "testorg"
	project := "testproject"
	team := "testteam"

	// Check Base
	endpointSimple := GetBaseEndpoint(org, "", "")
	expectationSimple := fmt.Sprintf("https://dev.azure.com/%s", org)

	if endpointSimple != expectationSimple {
		t.Errorf("Expectation: \"%s\"; Result: \"%s\"", expectationSimple,
			endpointSimple)
	}

	// Check Project
	endpointProject := GetBaseEndpoint(org, project, "")
	expectationProject := fmt.Sprintf("https://dev.azure.com/%s/%s", org, project)

	if endpointProject != expectationProject {
		t.Errorf("Expectation: \"%s\"; Result: \"%s\"", expectationProject, endpointProject)
	}

	// Check Team
	endpointTeam := GetBaseEndpoint(org, project, team)
	expectationTeam := fmt.Sprintf("https://dev.azure.com/%s/%s/%s", org, project, team)

	if endpointProject != expectationProject {
		t.Errorf("Expectation: \"%s\"; Result: \"%s\"", expectationTeam, endpointTeam)
	}
}

// Test Query Parameters Retrieval
func TestAddQueryParams(t *testing.T) {

	// Build Arguments
	args := map[string]string{
		"test": "A",
		"bird": "parrot",
	}

	// Build URL
	url := GetBaseEndpoint("test", "", "")

	// Add Parameters to URL
	url = AddQueryParams(url, args)

	// Run Check
	expectation := "https://dev.azure.com/test?api-version=7.1&test=A&bird=parrot"
	if url != expectation {
		t.Errorf("Expectation: \"%s\"; Result: \"%s\"", expectation, url)
	}

}

// Test WIQL Endpoint Retrieval
func TestGetWiqlEndpoint(t *testing.T) {

	// Build Endpoint
	org := "org"
	project := "project"
	team := "team"
	id := "id"
	url := GetWiqlEndpoint(org, project, team, id)

	// Run Checks
	expectation := fmt.Sprintf("https://dev.azure.com/%s/%s/%s/_apis/wit/wiql/%s?api-version=7.1", org, project, team, id)
	if url != expectation {
		t.Errorf("Expectation: \"%s\"; Result: \"%s\"", expectation, url)
	}

}
