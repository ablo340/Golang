package main

import (
	"fmt"
	"encoding/gob"
	"os"
	"io/ioutil"
	"bytes"
)

/* 
	Const define
*/
const (  // iota is reset to 0
	MENU_QUIT = iota
	MENU_SHOW_ALL = iota
	MENU_SHOW_UNPAID = iota
	MENU_SEARCH = iota
	MENU_EDIT = iota
	MENU_SAVE = iota
	MENU_LOAD = iota
)


type Student struct{
	Firstname string
	Lastname string
	Age int
	Paid bool
}

/*
About Student
*/


// editer un student 
func (s *Student) edit(){
	fmt.Printf("First name (maximum 31 chars): ")
	fmt.Scanf("%s\n", &s.Firstname)
	fmt.Printf("Last name (maximum 31 chars): ")
	fmt.Scanf("%s\n", &s.Lastname)
	fmt.Printf("Votre age: ")
	fmt.Scanf("%d\n", &s.Age)
	fmt.Printf("Has paid? (1 for yes, 0 for no): ")
	fmt.Scanf("%t\n", &s.Paid)
}


// infos sur un student
func (s *Student) infos() {
	fmt.Printf("----- student's informations -----\n\n");
	fmt.Printf("Student first name: %s\n", s.Firstname)
	fmt.Printf("Student last name: %s\n", s.Lastname)
	fmt.Printf("Student year of birth: %d\n", s.Age)
	fmt.Printf("Has paid?: %t\n\n", s.Paid)
}


/*
About student list
*/

//Ajouter student(s)
func addStudent(s *[]Student){
	var n int
	fmt.Printf("Enter the number of students: ")
	fmt.Scanf("%d\n", &n);

	for i := 0; i < n; i++ {
		var st Student
		fmt.Printf("Entering information for student number %d:\n", i+1)
		st.edit()
		*s = append(*s, st)
	}
}


// Montrer la liste de tous les students
func allStudentInfos(s []Student){
	fmt.Printf("----------- Listing student details -----------\n\n")
	if s == nil {
		fmt.Printf("No student in the database\n")
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("----------- Details for student number %d -----------\n", i+1)
		s[i].infos()
	}
}


// Montrer les student qui n'ont pas payé
func show_unpaid(s []Student){
	fmt.Printf("----------- Listing unpaid student details -----------\n\n")
	for i := 0; i < len(s); i++ {
		fmt.Printf("----------- Details for student number %d -----------\n", i+1)
		if !s[i].Paid{
			s[i].infos()
		}
	}
}


// rechercher un student
func search(s []Student) *Student{
	var st Student
	var testfn string
	var testln string

	fmt.Printf("First name (maximum 31 chars): ")
	fmt.Scanf("%s\n", &testfn)
	fmt.Printf("Last name (maximum 31 chars): ")
	fmt.Scanf("%s\n", &testln)

	for i := 0; i < len(s); i++{
		if s[i].Firstname == testfn && s[i].Lastname == testln{
			return &s[i]
		}
	}
	return &st
}


/*
About menu
*/


func do_menu() int{
	var choice int

	fmt.Printf("[%d] Quit\n", MENU_QUIT)
	fmt.Printf("[%d] Show all\n", MENU_SHOW_ALL)
	fmt.Printf("[%d] Show unpaid\n", MENU_SHOW_UNPAID)
	fmt.Printf("[%d] Search name\n", MENU_SEARCH)
	fmt.Printf("[%d] Edit student\n", MENU_EDIT)
	fmt.Printf("[%d] Save\n", MENU_SAVE)
	fmt.Printf("[%d] Load\n", MENU_LOAD)
	fmt.Printf("> ")
	fmt.Scanf("%d\n", &choice)
	
	return choice
}


func process_choice(choice int, s []Student){
	switch choice {
	case MENU_QUIT:
		defer fmt.Sprintln("Good bye :)")
		os.Exit(0)

	case MENU_SHOW_ALL:
		allStudentInfos(s)

	case MENU_SHOW_UNPAID:
		show_unpaid(s)

	case MENU_SEARCH:
		st := search(s)
		if st.Firstname != ""{
			fmt.Printf("Found\n")
			st.infos()
		}else{
			fmt.Printf("Not found this student\n")
		}

	case MENU_EDIT:
		st := search(s)
		if st.Firstname != ""{
			fmt.Printf("Enter the modifications\n")
			st.edit()
			fmt.Printf("Success\nNew informations : ")
			st.infos()
		}else{
			fmt.Printf("Not found this student\n")
		}
		
	case MENU_SAVE:
		save(s)

	case MENU_LOAD:
		load(s)

	default:
		fmt.Printf("Wrong input\n")
	}
}


/*
About file
*/

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// save student in file
func save(s []Student){
	var namefile string
	
	fmt.Printf("----------- Saving students to file -----------\n\n")
	fmt.Printf("Enter file name: ")
	fmt.Scanf("%s\n", &namefile)

	f, err := os.Create(namefile)
	defer f.Close()
	if err != nil {
		fmt.Printf("ERROR: Could not open file \"%s\" for saving\n", namefile)
		return
    }
	

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(s)
	check(err)
}

// load student in file
func load(s []Student){
	var namefile string

	fmt.Printf("----------- Loading students to file -----------\n\n")
	fmt.Printf("Enter file name: ")
	fmt.Scanf("%s\n", &namefile)

	f, err := ioutil.ReadFile(namefile)
	if err != nil {
		fmt.Printf("ERROR: Could not open file \"%s\" for loading\n", namefile)
		return
	}
	
	decoder := gob.NewDecoder(bytes.NewReader(f))
	err = decoder.Decode(&s)
}


func main() {
	var s []Student

	fmt.Printf("------------ Welcome to the sailing thing program. ---------------\n");

	addStudent(&s)

	for {
		process_choice(do_menu(), s)
	}
}
