package app

import (
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
		//SetScheduleClicks()
	}()
	return nil
}

func queryForSchedules() []any {
	return []any{"test"}
}

func createSchedule(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	w := Document.ByIdWrapped("add-schedule-form")

	go func() {
		m := w.MapOfInputs()
		code := network.DoPost("/schedules", m)
		fmt.Println(code)
	}()
	return nil
}
