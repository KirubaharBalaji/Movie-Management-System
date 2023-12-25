package main

import (
	"fmt"
	"strings"
)

var cinemaName string
var locationTheatre string
var movieName string

var firstName string
var lastName string
var email string
var mobileNumber int

func main() {

	var theatre int
	var locationTheatreChoice int
	var movieChoice int
	var fullName string
	var confirmChoice string

	//function call for userInput and assigns value to the firstName, lastName,email and MobileNumber
	firstName, lastName, email, mobileNumber = userInput()
	// function call for getName and assigns value to the fullName
	fullName = getName(firstName, lastName)
	// prints the details of the user
	fmt.Printf("Name:%v\nEmail:%v\nMobile Number:%v\n", fullName, email, mobileNumber)
	//function call for userConfirmation and assigns the new value for the firstName, lastName, email and MobileNumber
	firstName, lastName, email, mobileNumber = userConfirmation()
	//prints the details of the user after changes
	//fmt.Printf("Details aFter Change:\n1.First Name:%v\n2.Last Name:%v\n3.Email:%v\n4.Mobile Number:%v\n", firstName, lastName, email, mobileNumber)
	// funcction call for userValidator and assigns true(or)false for isValidName,isValidMobileNumber,isValidEmail
	isValidName, isValidEmail, isValidMobileNumber := userValidator(firstName, lastName, email, mobileNumber)
	//if everythng is valid we proceed to get ticket count
	if isValidName && isValidEmail && isValidMobileNumber {
		greetInUser(fullName)
	} else {
		fmt.Printf("your Inputs are Invalid.\n Please check your input and try again.\n")
		if !isValidEmail {
			fmt.Printf("your Email is Invalid %v.\n", email)
		}
		if !isValidName {
			fmt.Printf("Your Name is Invalid %v  %v.\n", firstName, lastName)
		}
		if !isValidMobileNumber {
			fmt.Printf("Your Mobile Number is Invalid %v.\n", mobileNumber)
		}
	}

	// Input for the theatre name
	fmt.Printf("The Following Theatres are available :\n 1.Arnold Cinemas\n 2.Rocky Cinemas\n 3.LUXE Cinemas\n 4.AGS Cinemas\n 5.INOX Cinemas\n 6.PVR Cinemas\n 7.SPI Cinemas\n")
	fmt.Println("Enter Your Choice...:")
	fmt.Scan(&theatre)
	// function call for getTheatre
	getTheatre(theatre)
	// Input for the theatre Location
	fmt.Printf("The Following Theatres are available in the Following Locations :\n 1.Aynavaram\n 2.Kolathur\n 3.Perambur\n 4.Mount Road\n 5.Kilpauk\n 6.T.Nagar\n 7.Numngambakkam\n")
	fmt.Println("Enter Your Choice...:")
	fmt.Scan(&locationTheatreChoice)
	// function call for getLocation
	getLocation(locationTheatreChoice)
	// Input for the Movie
	fmt.Printf("The Following Movies are available :\n 1.Vikram\n 2.Beast\n 3.Mamanidhan\n 4.K.G.F Chapter 2\n 5.Valimai\n 6.Don\n 7.Mahaan\n")
	fmt.Println("Enter Your Choice...:")
	fmt.Scan(&movieChoice)
	// function call for getMovieName
	getMovieName(movieChoice)
	// prints the details of theatre name, location, MovieName
	fmt.Printf("Theatre Name: %v\nLocation: %v\nMovie: %v\n", cinemaName, locationTheatre, movieName)
	fmt.Printf("Do you want to change your choice? (yes/no): \n")
	fmt.Scan(&confirmChoice)
	confirmChoice = strings.ToUpper(confirmChoice)
	if confirmChoice == "YES" {
		confirmator()
		afterConfirmation()
	} else {
		fmt.Print("OK \n")
	}

}
func greetInUser(fullName string) {
	fmt.Printf("Welcome %v to Ticket-Surfer\n", fullName)
}

func ticketer(){
	
}