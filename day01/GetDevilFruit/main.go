package main

import (
	"fmt"
)

func main() {
	name, power, color, owner := GetDevilFruit("Mera mera no mi", "Fire Control", "Red", "Sabo")
	fmt.Print("Name: ", name+"\n")
	fmt.Print("Power: ", power+"\n")
	fmt.Print("Color: ", color+"\n")
	fmt.Print("Owner: ", owner+"\n")
}
