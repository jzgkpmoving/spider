package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	const requestId = "requestId"
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()
		rid, _ := c.Get(requestId)
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elspad", time.Now().Sub(s)),
			zap.Int(requestId, rid.(int)))

	}, func(c *gin.Context) {
		c.Set(requestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(requestId); exists {
			h[requestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
