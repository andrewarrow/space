package app

import (
	"fmt"
	"space/wasm/network"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func createTicket(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	div := Document.ById("modal")
	wasm.AddClass(div, "hidden")
	w := Document.ByIdWrapped("ticket-form")

	go func() {
		m := w.MapOfInputs()
		code := network.DoPost("/ticket", m)
		fmt.Println(code)
	}()
	return nil
}
