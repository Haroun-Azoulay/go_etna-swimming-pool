package main

func RemoveMember(crew []string, index int) []string {
	return append(crew[:index], crew[index+1:]...)
}
