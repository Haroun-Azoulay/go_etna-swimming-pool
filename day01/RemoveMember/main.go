package main

import (
	"fmt"
)

func main() {
	crewMembers := []string{"Luffy", "Zoro", "Nami", "Sanji"}
	crewMembers = RemoveMember(crewMembers, 2)
	fmt.Println(crewMembers)
}
