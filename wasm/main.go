package main

import (
	"fmt"
	"math/rand"
	"space/wasm/app"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	app.Global, app.Document = wasm.NewGlobal()
	app.RegisterEvents()

	select {}
}
