package app

import (
	"fmt"
	"html/template"

	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {
	send := map[string]any{}
	c.LayoutMap["wasm"] = makeWasmScript("welcome")
	c.SendContentInLayout("welcome.html", send, 200)
}

func makeScript(s string) template.HTML {
	script := `<script>%s</script>`
	return template.HTML(fmt.Sprintf(script, s))
}

func makeWasmScript(s string) template.HTML {
	BuildTag := router.BuildTag
	t := fmt.Sprintf(wasmScript, BuildTag, s)
	return makeScript(t)
}

var wasmScript = `document.addEventListener("DOMContentLoaded", function() {
            const go = new Go();
  WebAssembly.instantiateStreaming(fetch("/assets/other/json.wasm.gz?id=%s"), go.importObject).then((result) => {
                go.run(result.instance);
                WasmReady('%s');
            });
});`
