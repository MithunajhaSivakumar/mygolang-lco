package main

import "fmt"

// Outside cannot use no var style
// jwtToken := 1000

// var name start with Uppercase - public variable
const LoginToken string = "abcdefgh"

func main() {
	var username string = "Mithunajha"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.787664432345678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit way of declaring variable
	var website = "Learning"
	fmt.Println(website)

	// no var style
	numberOfUser := 100
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
