package main

import (
	"fmt"
	"go-booking-app/helper"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	ticketsRequired uint
}

var wg = sync.WaitGroup{}

func main() {

	const conferenceTickets uint = 50
	var ticketsRemaining uint = 50

	helper.GreetUsers(ticketsRemaining, conferenceTickets)

	//var bookings = make([]map[string]string, 0)
	var bookings = make([]UserData, 0)

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

		//var userData = make(map[string]string)
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["email"] = email
		// userData["ticketsRequired"] = strconv.FormatUint(uint64(ticketsRequired), 10)
		var userData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			ticketsRequired: ticketsRequired,
		}

		bookings = append(bookings, userData)
		fmt.Printf("List of booking %v\n-----------------------------------------\n", bookings)

		fmt.Printf("Thanks %v %v, we have reserved %v tickets for you.  Confirmation sent to %v\n", firstName, lastName, ticketsRequired, email)
		println("")
		fmt.Printf("There are now %v tickets left\n", ticketsRemaining)
		println("")

		firstNames := getFirstNames(bookings)

		fmt.Printf("Current bookings are %v\n-----------------------------------------\n", firstNames)

		wg.Add(1)
		go sendTicket(ticketsRequired, firstName, lastName, email)

	}
	wg.Wait()
	fmt.Println("All tickets sold,  closing ticket office!\n-----------------------------------------")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func sendTicket(ticketsRequired uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("Sending %v Tickets to %v %v ", ticketsRequired, firstName, lastName)
	fmt.Println("\n###############################################")
	fmt.Printf("%vat email %v\n", ticket, email)
	fmt.Println("###############################################\n\n\n")
	wg.Done()
}
