package app

import "syscall/js"

func clickSchedules(this js.Value, params []js.Value) any {
	params[0].Call("preventDefault")
	return nil
}
