package app

var smartHomeDevices = []string{
	"Smart Fridge",
	"Smart Scale",
	"Nest Thermostat",
	"Amazon Alexa",
	"Smart Lighting",
	"Ring Doorbell",
	"Tesla",
	"iPhone",
	"Apple Watch",
}

var smartHomeDeviceMaps = []map[string]any{{"id": 1, "name": smartHomeDevices[0]},
	{"id": 2, "name": smartHomeDevices[1]},
	{"id": 3, "name": smartHomeDevices[2]},
	{"id": 4, "name": smartHomeDevices[3]},
	{"id": 5, "name": smartHomeDevices[4]},
	{"id": 6, "name": smartHomeDevices[5]},
	{"id": 7, "name": smartHomeDevices[6]},
	{"id": 8, "name": smartHomeDevices[7]},
	{"id": 9, "name": smartHomeDevices[8]}}

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
