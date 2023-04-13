package main

import (
	"encoding/json"
	"fmt"
)

type resourceKeys struct {
	Keys []resourceValues `json:"resources,omitempty"`
}

type resourceValues struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Mode      string      `json:"mode"`
	Instances interface{} `json:"instances,omitempty"`
}

func getResources(data []byte) error {
	//Unmarshal the JSON data into a outputKeys struct
	var resources resourceKeys

	err := json.Unmarshal(data, &resources)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	//Print the Header:
	fmt.Printf("\033[1m\nResources:\n===================\n\033[0m")
	//Set Managed Count to Zero
	mc := 0
	for r := range resources.Keys {
		// Name: aws_instance.example, aws_security_group.example, etc.
		n := resources.Keys[r].Name
		// Type: aws_instance, aws_security_group, etc.
		t := resources.Keys[r].Type
		// Mode: Managed or Data
		m := resources.Keys[r].Mode
		// Instance Count
		ic := len(resources.Keys[r].Instances.([]interface{}))
		// If Instance Count is greater than 1, Iterate over the Instances and Print the Resource
		if ic > 1 {
			for inst := range resources.Keys[r].Instances.([]interface{}) {
				//Increment Managed Count if Mode = Managed (vs. Data)
				if m == "managed" {
					mc++
					fmt.Printf("\033[32m%s:\033[0m %s.%s.[%d] \n", m, t, n, inst)
				} else {
					fmt.Printf("033[38;5;208m%s:\033[0m%s.%s.[%d] \n", m, t, n, inst)
				}
			}
			// If Instance Count is 1, Print the Resource
		} else {
			//Increment Managed Count if Mode = Managed (vs. Data)
			if m == "managed" {
				mc++
				fmt.Printf("\033[32m%s:\033[0m %s.%s \n", m, t, n)
			} else {
				fmt.Printf("\033[38;5;208m%s:\033[0m %s.%s \n", m, t, n)
			}
		}
	}

	fmt.Printf("\nResources Under Management(RUM): %d\n", mc)
	return nil
}
