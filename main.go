package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

const adress = ""
const password = ""
const db = 0

func main() {
	r := gin.Default()
	RedisClient()
	r.GET("/user/:name", getUser)
	r.POST("/user/:name/:serverip", setUser)
	r.Run()
}

func RedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     adress,
		Password: password,
		DB:       db,
	})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	val, err := client.Get(ctx, name).Result()
	switch {
	case err == redis.Nil:
		c.String(http.StatusNotFound, "Key not found")
		return
	case err != nil:
		c.String(http.StatusBadRequest, "%s", err)
		return
	case val == "":
		c.String(http.StatusNoContent, "Empty value")
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": name, "server": val})
}

func setUser(c *gin.Context) {
	name := c.Param("name")
	serverip := c.Param("serverip")
	fmt.Printf("Pushing to Redis -> %s : %s", name, serverip)
	err := client.Set(ctx, name, serverip, 0).Err()
	if err != nil {
		panic(err)
	}

}
