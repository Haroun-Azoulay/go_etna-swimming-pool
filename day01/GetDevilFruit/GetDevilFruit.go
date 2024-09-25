package main

import (
	"strings"
)

func GetDevilFruit(name string, power string, color string, owner string) (string, string, string, string) {
	name = strings.ToUpper(name)
	return name, power, color, owner
}
