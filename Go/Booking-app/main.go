package main // every file MUST be in a package

import (
	"fmt" // to use Print functionality etc
	"sync"
	"time"
)

// package level variables can be accessed across files in same package
// cannot be created using :=
var conferenceName = "Go Conference" // Go is statically typed, but uses Type Inference
const conferenceTickets = 50 // use of constants
var remainingTickets uint = 50	
// Declaration of array in Go: var bookings [50]string
// Syntax for initialization: var bookings = [50]string{}
// Note: [] is slice, [number] is array
// Previous slice version: var bookings = []string{}
// Previous map version: var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

// struct in Go
// Go structs can have functions
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// WaitGroup created so that main thread waits for sendTicket thread
// if 50 tickets sold out before priting sendTicket
var wg = sync.WaitGroup{}

func main() {
	greetUser()

	// Only have for loop, no while, do-while, foreach in Go
	// Syntax: for <condition> {}
	for {
		
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTickets := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidEmail || !isValidName || !isValidTickets {
			fmt.Println("Invalid Input")
			continue
		}
		
		bookTicket(userTickets, firstName, lastName, email)
		// go ... starts a new goroutine
		// goroutine is a lightweight thread managed by the Go runtime
		// Add() sets the number of goroutines to wait for by increasing a counter internally
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("List of bookings: %v\n", firstNames)
		
		// Conditionals in Go
		if remainingTickets == 0 {
			fmt.Println("Our connference is booked out. Come back next year.")
			break
		}
	}
	// Wait blocks until WaitGroup counter is 0
	wg.Wait()
}

// functions
func greetUser() {
	// Use of Printf
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// Use , in Print
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	// range iterates over elements for different DTs
	// for array and slices, range provides index and value
	for _, booking := range bookings {
		//previous version when booking is []string
		//var names = strings.Fields(booking) // same as split in Java
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//input output in Go
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// maps in Go
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// array: bookings[0] = userData
	bookings = append(bookings, userData)

	fmt.Printf("type: %T\n", bookings)
	fmt.Printf("length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickers. You will receive a confirmation email at %v\n", 
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Simulate workload
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	// Done() decrements the WaitGroup counter by 1
	wg.Done()
}