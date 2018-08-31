package main

import (
	"fmt"
)

func main() {

	expectedUIDs := map[string]string{
		"92379": "Kauningerhof",
		"30037": "Die kleine Gärtnerei",
		"256":   "Blatt und Blüte",
		"17":    "Fictive Tile",
	}

	blumen := input{UIDs: expectedUIDs}

	returnMap(blumen.UIDs)

}

func returnMap(x map[string]string) {
	for key, value := range x {
		fmt.Printf("%v - %v\n", key, value)
	}
}

type input struct {
	UIDs map[string]string
}

// func searchResult() {
// 	// key: UIDs value:sem_contact_name
// 	expectedUIDs := map[string]string{
// 		"92379": "Kauningerhof",
// 		"30037": "Die kleine Gärtnerei",
// 		"256":   "Blatt und Blüte",
// 		"175":   "Fictive Tile",
// 		"195":   "Second Fictive Tile",
// 	}

// 	actualUIDs := map[string]string{
// 		"92379": "Kauningerhof",
// 		"256":   "Blatt und Blüte",
// 		"30037": "Die kleine Gärtnerei",
// 	}

// 	fmt.Println("")
// 	fmt.Println(expectedUIDs)
// 	fmt.Println("blumen search result...\n")

// 	for expectedUID := range expectedUIDs {
// 		if _, expectedUIDExists := actualUIDs[expectedUID]; expectedUIDExists {
// 			fmt.Printf("Expected UID %v does exist in actual IDs.\n", expectedUID)
// 		} else {
// 			fmt.Printf("Expected UID %v does not exist in actual IDs.\n", expectedUID)
// 		}
// 	}
// }
