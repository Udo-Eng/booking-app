package customer

import (
	"fmt"
	"strings"
	"booking-app/models"
)



func CreateUsersData(remainingTickets *uint) []model.UserData  {

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint


	var bookings  = []model.UserData{}


	var NewUserData model.UserData 

	for {


		if *remainingTickets <= 0 {
			fmt.Println("Our conference is totally booked out come back next Year thanks ")
			break 
		}
 
		// Performing User Validation on user Input 

		// To scan input of the user and pass the value to userName variable in Memory
		fmt.Println("\nEnter your First name ?")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name ?")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address ?")
		fmt.Scan(&emailAddress)

		fmt.Println("Enter the number of tickets ?")
		fmt.Scan(&userTickets)

		// Boolea Values to impplement the validation Logic 
		isValidName := len(firstName)  >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(emailAddress,"@")
		isValidUserTickets :=  userTickets > 0  &&  userTickets <= *remainingTickets

		if isValidEmail && isValidName && isValidUserTickets {

			NewUserData = model.UserData{
				firstName: firstName,
				lastName: lastName,
				emailAddress: emailAddress,
				tickets: userTickets,
			}
		
			*remainingTickets = *remainingTickets - userTickets
		
			fmt.Printf("Thank you %v %v for booking %v Tickets with us , you will receive a confirmation email at %v \n", firstName, lastName, userTickets, emailAddress)

			fmt.Printf("\n%v Tickets remaining for this %v\n", *remainingTickets, "Go Conference")


			bookings = append(bookings,NewUserData)
		
		} else {
			fmt.Println("Your Inputs were Invalid Reason:")

			if !isValidName {
				fmt.Println("Your first name and last name entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Your email doesn't contains an @ sign")
			}

			if !isValidUserTickets {
				fmt.Println("The ticket number you entered is invalid")
			}

			fmt.Println("Try Again")
			
		}

	}
	
	return bookings

}
