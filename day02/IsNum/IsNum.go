package main

func IsNum(str string) bool {
	for _, char := range str {
		if !(char >= '1' && char <= '9') {	
			return false
		}
	}
	return true
}
