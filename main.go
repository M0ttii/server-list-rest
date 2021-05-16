package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func main() {
	r := gin.Default()
	RedisClient()
	r.GET("/user/:name", getUser)
	r.POST("/user/:name/:serverip", setUser)

	r.Run()

}

func RedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func setUser(c *gin.Context) {
	name := c.Param("name")
	serverip := c.Param("serverip")
	err := client.Set(ctx, name, serverip, 0).Err()
	if err != nil {
		panic(err)
	}

}
