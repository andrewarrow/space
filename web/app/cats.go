package app

import "github.com/andrewarrow/feedback/router"

var catList = []string{"gaming", "vehicles", "phones",
	"tablets", "laptops", "household", "music", "video", "assistants", "other"}

func Cats(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		cats(c)
		return
	}
	c.NotFound = true
}

func cats(c *router.Context) {
	send := map[string]any{}
	catMap := map[string]int{}
	items := c.All("device", "order by id", "")
	total := 0
	for _, item := range items {
		cat := item["cat"].(string)
		catMap[cat]++
		total++
	}
	send["cats"] = catList
	send["cat_map"] = catMap
	send["total"] = total
	c.SendContentAsJson(send, 200)
}
