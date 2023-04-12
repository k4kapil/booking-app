package main

import (
	//"booking-app/helper"
	"fmt" // fmt -> format
	"strings"
	//"strings"
)

// var bookings = make([]map[string]string, 0) // empty List of Map (intializing list)
var bookings = make([]UserData, 0) // structure definiton & declaration

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// Variable
	//var conferenceName = "Go Conference"          //Mthod 1
	//var conferenceName string = "Go Conference"		//Method 2
	conferenceName := "Go Conference" //Method 3 -> only used for variables not constants

	//var bookings []string
	//bookings := []string{}  // empty slice

	/* 1. Declaring variables

	var test string
	test = "to test variable"

		2. To get type of data (datatype)
		fmt.Printf("ConferenceName type is %T,conferenceTickets type is %T\n", conferenceName, conferenceTickets)
	*/

	var remainingTickets uint = 50
	// COnstants
	const conferenceTickets = 50

	// Using function
	greeting(conferenceName, remainingTickets, conferenceTickets)

	// Loops   (In go -> there is only one loop i.e for loop)

	//for true {
	for {

		firstName, lastName, email, userTickets := getUserInput()                                                           // function
		isValidName, isValidEmail, isValidUserTicket := getInputValidation(firstName, email, userTickets, remainingTickets) // function

		if isValidName && isValidEmail && isValidUserTicket {
			fmt.Printf("The details for the user is  :\nFirst Name%5v\nLast Name%5v\nEmail Id%5v\nUser Tickets%5v\n", firstName, lastName, email, userTickets)
			// Remaining ticket logic
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			/*Arrays
			//(fixed size)

			var bookings [50]string
			bookings[0] = firstName + " " + lastName
			fmt.Printf("The whole array is : %v\n", bookings)
			fmt.Printf("The first element is : %v\n", bookings[0])
			fmt.Printf("Size of arrays is : %v\n", len(bookings))
			fmt.Printf("The type of array is : %T\n", bookings)
			*/

			// Map -> to store key-value pair

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}
			/*	var userData = make(map[string]string)
				userData["First Name"] = firstName
				userData["Last Name"] = lastName
				userData["Email"] = email
				userData["Number fo Tickets"] = strconv.FormatUint(uint64(userTickets), 10)
			*/

			// Slice
			//bookings = append(bookings, firstName+" "+lastName)
			bookings = append(bookings, userData)
			fmt.Printf("All the bookings are %v\n", bookings)

			//firstNames := getFirstNames(bookings) // function
			firstNames := getFirstNames()
			fmt.Println("The first names are :", firstNames)

			//If - else condition
			if remainingTickets == 0 {
				fmt.Println("All the conference tickets are booked. Try next time.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("first name or last name is too short (min 2 CHARACTERS).")
			}
			if !isValidEmail {
				fmt.Printf("email address does not conatin '@'.")
			}
			if !isValidUserTicket {
				fmt.Printf("You have entered invalid ticket number.")
			}
		}

		//fmt.Printf("The whole slice is : %v\n", bookings)
		//fmt.Printf("The first element is : %v\n", bookings[0])
		//fmt.Printf("Size of slice is : %v\n", len(bookings))
		//fmt.Printf("The type of slice is : %T\n", bookings)
	}
}

func greeting(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	//fmt.Println("Welcome to", conferenceName, "Booking Application !!!")
	fmt.Printf("Welcome to %v Booking Application !!!\n", conferenceName)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("*****Get your Tickets from here!!*******")
}

func getFirstNames() []string {
	// For each loop
	firstNames := []string{}
	// to ignore the a variable, we put _
	for _, booking := range bookings {
		//var names = strings.Fields(booking) // to split the array
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// user defined input
	fmt.Println("\nEnter user details :")
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter First Name :")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last Name :")
	fmt.Scan(&lastName)
	fmt.Println("Enter Email Id :")
	fmt.Scan(&email)
	fmt.Println("Enter user Tickets :")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func getInputValidation(firstName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// User Input Validations
	isValidName := len(firstName) > 2 && len(firstName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicket
}
