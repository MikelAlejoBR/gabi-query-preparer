package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	whiteSpacesRegex = regexp.MustCompile(`\s{2,}|\n$`)
)

func main() {
	// Attempt grabbing the path passed by parameter.
	queryFilePath := flag.String("query-file-path", "", "the path to the query file you want to transform to a Gabi query file")
	flag.Parse()

	if queryFilePath == nil || *queryFilePath == "" {
		fmt.Println(`You did not specify the path to the query file. Please do so with the "query-file-path" parameter.`)
		os.Exit(1)
	}

	// Try to read the file.
	buffer, err := os.ReadFile(*queryFilePath)
	if err != nil {
		log.Fatalf("Unable to read the provided file: %s\n", err)
	}

	fileContents := string(buffer)

	// Remove the leading and trailing spaces.
	trimmedFileContents := strings.TrimSpace(fileContents)

	// Replace all the whitespaces, tabs and carriage returns that are longer than two characters. The goal is to put
	// the query in a single long line, and remove any formatted or pretty structure.
	cleanedFileContents := whiteSpacesRegex.ReplaceAllString(trimmedFileContents, " ")

	output := Output{
		Query: cleanedFileContents,
	}

	result, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		log.Fatalf("Unable to serialize the output as JSON: %s\n", err)
	}

	fmt.Println(string(result))
}
