package main

import (
    "fmt"
)

func main() {
    crew1 := []string{"Luffy", "Zoro"}
    crew2 := []string{"Sanji", "Nami"}
    mergedCrew := MergeCrews(crew1, crew2)
    fmt.Println("Équipage fusionné:", mergedCrew)
}