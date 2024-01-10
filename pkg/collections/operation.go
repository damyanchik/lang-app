package collections

import "math/rand"

func MapKeysToSlice(collection map[string]string) []string {
	keys := make([]string, 0, len(collection))
	for key := range collection {
		keys = append(keys, key)
	}

	return keys
}

func RandomOrderInSlice(collection []string) []string {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}
