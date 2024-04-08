package helpers

import (
	"encoding/json"
	"fmt"
)

func ObjectToString(obj interface{}) string {
	jsonString, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return string(jsonString)
}