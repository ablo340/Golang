package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAXSIZE = 20

// readData reads each line of the specified text file,
// converting each line as a floating-point number
func readData(path string) [MAXSIZE]float64 {
	file, _ := os.Open(path)
	defer file.Close()

	var result [MAXSIZE]float64
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.ParseFloat(scanner.Text(), 64)
		result[i] = value
		i++
	}
	return result
}

// analyse goes through all the numbers of an array and performs
// several analyses, each result being associated with a string label
func analyse(data [MAXSIZE]float64) map[string]float64 {
	result := make(map[string]float64)
	result["first"] = data[0]
	result["random"] = random(data)
	return result
}

// random selects one random value from an array
func random(data [MAXSIZE]float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return data[rand.Intn(MAXSIZE)]
}

func main() {
	data := readData("data.txt")
	results := analyse(data)
	fmt.Println(results)
}