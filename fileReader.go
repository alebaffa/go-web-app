package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSONReturnNewsletter(jsonFile string) Newsletter {

	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file not found!")
	}

	var newsletter Newsletter
	err = json.Unmarshal(file, &newsletter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in marshaling the file", err)
	}

	return newsletter
}
