package app

import (
	"encoding/json"
	"fmt"
	"space/wasm/network"
)

func queryForDevices() []any {
	jsonString := network.DoGet("/devices")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}
func queryForDeviceFunctions(guid any) []any {
	jsonString := network.DoGet(fmt.Sprintf("/devices/functions/%s", guid))
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}
func queryForDevicePayloads(guid any) []any {
	jsonString := network.DoGet(fmt.Sprintf("/devices/payloads/%s", guid))
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}
