package main

import "fmt"


// checks whether a given string contains a given character or not
/*func contains_error(sentence string, character string) (string, int) {
	for i, char := range sentence{
		if char == character{ // char ou sentence[i] est un type d'octet => casse explicite
			return character, true
		}
	}
	return string(character)
}*/


// checks whether a given string contains a given character or not
func contains_fix(sentence string, character string) (string, bool) {
	for _, char := range sentence{
		if string(char) == character{ // char ou sentence [i ] est un type d'octet => casse explicite
			return string(char), true
		}
	}
	return character, false
}

func main() {
	fmt.Println(contains_fix("je suis la", "z"))
}
