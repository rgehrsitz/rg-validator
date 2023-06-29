package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rgehrsitz/rg-validator/pkg/validator"
	"github.com/rgehrsitz/rulegopher/pkg/rules"
)

func main() {
	// Parse command-line arguments
	rulesFile := flag.String("rules", "", "Path to the JSON file containing the rules")
	flag.Parse()

	if *rulesFile == "" {
		fmt.Println("Please provide a rules file with the -rules option")
		os.Exit(1)
	}

	// Read the rules file
	data, err := ioutil.ReadFile(*rulesFile)
	if err != nil {
		fmt.Printf("Failed to read the rules file: %v\n", err)
		os.Exit(1)
	}

	// Decode the rules file
	var rule rules.Rule
	err = json.Unmarshal(data, &rule)
	if err != nil {
		fmt.Printf("Failed to decode the rules file: %v\n", err)
		os.Exit(1)
	}

	// Validate the rule
	err = validator.ValidateRule(rule)
	if err != nil {
		fmt.Printf("Invalid rule: %v\n", err)
		os.Exit(1)
	}

	// If the rule is valid, print a success message
	fmt.Println("The rule is valid!")
}
