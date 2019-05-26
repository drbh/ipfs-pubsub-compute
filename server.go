package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"

	// "net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
)

var sh = shell.NewShell("localhost:5001")

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	// since were only running compute locally we
	// start the listener with a Go Routine.

	// could be and should be running on a seperate machine that can access
	// the pubsub channel
	go ListenForExecute()

	r.GET("/execute", func(c *gin.Context) {

		code := c.Query("code")
		event := c.Query("event")
		fmt.Println("Publishing")
		msg := `{   
		    "action": "execute",
		    "data": {
		    	"event": "` + event + `",
		    	"code": "` + code + `"
		    }
		}`

		resp := sh.PubSubPublish("test", msg)
		fmt.Println(resp)

		sub, _ := sh.PubSubSubscribe("test-response")
		r, _ := sub.Next()

		pubresp := strings.TrimSuffix(string(r.Data), "\n")

		c.JSON(200, gin.H{
			"message": pubresp,
		})
	})

	r.Run(":8769") // listen and serve on 0.0.0.0:8080
}
