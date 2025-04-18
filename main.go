package main

import (
    "booking_app/helper" // Custom helper package for validation
    "fmt"
    "strconv"
    "time"
)

// Main function: Entry point of the application
func main() {
    // Conference details
    conferenceName := "Go Conference"
    const conferenceTickets = 50
    var remainingTickets uint = 50

    // Slice to store bookings as a list of maps
    var bookings = make([]map[string]string, 0)

    // Greet users and provide conference details
    greetUsers(conferenceName, conferenceTickets, remainingTickets)

    // Infinite loop to keep the booking process running
    for {
        // Get user input for booking
        firstName, lastName, email, userTickets := getUserInput()

        // Validate user input
        if helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) {
            // Check if requested tickets exceed remaining tickets
            if userTickets > remainingTickets {
                fmt.Printf("We only have %v tickets remaining. Please try again.\n", remainingTickets)
                continue
            }

            // Book tickets and update the bookings list
            bookTicket(firstName, lastName, email, userTickets, &remainingTickets, &bookings)

            // Simulate sending tickets asynchronously
            go sendTicket(userTickets, firstName, lastName)

            // Display the first names of all attendees
            firstNames := getFirstNames(bookings)
            fmt.Printf("The first names of bookings are: %v\n", firstNames)

            // Display remaining tickets
            fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

            // Check if tickets are sold out
            if remainingTickets == 0 {
                fmt.Println("Our conference is fully booked. Come back next year!")
                break
            }
        } else {
            fmt.Println("Invalid input. Please try again.")
        }
    }

    // Example of a switch statement for city-based messages
    city := "London"
    switch city {
    case "New York":
        fmt.Println("You are in New York.")
    case "London", "Paris":
        fmt.Println("You are in London or Paris.")
    default:
        fmt.Println("You are in an unknown city.")
    }
}

// greetUsers: Displays a welcome message and conference details
func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
    fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
    fmt.Printf("We have a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets now to attend this amazing event!")
    fmt.Println("---------------------------------------------------")
}

// getFirstNames: Extracts and returns the first names of all attendees
func getFirstNames(bookings []map[string]string) []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking["firstName"])
    }
    return firstNames
}

// getUserInput: Collects user input for booking tickets
func getUserInput() (string, string, string, uint) {
    var firstName, lastName, email string
    var userTickets uint

    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email address:")
    fmt.Scan(&email)

    fmt.Println("Enter the number of tickets you want to book:")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}

// bookTicket: Books tickets for the user and updates the bookings list
func bookTicket(firstName string, lastName string, email string, userTickets uint, remainingTickets *uint, bookings *[]map[string]string) {
    // Deduct the number of tickets booked from the remaining tickets
    *remainingTickets -= userTickets

    // Create a map to store user details
    var userData = make(map[string]string)
    userData["firstName"] = firstName
    userData["lastName"] = lastName
    userData["email"] = email
    userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

    // Add the user data to the bookings list
    *bookings = append(*bookings, userData)

    // Display booking confirmation
    fmt.Println("---------------------------------------------------")
    fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, userTickets, email)
    fmt.Println("---------------------------------------------------")
    fmt.Println("Current list of bookings:")
    fmt.Println(*bookings)
}

// sendTicket: Simulates sending tickets to the user (asynchronous)
func sendTicket(userTickets uint, firstName string, lastName string) {
    time.Sleep(10 * time.Second) // Simulate a delay in sending the ticket
    ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("#########################")
    fmt.Printf("Sending ticket:\n%v\n", ticket)
    fmt.Println("#########################")
}