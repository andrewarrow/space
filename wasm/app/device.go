package app

import (
	"encoding/json"
	"space/common"
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
	asString := common.FridgeJson
	var m map[string]any
	json.Unmarshal([]byte(asString), &m)
	Document.RenderToId("modal-content", "device_show", m)
	Document.ByIdWrap("modal").Show()
	return nil
}
