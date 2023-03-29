package common

import "strings"

func ValidateUserInput(remainingTickets uint,userName string, lastName string, email string, userTickets uint) (bool,bool,bool){
	isValidName := len(userName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail,isValidName,isValidTicket
}