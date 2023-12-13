package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func createTicket(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	div := Document.ById("modal")
	wasm.AddClass(div, "hidden")
	return nil
}
