package app

import (
	"fmt"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func RegisterSpaceEvents() {
	//div := Document.ById("modal")
	//wasm.RemoveClass(div, "hidden")
	//apiInvoke(Document.Document, []js.Value{})
	go handleQueryForDevices()
	Global.Click("back", clickBack)
	Global.Click("schedules", clickSchedules)
	Global.Click("messages", clickMessages)
}

func handleQueryForDevices() {
	items := queryForDevices()
	div := Document.ByIdWrap("devices")
	div.Set("innerHTML", fmt.Sprintf("%s", items))
}

func clickBack(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	if len(Global.Stack) == 0 {
		go queryForCategories()
		div := Document.ById("modal")
		wasm.AddClass(div, "hidden")
	} else {
		si := Global.Stack[len(Global.Stack)-1]
		mc := Document.ById("modal-content")
		mc.Set("innerHTML", si.HTML)
		if si.Callback != nil {
			si.Callback()
		}
		Global.Stack = Global.Stack[0 : len(Global.Stack)-1]
	}
	return nil
}
