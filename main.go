package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 30
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Welcome to %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Hailey", "Honda", "Alice"}
	var bookings = [50]string // 空陣列不一定要加 {}  var bookings = [50]string{}

	var firstName string
	var lastName string
	var email string
	var userTickets int32

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName) // save the value that user enters
	fmt.Printf("This is a pointer special variable %v\n", &lastName)
	fmt.Printf("This is the variable \"%v\" that is pointed by above\n", &lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array: %v")

	fmt.Printf("Thank you %v %v for booking %v tickets.\n You will receive a confirmation email at %v\n77g", firstName, lastName, userTickets, email)
}
