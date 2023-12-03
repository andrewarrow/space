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
	c.LayoutMap["wasm"] = makeScript(wasmScript)
	c.SendContentInLayout("welcome.html", send, 200)
}

func makeScript(s string) template.HTML {
	script := `<script>%s</script>`
	return template.HTML(fmt.Sprintf(script, s))
}

var wasmScript = `document.addEventListener("DOMContentLoaded", function() {
            const go = new Go();
  WebAssembly.instantiateStreaming(fetch("/assets/other/json.wasm.gz?ran={{uuid}}"), go.importObject).then((result) => {
                go.run(result.instance);
                WasmReady('test');
            });
});`
