package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func RegisterSpaceEvents() {
	Global.Click("create", createTicket)
	Global.Click("x", clickX)
}

func createTicket(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	div := Document.ById("modal")
	wasm.RemoveClass(div, "hidden")
	return nil
}

func clickX(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	div := Document.ById("modal")
	wasm.AddClass(div, "hidden")
	return nil
}
