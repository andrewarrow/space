package app

import (
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
	mc.Set("innerHTML", "")
	Document.ByIdWrap("modal").Show()

	go func() {
		items := queryForDevices()
		newHTML := Document.Render("cat_show", items)
		mc.Set("innerHTML", newHTML)
		//SetDeviceClicks()
	}()
	return nil
}
