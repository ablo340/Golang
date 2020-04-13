package main

import "fmt"

func slice_nil(a *[]int){
	*a = nil
}

func slice_append(a *[]int){
	*a = append(*a, 1, 2, 3, 4)
}

func subtactOne(numbers []int) {  
    numbers[0] = 2
	fmt.Println("Modified slice: ", numbers)
}

func myfun(element []string) { 
  
    // Here we only modify the slice 
    // Using append function 
    // Here, this function only modifies 
    // the copy of the slice present in  
    // the function not the original slice 
    element = append(element, "Java") 
    fmt.Println("Modified slice: ", element) 
} 

func main() {
	/*var a, b string = "Jean", "Baptiste!"
	fmt.Printf("Bonjour monsieur %s %s", a, b)*/
	//var d [] int
	/*da := make ([] int , 3)
	slice_nil(&da)
	fmt.Println(da)*/
	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
	// Creating slice 
    slc := []string{"C#", "Python", "C", "Perl"} 
      
    fmt.Println("Initial slice: ", slc) 
  
    // Passing the slice to the function 
    myfun(slc) 
      
    fmt.Println("Final slice:", slc)

}
