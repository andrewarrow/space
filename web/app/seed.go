package app

import (
	"encoding/json"

	"github.com/andrewarrow/feedback/router"
)

func Seed(c *router.Context) {
	id := makeDevice(c, "Living Room TV")
	id = makeDevice(c, "Bedroom TV")
	id = makeDevice(c, "Pet Feeder")
	makeFunction(c, id, "release_food(bay,grams)")
	makeFunction(c, id, "play_audio(file)")
	makeFunction(c, id, "take_photo")
	makePayload(c, id, PET_JSON)

	id = makeDevice(c, "Fridge")
	makeFunction(c, id, "set_temperature(celsius)")
	makeFunction(c, id, "set_temperature(fahrenheit)")
	makeFunction(c, id, "play_audio(file)")
	makeFunction(c, id, "take_photo")
	makePayload(c, id, FRIDGE)
	id = makeDevice(c, "Andrew's iPhone")
	makePayload(c, id, IPHONE)
	id = makeDevice(c, "Tesla 2017 Red")
	id = makeDevice(c, "Tesla 2019 White")
	id = makeDevice(c, "iBike e-Assist")
	id = makeDevice(c, "Nest")
	id = makeDevice(c, "Ring")
	id = makeDevice(c, "Bathroom Scale")
	id = makeDevice(c, "Bathroom Toothbrush")
	id = makeDevice(c, "Bathroom Mirror")
	id = makeDevice(c, "Coffee")
	id = makeDevice(c, "Jen iPhone")
	id = makeDevice(c, "Jordan iPhone")
	id = makeDevice(c, "Jason iPhone")
	id = makeDevice(c, "Athena iPad")
	id = makeDevice(c, "Austin iPad")
	id = makeDevice(c, "Playstation")
	id = makeDevice(c, "X-Box")
	id = makeDevice(c, "Alexa")
	id = makeDevice(c, "Small Speaker")
	id = makeDevice(c, "2022 MacBook Air")
	id = makeDevice(c, "2020 MacBook Pro")
	id = makeDevice(c, "2021 iMac")
	id = makeDevice(c, "Rouge Bad Guy Device")
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

	var m map[string]any
	json.Unmarshal([]byte(s), &m)
	c.Params["payload"] = m
	c.Params["device_id"] = id
	c.ValidateCreate("payload")
	c.Insert("payload")
	guid := c.Params["guid"]
	return guid
}

const FRIDGE = `{ "fridgeId": "SFR123456", "model": "SmartFreeze 9000", "temperature": 2.5, "humidity": 45, "status": "operational", "lastUpdate": "2023-12-16T08:30:00Z", "inventory": [ { "item": "Milk", "quantity": 2, "expiryDate": "2023-12-20" }, { "item": "Eggs", "quantity": 1, "expiryDate": "2023-12-18" }, { "item": "Yogurt", "quantity": 3, "expiryDate": "2023-12-22" }, { "item": "Cheese", "quantity": 1, "expiryDate": "2023-12-25" }, { "item": "Vegetables", "quantity": 5, "expiryDate": null } ], "energyConsumption": { "current": 120, "average": 100, "unit": "Watt-hours" }, "alerts": [ { "type": "temperature", "message": "Temperature is higher than usual. Check settings." }, { "type": "inventory", "message": "Low quantity of eggs. Consider restocking." } ], "settings": { "temperatureUnit": "Celsius", "humidityAlertThreshold": 50 } }`

const PET_JSON = `{"bays": [ {"id": "BAY_01", "sku": "MEOW_MIX_1010", "flavor": "Chicken and Yumminess", "grams": 2000}, {"id": "BAY_02", "sku": "MEOW_MIX_1020", "flavor": "Fresh Sardines", "grams": 1900} ] }`

const IPHONE = `{ "device_info": { "device_name": "Andrew iPhone", "model": "iPhone SE", "os_version": "iOS 15.3", "serial_number": "1234567890", "battery": { "battery_percentage": 75, "battery_status": "Charging", "charging_source": "Lightning Cable" }, "storage": { "total_storage": "128 GB", "available_storage": "64 GB", "used_storage": "64 GB" } }, "system_status": { "cpu_usage": "25%", "memory": { "total_memory": "4 GB", "used_memory": "2 GB", "free_memory": "2 GB" }, "network": { "wifi": { "status": "Connected", "signal_strength": "Excellent" }, "cellular": { "status": "4G", "signal_strength": "High" } } }, "apps": { "running_apps": [ { "name": "Safari", "status": "Active", "cpu_usage": "10%", "memory_usage": "150 MB" }, { "name": "Messages", "status": "Inactive", "cpu_usage": "5%", "memory_usage": "80 MB" }, { "name": "Music", "status": "Active", "cpu_usage": "8%", "memory_usage": "120 MB" } ] }, "security": { "face_id": { "enabled": false }, "touch_id": { "enabled": true }, "passcode": { "enabled": true, "complexity": "Medium" } } }`
