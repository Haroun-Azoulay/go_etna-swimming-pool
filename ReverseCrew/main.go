package main

import (
	"fmt"
)

func main() {
	crewMembers := [10]string{"Luffy", "Zoro", "Nami", "Usopp", "Sanji", "Chopper", "Robin", "Franky", "Brook", "Jinbe"}
	reversedCrew := ReverseCrew(crewMembers)
	fmt.Println(reversedCrew)
}
