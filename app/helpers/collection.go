package helpers

func IsInArray(str string, array []string) bool {
	for _, element := range array {
		if element == str {
			return true
		}
	}
	return false
}

func SearchKeyInMap(str string, collection map[string]string) bool {
	for key, _ := range collection {
		if key == str {
			return true
		}
	}
	return false
}

func SearchValueInMap(str string, collection map[string]string) bool {
	for _, value := range collection {
		if value == str {
			return true
		}
	}
	return false
}
