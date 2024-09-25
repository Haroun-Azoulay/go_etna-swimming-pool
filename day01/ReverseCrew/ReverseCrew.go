package main

func ReverseCrew(crew [10]string) [10]string {
	var slice [10]string
	for i := 0; i < 10; i++ {
		slice[i] = crew[10-i-1]
	}
	return slice
}
