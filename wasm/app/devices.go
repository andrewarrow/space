package app

import (
	"encoding/json"
	"space/wasm/network"
)

func queryForDevices(cat string) []any {
	jsonString := network.DoGet("/devices/" + cat)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items

	/*
		for _, input := range div.SelectAllByClass("cursor-pointer") {
			cat := NewCat(input.Id)
			input.Click(cat.Click)
		}*/
}
