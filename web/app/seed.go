package app

import "github.com/andrewarrow/feedback/router"

func Seed(c *router.Context) {
	id := makeDevice(c, "Living Room TV")
	id = makeDevie(c, "Bedroom TV")
	id = makeDevie(c, "Pet Feeder")
	makeFunction(c, id, "release_food(bay,grams)")
}

func makeDevice(c *router.Context, s string) any {
	c.Params = map[string]any{}

	c.Params["name"] = s
	c.ValidateCreate("device")
	c.Insert("device")
	guid := c.Params["guid"]
	one = c.One("device", "where guid=$1", guid)
	return one["id"]
}
func makeFunction(c *router.Context, id any, s string) any {
	c.Params = map[string]any{}

	c.Params["function"] = s
	c.Params["device_id"] = id
	c.ValidateCreate("function")
	c.Insert("function")
	guid := c.Params["guid"]
	return guid
}
