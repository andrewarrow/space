package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type DataDevice struct {
	Id   string
	Name string
}

func NewDataDevice(id, name string) *DataDevice {
	d := DataDevice{}
	d.Id = id[1:]
	d.Name = name
	return &d
}

func (d *DataDevice) Click(this js.Value, params []js.Value) any {
	mc := Document.ByIdWrap("modal-content")
	si := wasm.NewStackItem(mc.Get("innerHTML"))
	mc.Set("innerHTML", "")
	si.Callback = SetDeviceClicks
	Global.Stack = append(Global.Stack, si)

	send := map[string]any{}
	send["item"] = d.Name
	newHTML := Document.Render("device_data_show", send)
	mc.Set("innerHTML", newHTML)
	return nil
}
