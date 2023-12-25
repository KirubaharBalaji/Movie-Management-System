package main

import (
	"fmt"
	"strings"
)

func userInput() (string, string, string, int) {
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email Address:")
	fmt.Scan(&email)
	fmt.Println("Enter your Mobile Number:")
	fmt.Scan(&mobileNumber)
	return firstName, lastName, email, mobileNumber
}
func getName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func userConfirmation() (string, string, string, int) {
	var choice string
	var confirmChoice int
	fmt.Printf("Are you Sure about Your Choices?\nYES (OR) NO\n")
	fmt.Scan(&choice)
	choice = strings.ToUpper(choice)
	if choice == "NO" {
		fmt.Printf("Please choose what value you want to change:\n")
		fmt.Printf("1.firstName:%v\n2.lastName:%v\n3.email:%v\n4.mobileNumber:%v\n", firstName, lastName, email, mobileNumber)
		fmt.Scan(&confirmChoice)
		switch confirmChoice {
		case 1:
			fmt.Printf("Enter your First Name:\n")
			fmt.Scan(&firstName)
		case 2:
			fmt.Printf("Enter Your LastName:\n")
			fmt.Scan(&lastName)
		case 3:
			fmt.Printf("Enter Your Email Address:\n")
			fmt.Scan(&email)
		case 4:
			fmt.Printf("Enter Your Mobile Number Numbers:\n")
			fmt.Scan(&mobileNumber)
		default:
			fmt.Printf("Invalid Choice")
			fmt.Printf("Details aFter Change:\n1.First Name:%v\n2.Last Name:%v\n3.Email:%v\n4.Mobile Number:%v\n", firstName, lastName, email, mobileNumber)
		}
	} else if choice == "YES" {
		fmt.Printf("\n")
	} else {
		fmt.Println("Invalid Choice")
	}
	return firstName, lastName, email, mobileNumber
}

func iterativeDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

func userValidator(firstName string, lastName string, email string, mobileNumber int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".com")
	isValidMobileNumber := iterativeDigitsCount(mobileNumber) == 10
	return isValidName, isValidEmail, isValidMobileNumber
}
