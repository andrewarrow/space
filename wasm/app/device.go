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
	//asString := common.FridgeJson
	m := map[string]any{}
	//json.Unmarshal([]byte(asString), &m)
	if d.Id == "tesla" {
		m["location"] = "34.3,-118.34"
	} else if d.Id == "nest" {
		m["fahrenheit"] = 93
	} else if d.Id == "ring" {
		m["front_door"] = "https://some.com/image"
	} else if d.Id == "fridge" {
		m["eggs"] = 2
	} else if d.Id == "scale" {
		m["lbs"] = 0
	} else if d.Id == "smoke_detector" {
		m["alarm"] = false
	} else if d.Id == "smart_mirror" {
		m["background"] = "https://pleasant.com/image"
	} else if d.Id == "smart_lock" {
		m["back_door"] = "locked"
		m["front_door"] = "unlocked"
		m["east_side_door"] = "ajar"
	} else if d.Id == "smart_coffee" {
		m["brew_at"] = 1702758928
	} else if d.Id == "pet_feeder" {
		m["meals_every"] = 86400
	}

	Document.RenderToId("modal-content", "device_show", m)
	Document.ByIdWrap("modal").Show()
	return nil
}
