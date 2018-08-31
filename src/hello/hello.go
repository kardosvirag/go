package main

import "fmt"

func main() {
	name := "Johny"
	age := 22

	rect1 := Rectangle{width: 30, height: 10}

	fmt.Println(returnString(name))
	fmt.Println(returnInt(age))
	fmt.Println(rect1)
	fmt.Println(rect1.width)
	fmt.Println(rect1.height)
	fmt.Println()

	firstListNumbers := []string{"13", "78", "30", "4", "5"}
	//secondListNumbers := []string{"5", "13", "4", "30", "78", "8", "0"}
	secondListNumbers := []string{"5", "0", "3", "78", "8"}
	fmt.Println(arraysEqual(firstListNumbers, secondListNumbers))

	fmt.Println(secondListNumbers)
	for i := 0; i < len(firstListNumbers); i++ {
		fmt.Println(firstListNumbers[i])
	}

	for i := 0; i < len(firstListNumbers); i++ {
		fmt.Println("first numbers:", firstListNumbers[i])
		for j := 0; j < len(firstListNumbers); j++ {
			fmt.Println(secondListNumbers[j])
			if firstListNumbers[i] == secondListNumbers[j] {
				fmt.Println(firstListNumbers[i], " = ", secondListNumbers[j])
			}
		}
	}
}

func arraysEqual(a, b []string) bool {
	var result bool
	number := 0
	for i := 0; i < len(a); i++ {
		fmt.Println("first numbers:", a[i])
		for j := 0; j < len(a); j++ {
			fmt.Println(b[j])
			if a[i] == b[j] {
				fmt.Println(a[i], " = ", b[j])
				number = number + 1
			}
		}
	}

	if number == len(a) {
		result = true
	} else {
		result = false
	}
	//a must be the "should tiles"
	return result
}

func returnString(stringName string) string {
	return stringName
}

func returnInt(number int) int {
	return number
}

type Rectangle struct {
	width  int
	height int
}

// // TODO verify IDs

// expectedUIDs := []string{"256", "92379", "30037"}

// for i := 0; i < len(searchResponse.Response.Docs); i++ {
// 	fmt.Printf("%v\n", searchResponse.Response.Docs[i].UID)
// }
// var result bool
// number := 0

// //a variable must be the "should tiles"
// for i := 0; i < len(expectedUIDs); i++ {
// 	fmt.Println("should be:", expectedUIDs[i])
// 	for j := 0; j < len(expectedUIDs); j++ {
// 		fmt.Println(searchResponse.Response.Docs[j].UID)
// 		if expectedUIDs[i] == searchResponse.Response.Docs[j].UID[i] {
// 			fmt.Println(expectedUIDs[i], " = ", searchResponse.Response.Docs[j].UID, " true")
// 			number = number + 1
// 		}
// 	}
// }
// if number == len(expectedUIDs) {
// 	result = true
// } else {
// 	result = false
// }
// fmt.Println(result)
