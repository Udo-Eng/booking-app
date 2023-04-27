package main

import (
	"fmt"
	"time"
	"booking-app/models"
	"booking-app/customers"
)

var ConferenceName = "Go Conference"



func main() {


	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Hello welcome to ", ConferenceName, "Booking Application")

	fmt.Printf("We have a total of %v and  %v Tickets are still available \n", conferenceTickets, remainingTickets)


	fmt.Println("Get Your Tickets here to attend ")

	// Declaring an Array to store the  details of each user

	var bookings []model.UserData

// for loop to kep the application runing until the Tickets have been  fully booked 
 for {
		
		bookings = customer.CreateUsersData(&remainingTickets)


		
		for _,value  := range bookings {

			fmt.Printf("hello %v Thanks for booking with us \n",value.firstName)
			go sendTicket(value.tickets,ConferenceName,value.emailAddress)
			fmt.Println(value)
	
		}

		if remainingTickets <= 0 {
			break 
		}


	}
	
}


// Create a Function to ssend a Ticket
func sendTicket (numberTickets uint,ConferenceName string,email string){

	// use the time package to simulate a blocking operation on this single Thread 

	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("You bok %v tickets for the %v conference",numberTickets,ConferenceName)

	fmt.Println("##############")

	fmt.Printf("sending %v ticket to %v \n",ticket,email)
	fmt.Println("##############")
}


// To manipulate strig the strings package should be used 

// e.g split a string into an Array you need the 
// strings.Fields  function in the strings package 


// a  blank identifier is used to accept a value that would not be used in Golang programming 