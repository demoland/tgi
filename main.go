package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	//Define a -file flag to specify the JSON file name
	fileFlag := flag.String("file", "", "JSON file to parse")
	metaFlag := flag.Bool("meta", false, "Print statefile metadata")
	outputFlag := flag.Bool("outputs", false, "Print outputs")
	resourceFlag := flag.Bool("resources", false, "Print number of resources under management")
	flag.Parse()

	//Check if the -file flag was provided
	if *fileFlag == "" {
		fmt.Println("Please provide a JSON file to parse")
		return
	}

	// Read the contents of the file
	data, err := ioutil.ReadFile(*fileFlag)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//TODO: Turn this into a switch statement instead of a bunch of if/else statements:
	if *metaFlag {
		err := getMetadata(data)
		if err != nil {
			fmt.Println("Error getting metadata:", err)
			return
		}
	} else if *outputFlag {
		err := getOutput(data)
		if err != nil {
			fmt.Println("Error getting output:", err)
			return
		}
	} else if *resourceFlag {
		err := getResources(data)
		if err != nil {
			fmt.Println("Error getting resources:", err)
			return
		}
		return
	} else {
		err := getMetadata(data)
		if err != nil {
			fmt.Println("Error getting metadata:", err)
			return
		}
	}
}
