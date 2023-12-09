package app

import "github.com/andrewarrow/feedback/wasm"

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	if Global.Start == "welcome" {
		RegisterLoginEvents()
	} else if Global.Start == "space" {
		RegisterSpaceEvents()
	} else if Global.Start == "profile" {
		//RegisterProfileEvents()
	}
}
