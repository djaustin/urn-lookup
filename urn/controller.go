package urn

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type URNController struct {
	Redis *redis.Client
}

func NewController(redisAddress string) *URNController {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	return &URNController{
		Redis: client,
	}
}

func (con *URNController) GetURN(c *gin.Context) {
	id := c.Param("id")
	cmd := con.Redis.HGetAll(c, id)
	if err := cmd.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := cmd.Val()
	if len(result) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("urn '%s' not found", id),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
