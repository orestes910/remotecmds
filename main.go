package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {

	Server := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	Server.GET("/run", func(c *gin.Context) {
		cmd := c.Query("cmd")
		valid := checkCmd(cmd)
		if valid {
			cmdOut, cmdErr := exec.Command(cmd).CombinedOutput()
			if cmdErr != nil {
				log.Fatal(cmdErr)
			}
			c.String(http.StatusOK, "%s", cmdOut)
		} else {
			c.String(http.StatusBadRequest, "%s", "Unknown Command, try...")
		}
	})

	Server.Run(":8080")

}

func checkCmd(cmd string) bool {
	cmds := [2]string{"date", "time"}
	for _, c := range cmds {
		if cmd == c {
			return true
		}
	}
	return false
}
