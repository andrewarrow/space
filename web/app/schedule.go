package app

import "github.com/andrewarrow/feedback/router"

func Schedules(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		scheduleIndex(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		createSchedule(c)
		return
	}
	c.NotFound = true
}

func scheduleIndex(c *router.Context) {
	send := map[string]any{}
	items := c.All("schedule", "where user_id=$1 order by created_at desc", "",
		c.User["id"])
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
