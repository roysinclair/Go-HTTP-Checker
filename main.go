package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a Go Channel with the variable name of c
	c := make(chan string)

	for _, link := range links {
		/* NOTE: Create a Go Routine by adding the word "go" in front of a function only function can become a routine */
		go checkLink(link, c)
	}

	// Loop with this for loop indefinitely
	// As receiving stuff for a Channel is a blocking call this for loop will wait until it has received a reply and only then action
	// Wait for the channel to return some value after the channel has returned some value. Assign it to this variable l l in this case being short for link then run the body of the for loop
	for l := range c {

		/*
			NOTE: With routines that is that we never ever try to access the same variable from a different child routine where ever possible.
			We only share information with a child routine or a new routine that we create by passing it in as an argument or communicating with the child routine over channels.
			We never try to share variables directly between them otherwise for having had some really weird behavior
		*/

		// So we can create a function literal (JS,PHP calls it an anonymous function)
		go func(link string) {
			// Pause for 5 seconds
			time.Sleep(5 * time.Second)

			// Recheck the links
			checkLink(link, c)

		}(l) /* NOTE: <-- This set of parentheses is the set of parentheses that actually executes the function */
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send Link information back to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Send Link information back to the channel
	c <- link
}
