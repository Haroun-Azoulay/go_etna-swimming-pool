package main

func ToUpper(str string) string {
	var result string
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result += string(char)
		} else {
			result += string(char - ('a' - 'A'))
		}
	}
	return result
}
