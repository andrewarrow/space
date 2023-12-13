package app

import "github.com/andrewarrow/feedback/router"

func Ticket(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		ticketIndex(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		createTicket(c)
		return
	}
	c.NotFound = true
}

func ticketIndex(c *router.Context) {
	send := map[string]any{}
	items := c.All("ticket", "order by created_at desc", "")
	send["items"] = items
	c.LayoutMap["wasm"] = makeWasmScript("tickets")
	c.SendContentInLayout("tickets.html", send, 200)
}
