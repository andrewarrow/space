package app

import (
	"encoding/json"
	"space/wasm/network"
	"strings"
)

func queryForCategories() {
	jsonString := network.DoGet("/cats")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)

	countMap := m["cat_map"].(map[string]any)
	Document.ById("total").Set("innerHTML", m["total"])
	div := Document.ByIdWrap("cats")
	div.Set("innerHTML", "")
	cats := []string{}
	for _, cat := range m["cats"].([]any) {
		catMap := map[string]any{}
		catString := cat.(string)
		catMap["name"] = catString
		catMap["id"] = catString
		catMap["count"] = countMap[catString]
		newHTML := Document.Render("cat", catMap)
		current := div.Get("innerHTML")
		div.Set("innerHTML", current+newHTML)
		cats = append(cats, catString)
	}
	Global.Space["cats"] = strings.Join(cats, ",")

	for _, input := range div.SelectAllByClass("cursor-pointer") {
		cat := NewCat(input.Id)
		input.Click(cat.Click)
	}
}
