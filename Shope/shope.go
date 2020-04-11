package main

import (
        "io/ioutil"
        "log"
		"strings"
		"fmt"
		"strconv"
		"os"
)

type Game struct {
	Name     string
	Quantity int
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func edit_db(db []Game, file string) {
	input, err := ioutil.ReadFile(file)
	check(err)

	lines := strings.Split(string(input), "\n")
	
	for _, game := range db {
		for l, line := range lines {
			if strings.Contains(line, game.Name) {
				lines[l] = game.Name + ":" + strconv.Itoa(game.Quantity)
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	check(err)
}

func edit_file() {
	input, err := ioutil.ReadFile("file")
	check(err)

	lines := strings.Split(string(input), "\n")
	

	for l, line := range lines {
		if strings.Contains(line, "Ps4") {
			lines[l] = "[]"
		}
	}
	
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("file", []byte(output), 0644)
	check(err)
}



// Load database
func load_db(file string) []Game {
	var db []Game

	content, err := ioutil.ReadFile(file)
	check(err)

	data := strings.Split(string(content), "\r\n")

	for _, v := range data {
		lines := strings.Split(v, ":")
		for i := 0; i < len(lines); i = i + 2{
			n, err := strconv.Atoi(lines[i+1])
			check(err)
			game := Game {
				Name: lines[i],
				Quantity : n,
			}
			db = append(db, game)
		}
	}

	return db
}

// information on about a game
func (g *Game) infos() {
	fmt.Printf("Name : %s\n", g.Name)
	fmt.Printf("Quantity: %d\n", g.Quantity)
	fmt.Printf("----------------------\n\n")
}


// edit game
func (g *Game) edit(numb int) {
	g.Quantity = g.Quantity - numb
}

func display_db(db []Game) {
	fmt.Println("----------- Games in stock -----------")
	for i := 0; i < len(db); i++ {
		db[i].infos()
	}
}

func search(db []Game) *Game{
	var game Game
	var name string

	fmt.Printf("\nWhat's the name ")
	fmt.Scanf("%s\n", &name)

	for i := 0; i < len(db); i++{
		if db[i].Name == name{
			return &db[i]
		}
	}
	return &game
}


func order(db []Game, panier *[]Game){
	var numb int
	order := true

	for order {
		game := search(db)

		if game == nil {
			fmt.Print("\n<nThis game doesn't exit\n\n")
			return
		}

		fmt.Printf("How many ?")
		fmt.Scanf("%d\n", &numb)
		if numb <= 0 || numb > game.Quantity {
			fmt.Printf("\n\nNumber not possible\n\n")
			return
		}

		item := Game{
			Name: game.Name,
			Quantity: numb,
		}
		*panier = append(*panier, item)
		fmt.Println("Continue ?\n1 for yes and 0 for no")
		fmt.Scanf("%d\n", &numb)
		if numb == 0 {
			order = false
		}
	}
	
	fmt.Println("\nYour panier is update")
}

// see panier
func show_panier(panier []Game) {
	fmt.Println("----------- Panier -----------")

	if panier == nil {
		fmt.Printf("Panier is empty\n\n")
		return
	}
	for i := 0; i < len(panier); i++ {
		panier[i].infos()
	}
}

func confirm_order() bool {
	var confirm bool
	fmt.Println("Are you sure to confirm your order ?\n1 for yes and 0 for no")
	fmt.Scanf("%t\n", &confirm)
	return confirm
}


func update_db(db []Game, panier *[]Game, confirm bool) {
	
	if confirm == false{
		fmt.Printf("\nYou may continue your purchase\n\n")
		return
	}
	if *panier == nil {
		return
	}
	pn := *panier

	for p := 0; p < len(pn); p++ {
		for d := 0; d < len(db); d++ {
			if pn[p].Name == db[d].Name {
				db[d].edit(pn[p].Quantity)
			}
		}
	}
	fmt.Printf("\nYour order is confirmed\n\n")
	*panier = nil

}

/* 
	Const define
*/
const (  // iota is reset to 0
	MENU_QUIT = iota
	MENU_SHOW_PANIER = iota
	MENU_ORDER = iota
	MENU_CONFIRM = iota
)

/*
About menu
*/


func do_menu(db []Game) int{
	var choice int
	display_db(db)
	
	fmt.Printf("[%d] Quit\n", MENU_QUIT)
	fmt.Printf("[%d] Show panier\n", MENU_SHOW_PANIER)
	fmt.Printf("[%d] order\n", MENU_ORDER)
	fmt.Printf("[%d] confirm your order\n", MENU_CONFIRM)
	fmt.Printf("> ")
	fmt.Scanf("%d\n", &choice)
	
	return choice
}


func process_choice(choice int, db []Game, p *[]Game){

	switch choice {
	case MENU_QUIT:
		defer fmt.Sprintln("Good bye :)")
		os.Exit(0)

	case MENU_SHOW_PANIER:
		show_panier(*p)

	case MENU_ORDER:
		order(db, p)
	
	case MENU_CONFIRM:
		show_panier(*p)
		update_db(db, p, confirm_order())
		edit_db(db, "db")

	default:
		fmt.Printf("Wrong input\n")
	}
}


func main() {

	fmt.Printf("------------ Welcome to the koolwx shop ---------------\n")
	database := load_db("db")
	var p []Game

	for {
		process_choice(do_menu(database), database, &p)
	}
	//edit_file()

}