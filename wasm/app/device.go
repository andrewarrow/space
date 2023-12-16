package app

import (
	"fmt"
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
	fmt.Println(d.Id)
	return nil
}
