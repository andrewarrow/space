package app

import (
	"encoding/json"
	"space/wasm/network"
)

func queryForDevice() {
	jsonString := network.DoGet("/cats")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)

	/*
		for _, input := range div.SelectAllByClass("cursor-pointer") {
			cat := NewCat(input.Id)
			input.Click(cat.Click)
		}*/
}
