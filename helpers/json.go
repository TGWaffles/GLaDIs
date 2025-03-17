package helpers

import (
	"encoding/json"
	"fmt"
)

func PrintStructAsJSON(s any) {
	st, err := json.MarshalIndent(s, "", "  ")

	if err != nil {
		fmt.Printf("failed to print struct")
	}

	fmt.Println(string(st))
}
