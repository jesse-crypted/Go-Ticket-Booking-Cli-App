package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// Declare package level variable
const conferenceTickets = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

// var bookings = make([]map[string]string, 0)

// get time
var now time.Time = time.Now()
var formatted string = now.Format(time.RFC3339)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	time            string
}

func main() {
	greetUsers() // function call

	// loop through the program
	for {

		firstName, lastName, email, userTickets := getUserInput() // function to get user inputs
		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// check if userTicket input is greater than remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email, formatted)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
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

func greetUsers() {

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, time string) {

	remainingTickets = remainingTickets - userTickets

	//CREATE MAP FOR USER
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastname"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Struct
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
		time:            formatted,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // stops the execution for 10secs
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("####################")

}
