package app

import (
	"fmt"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

func RegisterEvents(g *wasm.Global) {
	g.Click("login", login)
}

func login(this js.Value, p []js.Value) any {
	fmt.Println("login")
	return nil
}
