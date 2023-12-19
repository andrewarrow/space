package app

import (
	"space/wasm/network"
	"syscall/js"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

func RegisterLoginEvents() {
	Global.Click("login", login)
	Global.Click("register", register)
	Global.Submit("login-form", loginForm)
	Global.Submit("register-form", registerForm)
}

func login(this js.Value, p []js.Value) any {
	form := Document.ById("login-form")
	wasm.RemoveClass(form, "hidden")
	form = Document.ById("register-form")
	wasm.AddClass(form, "hidden")
	return nil
}

func register(this js.Value, p []js.Value) any {
	form := Document.ById("login-form")
	wasm.AddClass(form, "hidden")
	form = Document.ById("register-form")
	wasm.RemoveClass(form, "hidden")
	return nil
}

func loginForm(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	form := Document.ByIdWrapped("login-form")

	go func() {
		_, code := network.DoPost("/login", form.MapOfInputs())
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

func registerForm(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	form := Document.ByIdWrapped("register-form")

	go func() {
		_, code := network.DoPost("/register", form.MapOfInputs())
		if code == 200 {
			Global.Location.Set("href", "/space")
			return
		}
		flash := Document.ById("flash")
		flash.Set("innerHTML", "invalid username or password")
		time.Sleep(time.Second * 3)
		flash.Set("innerHTML", "")
	}()

	return false
}
