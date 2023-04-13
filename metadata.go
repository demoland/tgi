package main

import (
	"encoding/json"
	"fmt"
)

type statefileMeta struct {
	CheckResults     string `json:"check_results"`
	Lineage          string `json:"lineage"`
	Version          int16  `json:"version"`
	Serial           int16  `json:"serial"`
	TerraformVersion string `json:"terraform_version"`
}

func getMetadata(data []byte) error {
	//Unmarshal the JSON data into a statefileMeta struct
	var meta statefileMeta
	err := json.Unmarshal(data, &meta)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	// Print the metadata
	// fmt.Printf("")
	fmt.Printf("\033[1m\nStatefile Metadata:\n===================\n\033[0m")
	if meta.CheckResults == "" {
		meta.CheckResults = "null"
	}
	fmt.Printf("Check Results: %s\n\033[0m", meta.CheckResults)
	fmt.Printf("Lineage: \033[32m%s\n\033[0m", meta.Lineage)
	fmt.Printf("Version: \033[32m%d\n\033[0m", meta.Version)
	fmt.Printf("Serial: \033[32m%d\n\033[0m", meta.Serial)
	fmt.Printf("Terraform Version: \033[32m%s\n\033[0m", meta.TerraformVersion)
	return nil

}
