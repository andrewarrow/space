package app

import "github.com/andrewarrow/feedback/router"

func Ticket(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		createTicket(c)
		return
	}
	c.NotFound = true
}
