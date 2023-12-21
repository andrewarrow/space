package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type DataDevice struct {
	Id string
}

func NewDataDevice(id string) *DataDevice {
	d := DataDevice{}
	d.Id = id[1:]
	return &d
}

func (d *DataDevice) Click(this js.Value, params []js.Value) any {
	mc := Document.ByIdWrap("d" + d.Id)
	current := mc.Get("innerHTML")
	mc.Set("innerHTML", current+"<p>hi</p>")

	return nil
}

func apiInvoke(this js.Value, params []js.Value) any {
	mc := Document.ByIdWrap("modal-content")
	si := wasm.NewStackItem(mc.Get("innerHTML"))
	mc.Set("innerHTML", "")
	si.Callback = func() { Global.Click("api1", apiInvoke) }
	Global.Stack = append(Global.Stack, si)

	send := map[string]any{}
	send["item"] = "api1"
	newHTML := Document.Render("device_data_invoke", send)
	mc.Set("innerHTML", newHTML)
	return nil
}
