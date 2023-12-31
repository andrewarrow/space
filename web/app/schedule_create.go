package app

import "github.com/andrewarrow/feedback/router"

func createSchedule(c *router.Context) {
	c.ReadJsonBodyIntoParams()

	c.ValidateCreate("schedule")
	c.Params["user_id"] = c.User["id"]
	c.Insert("schedule")
	guid := c.Params["guid"]
	row := c.One("schedule", "where guid=$1", guid)
	c.SendContentAsJson(row, 200)
}
func deleteSchedule(c *router.Context, guid string) {
	c.FreeFormUpdate("delete from schedules where guid=$1", guid)
	c.SendContentAsJson("", 200)
}
