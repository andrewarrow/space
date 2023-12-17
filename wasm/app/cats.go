package app

import (
	"encoding/json"
	"space/wasm/network"
)

func queryForCategories() {
	jsonString := network.DoGet("/cats")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)

	div := Document.ByIdWrap("cats")
	for _, cat := range m["cats"].([]any) {
		catMap := map[string]any{}
		catMap["name"] = cat
		catMap["id"] = cat
		catMap["count"] = 0
		newHTML := Document.Render("cat", catMap)
		newDiv := Document.Document.Call("createElement", "div")
		newDiv.Set("innerHTML", newHTML)
		div.Call("appendChild", newDiv)
	}

	for _, input := range div.SelectAllByClass("cursor-pointer") {
		cat := NewCat(input.Id)
		input.Click(cat.Click)
	}
}
