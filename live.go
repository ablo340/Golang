package main

import "fmt"


// checks whether a given string contains a given character or not
func count_voy(sentence string) int {
	r := 0
	for _, char := range sentence{
		c := string(char)
		if c == "a" || c == "e" || c == "o" || c == "i" || c == "y" || c == "u" {
			r++
		}
	}
	return r
}

func main() {
	fmt.Println(count_voy("je suis la"))
}
