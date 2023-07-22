// handlers/get_handler.go
package handlers

import (
	"Go/src/go/go-sensibull/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("redis connection init..")
	utils.InitRedis("localhost:6379", "", 0)
}
func GetHandler(c *gin.Context) {
	// Check if the data is present in the cache
	cachedData, err := utils.GetFromCache("my_cache_key")
	if err == nil && cachedData != "" {
		c.JSON(http.StatusOK, gin.H{"data": cachedData})
		return
	}

	// If data is not in cache, fetch it from the database or another data source
	// For now, let's just return some example data
	data := "Hello from GET API!"

	// Save the fetched data to the cache
	utils.SetToCache("my_cache_key", data, 30*time.Second)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
