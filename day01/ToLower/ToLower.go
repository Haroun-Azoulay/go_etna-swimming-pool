package main

import "fmt"

func ToLower(str string) string {
	var lowCharacteres string = ""
	for _, chr := range str {
		if chr >=65 && chr <= 90 {
			chr +=32
		}
		lowCharacteres += fmt.Sprintf("%c", chr)
	}
	return lowCharacteres
}
