package helper

func IsInArray(arr []interface{}, key interface{}) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
