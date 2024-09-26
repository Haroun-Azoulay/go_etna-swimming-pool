package main

func CompareCrews(crew1 []string, crew2 []string) bool {

	for i := range crew1 {
		if crew1[i] != crew2[i] {
			return false
		}
	}
	return true
}
