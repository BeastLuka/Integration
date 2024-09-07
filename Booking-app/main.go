package main

import (
	"Booking-app/helper"
	"fmt"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)         // map

var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

func main() {

	// function to greet users
	greetUsers()

	for len(bookings) < 50 && remainingTickets > 0 {

		//func to get user input
		firstName, lastName, email, userTickets := getUserInput()

		//func to validatye user-input
		isValidName, isValidemail, isvaliduserTickets := helper.UserInputValidation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidemail && isvaliduserTickets {

			//func to book ticket

			bookTicket(firstName, lastName, userTickets, email)
			go sendTicket(firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
			fmt.Printf("the booking data is : %v\n", bookings)

			// iterating through a loop to get the firstnames of users in a slice

			//function to print firstnames
			firstNames := getFirstNames()
			fmt.Printf("The firstnames of booked users is : %v\n", firstNames)

			//checking tickets availability

			if remainingTickets == 0 {
				fmt.Println("we are booked out. please try next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("your firstname or lastname should be morethan 2 characters.")
			}
			if !isValidemail {
				fmt.Println("your email should contain @ sign  ")
			}
			if !isvaliduserTickets {
				fmt.Println("number of tickets is invalid ")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v \n", conferenceName)
	fmt.Printf("Welcome to  %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v tickets are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Grab as fast as you can to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	// index refers to getting the order of elements in the slice
	// for index, booking := range bookings {
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint, string) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("enter your no of Tickets: ")
	fmt.Scan(&userTickets)
	fmt.Println("enter your email: ")
	fmt.Scan(&email)

	return firstName, lastName, userTickets, email
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - (userTickets)

	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	// for map data type
	// var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)

	fmt.Printf("list of bookings : %v", bookings)
	fmt.Printf("user %v %v booked %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("congratualtions, you will get a confirmation mail to %v\n", email)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("sending %v tickets to the user %v %v\n", userTickets, firstName, lastName)

	fmt.Println("############################")
	fmt.Printf("sent %v details to %v address successfully\n", ticket, email)

	fmt.Println("############################")
}
