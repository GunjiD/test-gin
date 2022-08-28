package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	bashExecute := false
	cmd := exec.Command("bash")
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":  "hello gin",
			"hogefuga": "hogefuga",
		})
	})

	r.GET("/bash_start", func(c *gin.Context) {
		if bashExecute {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"bash_state": "bash が実行中です",
			})
			return
		}

		cmd.Start()
		bashExecute = true
		c.HTML(http.StatusOK, "index.html", gin.H{
			"bash_state": "bash が実行されました",
		})
	})

	r.GET("/bash_stop", func(c *gin.Context) {
		if !bashExecute {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"bash_state": "bash は停止中です",
			})
			return
		}

		cmd.Process.Kill()
		cmd.Wait()
		bashExecute = false
		c.HTML(http.StatusOK, "index.html", gin.H{
			"bash_state": "bash が停止されました",
		})
	})

	r.Run()
}
