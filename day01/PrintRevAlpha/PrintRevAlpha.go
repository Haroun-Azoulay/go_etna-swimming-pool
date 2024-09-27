package main

import "fmt"

func PrintRevAlpha() {
	for i := 122; i >= 97; i-- {
		fmt.Printf("%c", i)
	}
}
