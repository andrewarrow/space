package app

import (
	"syscall/js"
)

type Device struct {
	Id string
}

func NewDevice(id string) *Device {
	d := Device{}
	d.Id = id
	return &d
}

func (d *Device) Click(this js.Value, params []js.Value) any {
	Document.RenderToId("modal-content", "device_show", nil)
	Document.ByIdWrap("modal").Show()
	return nil
}
