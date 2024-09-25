package main

import (
    "fmt"
)

func main() {
    crewMembers := []string{"Luffy", "Zoro", "Nami", "Sanji"}
    count := CountMembersWithLength(crewMembers, 4)
    fmt.Println("Number of members with a length of 4 characters", count)
}