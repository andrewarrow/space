package app

import (
	"fmt"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global

func RegisterEvents() {
	Global.Click("login", login)
}

func login(this js.Value, p []js.Value) any {
	fmt.Println("login")
	return nil
}
