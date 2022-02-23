package urn

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// URNController encapsulates structs and logic required for fulfilling URN requests
type URNController struct {
	Redis *redis.Client
}

// NewController returns a pointer to a URNController with a Redis client
// for the provided address attached
func NewController(redisAddress string) *URNController {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	return &URNController{
		Redis: client,
	}
}

// GetURN retrieves all hash data associated with the :id param.
// The hash data is returned to the client in a JSON payload.
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
