package app

import (
	"github.com/andrewarrow/feedback/router"
)

func createTicket(c *router.Context) {
	c.ReadJsonBodyIntoParams()

	c.ValidateCreate("ticket")
	c.Insert("ticker")
	guid := c.Params["guid"]
	row := c.One("ticket", "where guid=$1", guid)
	c.SendContentAsJson(row, 200)
}
