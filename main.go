package main

import "fmt"

func main() {
	var conferenceName = "our conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("conference name: %T conference Tickets: %T remaining tickets: %T", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend tickets avaliable is: %v avaliable ticket is: %v \n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	fmt.Printf("userName: %v user ticket is %v \n", userName, userTickets)
	fmt.Printf("remaing tickets are: %v \n", remainingTickets)
}
