package main

import (
	"fmt"
	"math/rand"
	"space/wasm/app"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var g *wasm.Global

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	g = wasm.NewGlobal()
	app.RegisterEvents(g)

	select {}
}
