package app

import (
	"space/wasm/network"
	"syscall/js"
)

type CatPatch struct {
	DeviceId string
	Cat      string
}

func NewCatPatch(did, cat string) *CatPatch {
	d := CatPatch{}
	d.DeviceId = did
	d.Cat = cat[4:]
	return &d
}

func (d *CatPatch) Click(this js.Value, params []js.Value) any {

	Document.ByIdWrap("back").FireClick()
	Document.ByIdWrap("w" + d.DeviceId).Hide()
	go func() {
		payload := map[string]any{"cat": d.Cat}
		network.DoPatch("/devices/"+d.DeviceId, payload)
	}()

	return nil
}
