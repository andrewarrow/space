package app

import (
	"space/wasm/network"
	"syscall/js"
)

type DataDevice struct {
	Id   string
	Name any
}

func NewDataDevice(id string, name any) *DataDevice {
	d := DataDevice{}
	d.Id = id[1:]
	d.Name = name
	return &d
}

func (d *DataDevice) Click(this js.Value, params []js.Value) any {
	buttons := Document.ByIdWrap("device")
	all := buttons.SelectAll(".cursor-pointer")
	for _, input := range all {
		other := Document.ByIdWrap("h" + input.Id[1:])
		other.Hide()
	}
	h := Document.ByIdWrap("h" + d.Id)
	Global.Submit("h"+d.Id, d.apiInvoke)
	h.Show()
	return nil
}

func (d *DataDevice) apiInvoke(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	Document.ByIdWrap("back").FireClick()

	payload = map[string]any{"name": d.Name, "when": "now"}
	go network.DoPost("/devices/schedule/"+d.Id, payload)
	return nil
}
