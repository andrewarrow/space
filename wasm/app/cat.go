package app

import (
	"space/common"
	"syscall/js"
)

type Cat struct {
	Id string
}

func NewCat(id string) *Cat {
	d := Cat{}
	d.Id = id
	return &d
}

func (d *Cat) Click(this js.Value, params []js.Value) any {
	mc := Document.ById("modal-content")
	newHTML := Document.Render("cat_show", common.SmartHomeDeviceMaps)
	mc.Set("innerHTML", newHTML)
	Global.Stack = append(Global.Stack, newHTML)
	div := Document.ByIdWrap("devices")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		device := NewDevice(input.Id)
		input.Click(device.Click)
	}

	Document.ByIdWrap("modal").Show()
	return nil
}
