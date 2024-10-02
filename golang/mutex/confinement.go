package main

import (
	"fmt"
	"sync"
)

func buyTicketConfinement(wg *sync.WaitGroup, ticketChan chan int, userID int) {
	defer wg.Done()
	ticketChan <- userID
}

func manageTicketconfinenment(ticketChan chan int, doneChan chan bool, totalTickets *int) {
	for {

		select {
		case userID := <-ticketChan:
			if *totalTickets > 0 {
				*totalTickets--
				fmt.Printf("\nUserID: %d bought a ticket, remain ticket: %d", userID, *totalTickets)
			}
		case <-doneChan:
			fmt.Println("All tickets are sold")
		}

	}
}

func usingConfinement(totalTicket int) {
	var wg sync.WaitGroup
	ticketChan := make(chan int)
	doneChan := make(chan bool)

	go manageTicketconfinenment(ticketChan, doneChan, &totalTicket)

	for userID := 0; userID < 2000; userID++ {
		wg.Add(1)
		go buyTicketConfinement(&wg, ticketChan, userID)
	}

	wg.Wait()
	doneChan <- true
}
