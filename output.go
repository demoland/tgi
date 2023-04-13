package main

import (
	"encoding/json"
	"fmt"
)

type outputKeys struct {
	Keys map[string]outputValues `json:"outputs,omitempty"`
}

type outputValues struct {
	Values interface{} `json:"value"`
	Type   interface{} `json:"type,omitempty"`
}

func getOutput(data []byte) error {

	//Unmarshal the JSON data into a outputKeys struct
	var output outputKeys
	err := json.Unmarshal(data, &output)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Print the Header:
	fmt.Printf("\033[1m\nOutputs:\n===================\n\033[0m")
	//	fmt.Println(output.Keys)
	for key := range output.Keys {
		v := output.Keys[key].Values
		if _, ok := v.(string); ok {
			fmt.Printf("%v:\n\t\033[32m%v\033[0m\n", key, v)
		} else if _, ok := v.([]interface{}); ok {
			//fmt.Println("\n", key, ":")
			fmt.Println(key, ":")
			for _, e := range v.([]interface{}) {
				fmt.Printf("\t\033[32m%v\033[0m\n", e)
			}
		}
	}
	return nil

}
