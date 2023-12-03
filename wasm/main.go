package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	wasm.NewGlobal()

	select {}
}
