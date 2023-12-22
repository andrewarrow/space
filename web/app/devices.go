package app

import (
	"regexp"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func Devices(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		devicesIndex(c)
		return
	}
	if second == "functions" && third != "" && c.Method == "GET" {
		devicesFunctions(c, third)
		return
	}
	if second == "payloads" && third != "" && c.Method == "GET" {
		devicesPayloads(c, third)
		return
	}
	if second == "schedule" && third != "" && c.Method == "POST" {
		deviceSchedule(c, third)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		devicesByCat(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "PATCH" {
		devicePatch(c, second)
		return
	}
	c.NotFound = true
}

func devicesIndex(c *router.Context) {
	c.TableJson("device")
}

var functionRegex = regexp.MustCompile(`\((.*?)\)`)

func devicesFunctions(c *router.Context, guid string) {
	device := c.One("device", "where guid=$1", guid)
	send := map[string]any{}
	items := c.All("function", "where device_id=$1 order by id", "", device["id"])
	for _, item := range items {
		f := item["function"].(string)
		match := functionRegex.FindStringSubmatch(f)
		item["params"] = []any{}
		if len(match) == 2 {
			tokens := strings.Split(match[1], ",")
			item["params"] = tokens
		}
	}
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
func devicesPayloads(c *router.Context, guid string) {
	device := c.One("device", "where guid=$1", guid)
	c.TableJsonParams("payload", "where device_id=$1", device["id"])
}
func deviceSchedule(c *router.Context, guid string) {
	c.ReadJsonBodyIntoParams()
	device := c.One("device", "where guid=$1", guid)
	name := c.Params["name"].(string)
	if name == "Pet Feeder" {
		petJson := `{"photo": "https://i.imgur.com/M4UMA2Z.png"}`
		makePayload(c, device["id"], petJson)
	}
	c.SendContentAsJson("", 200)
}
func devicesByCat(c *router.Context, cat string) {
	send := map[string]any{}
	items := c.All("device", "where cat=$1 order by name", "", cat)
	send["items"] = items
	c.SendContentAsJson(send, 200)
}

func devicePatch(c *router.Context, id string) {
	c.ReadJsonBodyIntoParams()
	c.Update("device", "where guid=", id)
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
