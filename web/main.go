package main

import (
	"embed"
	"math/rand"
	"os"
	"space/web/app"
	"time"

	"github.com/andrewarrow/feedback/common"
	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}
	arg := os.Args[1]

	if arg == "render" {
		router.RenderMarkup()
	} else if arg == "seed" {
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		app.Seed(r.ToContext())
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		tf := common.TemplateFunctions()
		router.CustomFuncMap = &tf
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.HandleWelcome
		r.Paths["space"] = app.HandleSpace
		r.Paths["login"] = app.Login
		r.Paths["register"] = app.Register
		r.Paths["tickets"] = app.Ticket
		r.Paths["schedules"] = app.Schedules
		r.Paths["devices"] = app.Devices
		r.Paths["cats"] = app.Cats
		r.Paths["markup"] = app.Markup
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}

}
