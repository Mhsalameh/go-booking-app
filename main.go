package main

import (
	"booking-app/common"
	"fmt"
	"sync"
	"time"
)
const conferenceTickets uint = 50
const conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData,0)
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


func greetUsers(){
	fmt.Printf("welcome to %v booking app\n", conferenceName)
	fmt.Printf( "%v Tickets available from total of %v\n",remainingTickets, conferenceTickets)	
}

func getFirstNames() []string{
	firstNames := []string{}
	for _ , booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func takeUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("your name: ")
	fmt.Scanln(&firstName)
	fmt.Print("last name: ")
	fmt.Scanln(&lastName)
	fmt.Print("email: ")
	fmt.Scanln(&email)
	fmt.Print("number of tickets to buy: ")
	fmt.Scanln(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets -= userTickets
	var userData = UserData {
		firstName :firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings,  userData)
	fmt.Printf("Thank you %v %v for your purchase of %v tickets, we will send you a confirmation email to %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("only %v tickets remaining in %v\n", remainingTickets, conferenceName )
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName )
	fmt.Println("#####################################")
	fmt.Printf("Sending ticket,\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################################")
	wg.Done()
}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()

	firstName, lastName, email, userTickets := takeUserInput()
	isValidEmail, isValidName, isValidTicket := common.ValidateUserInput(remainingTickets,firstName, lastName, email, userTickets)
	if !isValidEmail || !isValidName || !isValidTicket {
		if !isValidName{
			fmt.Println("Your name input is invalid, try again, name be of atleast length 2")
		}
		if !isValidEmail{
			fmt.Println("Your email input is invalid, try again, email should atleast include '@' symbol")
		}
		if !isValidTicket{
			fmt.Printf("Your ticket input is invalid, only %v tickets remaining, tickets can't be less or equal '0'\n", remainingTickets)
		}

	} else {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets , firstName, lastName, email)
		fmt.Printf("our bookings: %v\n", getFirstNames())

		if remainingTickets == 0 {
			fmt.Println("We are booked out, come early next year!")
			// break
		}
	}
	fmt.Println(bookings)
	wg.Wait()
}