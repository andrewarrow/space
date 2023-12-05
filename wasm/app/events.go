package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global

func RegisterEvents() {
	Global.Click("login", login)
	Global.Submit("login-form", loginForm)
}

func login(this js.Value, p []js.Value) any {
	Global.Location.Set("href", "/space")
	return nil
}

func loginForm(this js.Value, p []js.Value) any {
	this.Call("preventDefault")
	//Global.Location.Set("href", "/space")
	return nil
}
