package app

import "github.com/andrewarrow/feedback/router"

func Login(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "POST" {
		handleLogin(c)
		return
	}
	c.NotFound = true
}

func handleLogin(c *router.Context) {
}
