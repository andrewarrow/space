package app

import "github.com/andrewarrow/feedback/router"

func Devices(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		devices(c)
		return
	}
	c.NotFound = true
}

func devices(c *router.Context) {
	send := map[string]any{}
	catMap := map[string]int{}
	items := c.All("device", "order by id", "")
	total := 0
	for _, item := range items {
		cat := item["cat"].(string)
		catMap[cat]++
		total++
	}
	send["cats"] = cats
	send["cat_map"] = catMap
	send["total"] = total
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
