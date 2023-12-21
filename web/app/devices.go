package app

import "github.com/andrewarrow/feedback/router"

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

func devicesFunctions(c *router.Context, guid string) {
	device := c.One("device", "where guid=$1", guid)
	c.TableJsonParams("function", "where device_id=$1", device["id"])
}
func devicesPayloads(c *router.Context, guid string) {
	device := c.One("device", "where guid=$1", guid)
	c.TableJsonParams("payload", "where device_id=$1", device["id"])
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

var fridgeJson = `{
  "fridgeId": "SFR123456",
  "model": "SmartFreeze 9000",
  "temperature": 2.5,
  "humidity": 45,
  "status": "operational",
  "lastUpdate": "2023-12-16T08:30:00Z",
  "inventory": [
    {
      "item": "Milk",
      "quantity": 2,
      "expiryDate": "2023-12-20"
    },
    {
      "item": "Eggs",
      "quantity": 1,
      "expiryDate": "2023-12-18"
    },
    {
      "item": "Yogurt",
      "quantity": 3,
      "expiryDate": "2023-12-22"
    },
    {
      "item": "Cheese",
      "quantity": 1,
      "expiryDate": "2023-12-25"
    },
    {
      "item": "Vegetables",
      "quantity": 5,
      "expiryDate": null
    }
  ],
  "energyConsumption": {
    "current": 120,
    "average": 100,
    "unit": "Watt-hours"
  },
  "alerts": [
    {
      "type": "temperature",
      "message": "Temperature is higher than usual. Check settings."
    },
    {
      "type": "inventory",
      "message": "Low quantity of eggs. Consider restocking."
    }
  ],
  "settings": {
    "temperatureUnit": "Celsius",
    "humidityAlertThreshold": 50
  }
}
`
