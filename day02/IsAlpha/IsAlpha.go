package main

func IsAlpha(str string) bool {
	var i int
	var result bool = false
	for i = 0; i < len(str); i++ {
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') {
			result = true
		} else {
			return false
		}
	}
	return result
}
