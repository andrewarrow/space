package app

import (
	"strings"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Device struct {
	Id string
}

func NewDevice(id string) *Device {
	d := Device{}
	d.Id = id
	return &d
}

func SetDeviceClicks() {
	div := Document.ByIdWrap("devices")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		device := NewDevice(input.Id)
		input.Click(device.Click)
	}
}

func (d *Device) Click(this js.Value, params []js.Value) any {
	mc := Document.ByIdWrap("modal-content")
	si := wasm.NewStackItem(mc.Get("innerHTML"))
	mc.Set("innerHTML", "")
	si.Callback = SetDeviceClicks
	Global.Stack = append(Global.Stack, si)

	send := map[string]any{}
	send["item"] = d.Id
	tokens := strings.Split(Global.Space["cats"], ",")
	send["cats"] = tokens
	newHTML := Document.Render("device_show", send)
	mc.Set("innerHTML", newHTML)
	div := Document.ByIdWrap("device-cats")
	for _, input := range div.SelectAllByClass("cursor-pointer") {
		cat := NewCatPatch(d.Id, input.Id)
		input.Click(cat.Click)
	}
	return nil
}

func (d *Device) OldClick(this js.Value, params []js.Value) any {
	//asString := common.FridgeJson
	send := map[string]any{"name": d.Id}
	m := map[string]any{}
	//json.Unmarshal([]byte(asString), &m)
	if d.Id == "tesla" {
		v := map[string]any{}
		v["location"] = "34.3,-118.34"
		m["values"] = v
		m["commands"] = []string{"self_drive_to"}
	} else if d.Id == "nest" {
		v := map[string]any{}
		v["fahrenheit"] = 93
		m["values"] = v
		m["commands"] = []string{"set_temperature"}
	} else if d.Id == "ring" {
		v := map[string]any{}
		v["front_door"] = "https://i.imgur.com/GX6D5hj.png"
		m["values"] = v
		m["commands"] = []string{"take_photo", "record_audio"}
	} else if d.Id == "fridge" {
		v := map[string]any{}
		v["eggs"] = 2
		v["milk_gallons"] = 1.5
		m["values"] = v
		m["commands"] = []string{"set_temperature", "list_items"}
	} else if d.Id == "scale" {
		v := map[string]any{}
		m["bob_lbs"] = 190
		m["sue_lbs"] = 130
		m["values"] = v
		m["commands"] = []string{"list_profiles", "add_profile"}
	} else if d.Id == "smoke_detector" {
		v := map[string]any{}
		m["alarm"] = false
		m["values"] = v
		m["commands"] = []string{"sound_alarm"}
	} else if d.Id == "smart_mirror" {
		v := map[string]any{}
		v["background"] = "https://i.imgur.com/fg3virE.png"
		m["values"] = v
		m["commands"] = []string{"set_text"}
	} else if d.Id == "smart_lock" {
		v := map[string]any{}
		v["back_door"] = "locked"
		v["front_door"] = "unlocked"
		v["east_side_door"] = "ajar"
		m["values"] = v
		m["commands"] = []string{"panic_mode", "list_locks"}
	} else if d.Id == "smart_coffee" {
		val := map[string]any{}
		val["brew_at"] = 1702758928
		val["cups"] = 2
		val["flavor"] = "house_blend"
		val["intensity"] = 11
		m["values"] = val
		m["commands"] = []string{"list_flavors", "brew_now"}
	} else if d.Id == "pet_feeder" {
		val := map[string]any{}
		val["meals_every"] = 86400
		val["grams"] = 90
		val["flavor"] = "meow_mix"
		val["photo"] = "https://i.imgur.com/M4UMA2Z.png"
		m["values"] = val
		m["commands"] = []string{"list_flavors", "feed_now", "take_photo"}
	}
	send["data"] = m

	mc := Document.ByIdWrap("modal-content")
	si := wasm.NewStackItem(mc.Get("innerHTML"))
	si.Callback = SetDeviceClicks
	Global.Stack = append(Global.Stack, si)

	newHTML := Document.Render("device_show", send)
	mc.Set("innerHTML", newHTML)

	Document.ByIdWrap("modal").Show()
	return nil
}
