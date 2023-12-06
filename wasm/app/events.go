package app

import (
	"space/wasm/network"
	"syscall/js"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	Global.Click("login", login)
	Global.Submit("login-form", loginForm)
}

func login(this js.Value, p []js.Value) any {
	Global.Location.Set("href", "/space")
	return nil
}

func loginForm(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	form := Document.ByIdWrapped("login-form")
	m := map[string]any{}
	for _, input := range form.SelectAll("input") {
		m[input.Id] = input.Value
	}
	go func() {
		code := network.DoPost("/login", m)
		if code == 200 {
			Global.Location.Set("href", "/space")
			return
		}
		flash := Document.ById("flash")
		flash.Set("innerHTML", "invalid login")
		time.Sleep(time.Second * 3)
		flash.Set("innerHTML", "")
	}()
	return false
}
