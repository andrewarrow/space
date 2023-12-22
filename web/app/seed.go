package app

import "github.com/andrewarrow/feedback/router"

func Seed(c *router.Context) {
	id := makeDevice(c, "Living Room TV")
	id = makeDevice(c, "Bedroom TV")
	id = makeDevice(c, "Pet Feeder")
	makeFunction(c, id, "release_food(bay,grams)")
	makePayload(c, id, PET_JSON)
}

func makeDevice(c *router.Context, s string) any {
	c.Params = map[string]any{}

	c.Params["name"] = s
	c.ValidateCreate("device")
	c.Insert("device")
	guid := c.Params["guid"]
	one := c.One("device", "where guid=$1", guid)
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
func makePayload(c *router.Context, id any, s string) any {
	c.Params = map[string]any{}

	c.Params["payload"] = s
	c.Params["device_id"] = id
	c.ValidateCreate("payload")
	c.Insert("payload")
	guid := c.Params["guid"]
	return guid
}

const PET_JSON = `{"bays": [ {"id": "BAY_01", "sku": "MEOW_MIX_1010", "flavor": "Chicken and Yumminess", "grams": 2000}, {"id": "BAY_02", "sku": "MEOW_MIX_1020", "flavor": "Fresh Sardines", "grams": 1900} ] }`
