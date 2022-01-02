package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

type user struct {
	Name    string
	Message string
	Number  int
}

func main() {

	r := gin.Default()

	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"type": "JSON", "status": 200})
	})
	r.GET("/someStruct", func(c *gin.Context) {

		msg := user{
			Name:    "ok",
			Message: "sources",
			Number:  http.StatusOK,
		}
		fmt.Println(msg)
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/somexml", func(c *gin.Context) {
		c.XML(http.StatusOK, "message:XML")
	})
	r.GET("/someYaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, "message:yaml")
	})

	r.GET("someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		Label := "Liable"

		data := &protoexample.Test{
			Reps:  reps,
			Label: &Label,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "模板渲染"})
	})

r.GET("/someredirect", func(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
})
	r.GET("/somesync", func(c *gin.Context) {
		copySync:=c.Copy()
		time.Sleep(time.Second*3)
		log.Println("异步执行"+copySync.Request.URL.Path)
		c.JSON(http.StatusOK,gin.H{"data":copySync.Request.URL.Path})
	})
	r.GET("/moreasync", func(c *gin.Context) {
		log.Println("同步url"+c.Request.URL.Path)
	})
	err := r.Run("localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
