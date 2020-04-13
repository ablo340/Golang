package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"log"
	"errors"
)

func check(e error) {
    if e != nil {
		log.Fatal(e)
	}
}

// readData reads each line of the specified text file,
// converting each line as a floating-point number
func readData(path string) ([]float64, error)  {
	file, e := os.Open(path)
	defer file.Close()

	var result []float64

	if e != nil {
		return result, e
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, e := strconv.ParseFloat(scanner.Text(), 64)
		if e != nil {
			return result, e
		}
		result = append(result, value)
	}
	return result, nil
}

// analyse goes through all the numbers of an array and performs
// several analyses, each result being associated with a string label
func analyse(data []float64) map[string]float64 {
	result := make(map[string]float64)
	result["first"] = data[0]
	result["random"] = random(data)
	result["pair"] = pair(data)  // part coding in live
	return result
}

func greatest(data []float64) float64 {
	var result float64 = 0.0
	for _, v := range data{
		value := float64(v)
		if value > result {
			result = value
		}
	}
	
	return result
}


// number of pair number
func pair(data []float64) float64 {
	var result float64 = 0.0
	for _, v := range data{
		if int(v)%2 == 0  {
			result++
		}
	}
	
	return result
}


// sum of all number in file
func sum(data []float64) float64 {
	var result float64 = 0.0
	for _, v := range data{
		result = result + float64(v)
	}
	
	return result
}

// random selects one random value from an array
func random(data []float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return data[rand.Intn(len(data))]
}

// factorial
func fact(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Fact . of negative number undefined .")
	}
	if n == 0 {
		return 1, nil
	}

	prev, _ := fact(n - 1)
	return n * prev, nil
}

func main() {
	data, err := readData("data.txt")
	check(err)
	results := analyse(data)
	fmt.Println(results)
}