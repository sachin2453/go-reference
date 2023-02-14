package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go conference")
	message := "Go lang" // its a var
	const tickets = 50
	var remainingtickets uint = 50
	var cost float64 = 100
	fmt.Println("1) Let's learn", message)
	fmt.Println("2) You have", remainingtickets, "of", tickets, "tickets")
	fmt.Printf("3) Ticket balance:  %v tickets of %v tickets\n", remainingtickets, tickets)
	fmt.Printf("4) Ticket balance:  %T tickets of %T tickets\n", remainingtickets, tickets) // type of variable
	fmt.Printf("5) Each tickts costs $%v\n", cost)
	fmt.Println("6) memory location of remainingtickets is ", &remainingtickets, ".")
	var userName string
	var firstName string
	var lastName string
	var email string
	var pin uint
	// Ask for user name
	fmt.Println("Fill form")
	fmt.Print("\tEnter Username : ")
	fmt.Scan(&userName)
	fmt.Print("\tEnter pin : ")
	fmt.Scan(&pin)
	isValid := len(userName) >= 4
	fmt.Print("\tUsername is isValid :", isValid)
	fmt.Print("\tfirstName : ")
	fmt.Scan(&firstName)
	fmt.Print("\tlastName : ")
	fmt.Scan(&lastName)
	fmt.Print("\temailAddress : ")
	fmt.Scan(&email)
	isValidEmail := strings.Contains(email, "@")
	fmt.Print("\t isValidEmail : ", isValidEmail)
	fmt.Print("\tBuy tickets : ")
	fmt.Scan(&remainingtickets)
	totalCost := float64(remainingtickets) * cost
	fmt.Printf("****Thank for the regristration %v. You're total cost is  %v **** \n", firstName, totalCost)
	var array_letters = [50]string{"a", "b", "c"}
	fmt.Print("7) array_letters : \n\t", array_letters, "\n")
	var bookings = [50]string{}
	bookings[0] = "Dear " + firstName
	bookings[1] = lastName
	fmt.Print("\t", bookings, "\n")
	fmt.Printf("\tType : %T \n", bookings[0])
	fmt.Printf("\tlength : %v \n ", len(bookings[0]))
	var array_alphanum = []string{}
	fmt.Print("8) slice_alnum \n ")
	array_alphanum = append(array_alphanum, "a1")
	print("\t", array_alphanum[0], "\n")
	print("\t", array_alphanum)
}
