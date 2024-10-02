package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicket(wg *sync.WaitGroup, userID int, totalTicketRemain *int) {
	defer wg.Done()

	mutex.Lock()
	if *totalTicketRemain > 0 {
		*totalTicketRemain--
		fmt.Printf("\nUserID: %d bought a ticket, remain ticket: %d", userID, *totalTicketRemain)
	}
	mutex.Unlock()
}

func usingMutex(totalTicket int) {
	var wg sync.WaitGroup
	for userID := 0; userID < 2000; userID++ {
		wg.Add(1)
		go buyTicket(&wg, userID, &totalTicket)
	}
	wg.Wait()
}
