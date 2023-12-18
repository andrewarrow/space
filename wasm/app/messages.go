package app

import "syscall/js"

func clickMessages(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	return nil
}
