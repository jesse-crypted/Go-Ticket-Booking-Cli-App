package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	// loop through the program
	for {
		// defining user variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask for their name

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		// user input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// check if userTicket input is greater than remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("The length of slice: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("The first names of users who booked are: %v\n", firstNames)

			// check if tickets are sold out and stop its execution
			if remainingTickets == 0 {
				fmt.Println("Chief!, Tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firstname or lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address input is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// continue
		}
	}
}
