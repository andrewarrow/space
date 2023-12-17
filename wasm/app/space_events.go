package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func RegisterSpaceEvents() {
	div := Document.ByIdWrap("cats")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		cat := NewCat(input.Id)
		input.Click(cat.Click)
	}
	Global.Click("back", clickBack)
}

func showCreateTicket(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	div := Document.ById("modal")
	wasm.RemoveClass(div, "hidden")
	Document.ById("title").Set("value", "")
	Document.ById("title").Call("focus")
	Document.ById("description").Set("value", "")
	Global.Submit("ticket-form", createTicket)
	return nil
}

func clickBack(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	if len(Global.Stack) == 0 {
		div := Document.ById("modal")
		wasm.AddClass(div, "hidden")
	} else {
		si := Global.Stack[len(Global.Stack)-1]
		mc := Document.ById("modal-content")
		mc.Set("innerHTML", si.HTML)
		si.Callback()
		Global.Stack = Global.Stack[0 : len(Global.Stack)-1]
	}
	return nil
}
