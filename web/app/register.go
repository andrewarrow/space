package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func Register(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "POST" {
		handleRegister(c)
		return
	}
	c.NotFound = true
}

func handleRegister(c *router.Context) {
	c.ReadJsonBodyIntoParams()

	c.Params["username"] = c.Params["register-username"]
	c.Params["password"] = c.Params["register-password"]

	message := c.ValidateCreate("user")
	if message != "" {
		c.SendContentAsJson(message, 500)
		return
	}

	c.Params["password"] = router.HashPassword(c.Params["password"].(string))
	message = c.Insert("user")
	if message != "" {
		c.SendContentAsJson(message, 500)
		return
	}

	row := c.One("user", "where username=$1", c.Params["username"])
	guid := util.PseudoUuid()
	c.Params = map[string]any{"guid": guid, "user_id": row["id"]}
	c.Insert("cookie_token")
	router.SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
	c.SendContentAsJson(row, 200)
}
