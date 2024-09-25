package main

func Div(nb1 int, nb2 int) float64 {
	if nb2 == 0 {
		return -1

	} else {
		var restulat float64 = float64(nb1) / float64(nb2)
		return restulat
	}
}
