package main

import "fmt"
import "math"
import "errors"

// returns the maximal number from the three number received in parameter;
func max(a int, b int, c int) int {

	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}

// checks whether a given string contains a given character or not
func contains(sentence string, character string) bool {
	for _, char := range sentence {
		if string(char) == character { // char ou sentence [i ] est un type d'octet => casse explicite
			return true
		}
	}
	return false
}

// computes the factorial of a given natural number
func fact(num float64) float64 {
	if num == 0 { // fact de 0
		return 1.0
	}

	if num < 0 || math.Mod(num, 1) != 0 { //  fact de nombres negatifs et nombre à virgule(Erreur)
		return -1.0
	}

	var result = 1.0
	for num > 1 { // fact de nombres positifs
		result *= num
		num--
	}

	return result
}

func facte(n int) (int, error) {
	if n < 0 {
		return -1, errors.New(" Fact . of negative number undefined .")
	}
	if n == 0 {
		return 1, nil
	}

	prev, _ := facte(n - 1)
	return n * prev, nil
}

func main() {
	fmt.Println(facte(5.0))
}
