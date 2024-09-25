package main

import (
    "fmt"
)

func CountMembersWithLength(crew []string, length int) int {
	count := 0
	for _, a := range crew {
		if len(a) == length {
			count ++
		}
    }

	return count
}
