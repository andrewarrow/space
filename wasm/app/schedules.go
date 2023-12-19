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

	go queryForSchedules()
	return nil
}

func queryForSchedules() {
	jsonString := network.DoGet("/schedules")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	newHTML := Document.Render("schedules", m["items"])
	mc := Document.ById("modal-content")
	mc.Set("innerHTML", newHTML)
	Global.Submit("add-schedule-form", addSchedule)
	links := Document.SelectAllFrom("schedule-list", "a")
	for _, w := range links {
		s := NewSchedule(w.Get("id"))
		w.Click(s.doDelete)
	}
}

func addSchedule(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	w := Document.ByIdWrapped("add-schedule-form")

	go func() {
		m := w.MapOfInputs()
		code := network.DoPost("/schedules", m)
		fmt.Println(code)
		queryForSchedules()
	}()
	return nil
}

type Schedule struct {
	Id string
}

func NewSchedule(id string) *Schedule {
	d := Schedule{}
	d.Id = id[1:]
	return &d
}

func (d *Schedule) doDelete(this js.Value, params []js.Value) any {
	Document.ByIdWrap("w" + d.Id).Hide()
	return nil
}
