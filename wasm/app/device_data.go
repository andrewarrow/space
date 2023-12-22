package app

import (
	"syscall/js"
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
	buttons := Document.ByIdWrap("device")
	all := buttons.SelectAll(".cursor-pointer")
	for _, input := range all {
		other := Document.ByIdWrap("h" + input.Id[1:])
		other.Hide()
	}
	h := Document.ByIdWrap("h" + d.Id)
	Global.Submit("h"+d.Id, apiInvoke)
	h.Show()
	return nil
}

func apiInvoke(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	Document.ByIdWrap("back").FireClick()
	return nil
}
