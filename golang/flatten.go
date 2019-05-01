package main

import "fmt"
import "log"

func main() {
	var one = 1
	var two = List(2, 3)
	var three = List(two, 3, 4)
	var four = List(three, 5, one, two, three)
	var list = List(four,
		7,
		8,
		List(),
	)
	fmt.Println(list)
	fmt.Println(Flatten(list))
}

// Flatten function receives an array of arbitrarily nested arrays of integers
// then returns a flat array of integers
func Flatten(list []interface{}) (result []int) {
	for _, element := range list {
		switch data := element.(type) {
		case int:
			log.Printf("Data int: %v", data)
			result = append(result, data)
		case []interface{}:
			log.Printf("Data interface: %v", data)
			result = append(result, Flatten(data)...)
		}
	}
	return
}

// List function creates a list of elements in an easy way
func List(element ...interface{}) []interface{} {
	return element
}
