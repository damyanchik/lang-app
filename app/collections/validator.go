package collections

func IsInArray(str string, array []string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

func IsKeyInMap(str string, collection map[string]string) bool {
	for key := range collection {
		if key == str {
			return true
		}
	}
	return false
}

func IsValueInMap(str string, collection map[string]string) bool {
	for _, value := range collection {
		if value == str {
			return true
		}
	}
	return false
}
