package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"time"
	"os"
)

type ContactBook struct {
	Name     string
	Contacts []map[string]string
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		//log.Fatal(e)
		//return
		os.Exit(0)
		//panic(e)
	}
}

// meanBirthYear computes the mean birth year
// of the contacts from a given contact book
func meanBirthYear(cb *ContactBook) float64 {
	var mean float64

	for i := 0; i < len(cb.Contacts); i++ {
		birth, err := time.Parse(time.RFC3339, cb.Contacts[i]["birthday"])
		check(err)
		year := birth.Year()
		mean += float64(year)
	}

	mean /= float64(len(cb.Contacts))

	return mean
}

func main() {

	cb := ContactBook{}
	content, err := ioutil.ReadFile("contacts.json")
	check(err)

	err = json.Unmarshal(content, &cb)
	check(err)

	fmt.Printf("Mean birth year: %.2f\n", meanBirthYear(&cb))
}
