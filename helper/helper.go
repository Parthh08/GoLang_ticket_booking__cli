package helper
import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {

    isValid := len(firstName) >= 2 && len(lastName) >= 2 && userTickets > 0 && userTickets <= remainingTickets && strings.Contains(email, "@")
    return isValid
}