package main

import (
	"fmt"
	"go-booking-app/helper"
)

func main() {

	const conferenceTickets uint = 50
	var ticketsRemaining uint = 50

	helper.GreetUsers(ticketsRemaining, conferenceTickets)

	var bookings []string

	for ticketsRemaining > 0 {
		var firstName string
		var lastName string
		var email string
		var ticketsRequired uint

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email")
		fmt.Scan(&email)
		fmt.Printf("Hi %v %v, how many tickets would you like?\n", firstName, lastName)
		fmt.Scan(&ticketsRequired)

		isValidFirstName, isValidLastName, isValidEmail, isValidTicketsRequired, isValidTicketNumbersRemaining := helper.GetValidInput(firstName, lastName, email, ticketsRequired, ticketsRemaining)

		if !isValidFirstName {
			fmt.Print("You have entered an invalid first name (minimum 2 characters),  try again\n-----------------------------------------\n")
			continue
		}

		if !isValidLastName {
			fmt.Print("You have entered an invalid last name (minimum 2 characters),  try again\n-----------------------------------------\n")
			continue
		}

		if !isValidEmail {
			fmt.Print("You have entered an invalid email (missing @),  try again\n-----------------------------------------\n")
			continue
		}

		if !isValidTicketsRequired {
			fmt.Print("You have entered an invalid number of tickets required, must be more than zero,  try again\n-----------------------------------------\n")
			continue
		}
		if !isValidTicketNumbersRemaining {
			fmt.Print("You have entered an invalid number of tickets required, exceeds number available,  try again\n-----------------------------------------\n")
			continue
		}

		println("")

		ticketsRemaining -= ticketsRequired

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thanks %v %v, we have reserved %v tickets for you.  Confirmation sent to %v\n", firstName, lastName, ticketsRequired, email)
		println("")
		fmt.Printf("There are now %v tickets left\n", ticketsRemaining)
		println("")

		firstNames := helper.GetFirstNames(bookings)

		fmt.Printf("Current bookings are %v\n-----------------------------------------\n", firstNames)

	}
	fmt.Println("All tickets sold,  closing ticket office!\n-----------------------------------------")
}
