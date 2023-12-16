package common

var FridgeJson = `{
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
  "alerts3": [
    {
      "type": "temperature",
      "message": "Temperature is higher than usual. Check settings."
    },
    {
      "type": "inventory",
      "message": "Low quantity of eggs. Consider restocking."
    }
  ],
  "alerts2": [
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
