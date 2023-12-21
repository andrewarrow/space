package app

import (
	"encoding/json"
	"space/wasm/network"
)

func queryForDevices() []any {
	jsonString := network.DoGet("/devices")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}
func queryForDeviceFunctions(guid string) []any {
	jsonString := network.DoGet("/devices/functions/" + guid)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}
