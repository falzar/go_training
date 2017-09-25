package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://jyodroid.github.io",
	}

	// "Chanels are the way we comunicate between routines"
	// "Routines have priority level"
	// "Main routine have bigger priority than other routines"
	// "Main have to be awared about other routines"
	c := make(chan string)

	// send data into the channel: channel <- value
	// Wait for a value (Specific) to be sent into the channel: value <- channel
	// Wait for a value (ANY) to be sent into a channel: <- channel

	for _, site := range sites {
		// go reserved word creates routines
		// is great for blocking opperations
		// _go scheduler_ uses one CPU by default, but also can run several CPU at time.
		// "concurrency is not parallelism"

		go checkSite(site, c)
		// After every go routine, the main program go and wait an answer for the channel
	}

	// Receive messages from channel
	// fmt.Println(<-c) //blocking code, we just receive the quickest one

	// what if i use a wait statement for each site
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(i, " ", <-c) // Still blocking
	// }

	// Execute check perpetually
	// for {
	// 	go checkSite(<-c, c)
	// }

	for s := range c { // The range keyword will take every channel return

		// function literal is equivalent to an anonimus function (wrap some of code)
		// go func() { //Using the same s of outer routine
		go func(s string) {
			time.Sleep(5 * time.Second) //blocking operation. Try to no use it on main routine
			checkSite(s, c)             // Function literal capturing s range variable
			// never attempt to use trhe same variable inside different go routines
			// provide the argument to the child routine (copied value)
			// }() Without the s argument to use the same outer variable (dont)
		}(s)
	}

}

func checkSite(s string, c chan string) {

	down := "Site migth be down!"
	up := "Is up"

	// With _ we can ignore the return argument
	_, err := http.Get(s)

	if err != nil {
		fmt.Println(s, down, err)
		// send a message to the channel
		// c <- down
		c <- s
		return
	}

	fmt.Println(s, up)
	// c <- up
	// Sleep pauses the current go routine for at least the duration d
	// time.Sleep(10 * time.Second) //blocking operation. Try to no use it on main routine
	c <- s
}
