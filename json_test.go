package main

import (
	"testing"
	"encoding/json"
	"os"
	"fmt"
)

func TestReadAndUnmarshal(t *testing.T) {
	var expectedNewsletter Newsletter
	jsonTestfile := "./db/newsletter.json"

	jsonFile, err := os.Open(jsonTestfile)
	if err != nil {
		fmt.Errorf("opening config file", err.Error())
	}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&expectedNewsletter); err != nil {
		fmt.Errorf("parsing config file", err.Error())
	}

	actualNewsletter := readJsonReturnNewsletter(jsonTestfile)
	if !areNewsletterDeepEqual(actualNewsletter, expectedNewsletter){
		t.Error("the newsletters are not equal!")
	}

}

func areNewsletterDeepEqual(newsletterActual, newsletterExpected Newsletter) bool {
	for index, issue := range newsletterActual.Issues {
		issue2 := newsletterExpected.Issues[index]
		if issue.Id == issue2.Id && issue.Status == issue2.Status && areLinksEqual(issue, issue2){
			return true
		}
	}
	return false
}

func areLinksEqual(issue1, issue2 Issue) bool {
	for index, link1 := range issue1.Links {
		link2 := issue2.Links[index]
		if link1.Contributor == link2.Contributor && link1.Link == link2.Link && link1.Title == link2.Title {
			return true
		}
	}
	return false
}
