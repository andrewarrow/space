package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "POST" {
		handleLogin(c)
		return
	}
	if second == "" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	c.NotFound = true
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func handleLogin(c *router.Context) {
	c.ReadJsonBodyIntoParams()

	c.Params["username"] = c.Params["login-username"]
	c.Params["password"] = c.Params["login-password"]
	password := c.Params["password"].(string)

	row := c.One("user", "where username=$1", c.Params["username"])

	if len(row) > 0 && checkPasswordHash(password, row["password"].(string)) {
		guid := util.PseudoUuid()
		c.Params = map[string]any{"guid": guid, "user_id": row["id"]}
		c.Insert("cookie_token")
		router.SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
	} else {
		c.SendContentAsJson("username not found", 500)
		return
	}

	c.SendContentAsJson(row, 200)
}
