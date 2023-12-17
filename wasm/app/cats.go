package app

import (
	"encoding/json"
	"space/wasm/network"
)

func queryForCategories() {
	jsonString := network.DoGet("/cats")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)

	countMap := m["cat_map"].(map[string]any)
	//total := m["total"]
	div := Document.ByIdWrap("cats")
	for _, cat := range m["cats"].([]any) {
		catMap := map[string]any{}
		catMap["name"] = cat
		catMap["id"] = cat
		catMap["count"] = countMap[cat.(string)]
		newHTML := Document.Render("cat", catMap)
		current := div.Get("innerHTML")
		div.Set("innerHTML", current+newHTML)
	}

	for _, input := range div.SelectAllByClass("cursor-pointer") {
		cat := NewCat(input.Id)
		input.Click(cat.Click)
	}
}
