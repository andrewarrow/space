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
	SetDeviceClicks()

	Document.ByIdWrap("modal").Show()
	return nil
}
