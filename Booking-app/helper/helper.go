package helper

import "strings"

func UserInputValidation(firstName string, lastName string, userTickets uint, email string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isvaliduserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidemail, isvaliduserTickets
}
