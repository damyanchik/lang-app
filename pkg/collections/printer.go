package collections

import "fmt"

func PrintMap(collection map[string]string) {
	i := 1
	for key, value := range collection {
		fmt.Printf("%v: %v - %v\n", i, key, value)
		i++
	}
}
