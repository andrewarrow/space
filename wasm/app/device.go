package app

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Device struct {
	Guid any
	Name any
}

func NewDevice(device map[string]any) *Device {
	d := Device{}
	d.Guid = device["guid"]
	d.Name = device["name"]
	return &d
}

/*
func SetDeviceClicks() {
	div := Document.ByIdWrap("devices")
	for _, input := range div.SelectAll(".cat-click") {
		name := Document.ByIdWrap("n" + input.Id[1:])
		device := NewDevice(input.Id, name.Get("innerHTML"))
		input.Click(device.Click)
	}
	for _, input := range div.SelectAll(".data-click") {
		name := Document.ByIdWrap("n" + input.Id[1:])
		device := NewDataDevice(input.Id, name.Get("innerHTML"))
		input.Click(device.Click)
	}
}*/

func (d *Device) Click(this js.Value, params []js.Value) any {
	mc := Document.ByIdWrap("modal-content")
	mc.Set("innerHTML", "")
	Document.ById("title").Set("innerHTML", d.Name)

	go func() {
		items := queryForDeviceFunctions(d.Guid)
		mc.Set("innerHTML", wasm.Render("device_show", items))
	}()
	modal := Document.ByIdWrap("modal")
	modal.Show()
	return nil
}
