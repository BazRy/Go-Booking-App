package helper

import (
	"fmt"
	"strings"
)

func GetValidInput(firstName string, lastName string, email string, ticketsRequired uint, ticketsRemaining uint) (bool, bool, bool, bool, bool) {

	isValidFirstName := len(firstName) > 1
	isValidLastName := len(lastName) > 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsRequired := ticketsRequired > 0
	isValidTicketNumbersRemaining := ticketsRequired <= ticketsRemaining

	return isValidFirstName, isValidLastName, isValidEmail, isValidTicketsRequired, isValidTicketNumbersRemaining
}

func GreetUsers(ticketsRemaining uint, conferenceTickets uint) {
	conferenceName := "Bazza's Conference"
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("Welcome To %v Bookings\n", conferenceName)
	fmt.Printf("Get your tickets here to attend, there are %v in total, with %v remaining\n", conferenceTickets, ticketsRemaining)
	fmt.Println("-------------------------------------------------------------------------")

}

func GetFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}
