package main

import (
	"fmt"
)

func main() {
	crew1 := []string{"Luffy", "Zoro", "Sanji"}
	crew2 := []string{"Luffy", "Zoro", "Sanji"}
	isEqual := CompareCrews(crew1, crew2)
	fmt.Println("Les équipages sont-ils identiques ?", isEqual)
}
