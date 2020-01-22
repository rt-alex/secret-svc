package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type kv struct {
	Key string `json:'key'`
	Val string `json:'val'`
}

func main() {
	log.SetOutput(os.Stdout)
	r := gin.Default()
	r.GET("/env", getENV)
	r.GET("/secret", getSecret)
	r.Run(":80")
}

func getSecret(c *gin.Context) {
	v := os.Getenv("SECRET")
	data := kv{}
	if len(v) == 0 {
		data = kv{
			Key: "SECRET",
			Val: "UNSET",
		}
	} else {
		data = kv{
			Key: "SECRET",
			Val: v,
		}
	}
	c.JSON(http.StatusOK, data)
}

func getENV(c *gin.Context) {
	v := os.Getenv("ENV")
	data := kv{}
	if len(v) == 0 {
		data = kv{
			Key: "ENV",
			Val: "UNSET",
		}
	} else {
		data = kv{
			Key: "ENV",
			Val: v,
		}
	}
	fmt.Println(data)
	c.JSON(http.StatusOK, data)
}
