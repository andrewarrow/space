package app

import (
	"encoding/json"
	"fmt"
	"space/wasm/network"
	"syscall/js"
)

func clickSchedules(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	mc := Document.ById("modal-content")
	mc.Set("innerHTML", "")
	Document.ByIdWrap("modal").Show()

	go func() {
		items := queryForSchedules()
		newHTML := Document.Render("schedules", items)
		mc.Set("innerHTML", newHTML)
		Global.Submit("add-schedule-form", addSchedule)
		//SetScheduleClicks()
	}()
	return nil
}

func queryForSchedules() []any {
	jsonString := network.DoGet("/schedules")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	return items
}

func addSchedule(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	w := Document.ByIdWrapped("add-schedule-form")

	go func() {
		m := w.MapOfInputs()
		code := network.DoPost("/schedules", m)
		fmt.Println(code)
	}()
	return nil
}
