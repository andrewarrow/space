package main

import (
	"embed"
	"fmt"
	"math/rand"
	"space/wasm/app"
	"space/wasm/network"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

//go:embed views/*.html
var embeddedTemplates embed.FS

var useLive string
var viewList string

func main() {
	fmt.Println(viewList)
	wasm.UseLive = useLive == "true"
	wasm.EmbeddedTemplates = embeddedTemplates
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	app.Global, app.Document = wasm.NewGlobal()

	<-app.Global.Ready
	if wasm.UseLive {
		go wasm.LoadAllTemplates(viewList, network.DoGet)
	}

	app.RegisterEvents()

	select {}
}
