package main

import (
	"fmt"
	"net/http"

	"github.com/erparks/lology/pkg/riotclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	apiKey := initClientConfig()

	client := riotclient.New(apiKey, riotclient.NA1)

	r.GET("/summoner", func(c *gin.Context) {

		summonerName := c.Query("name")
		puuid, err := client.SummonerPUUID(summonerName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusOK, gin.H{
			"puuid": puuid,
		})
	})

	r.GET("/matches", func(c *gin.Context) {

		matches, err := client.MatchesByPUUID("oTvu-fQH1ES2SPqZVQbOTk1QP2uDttiP0jwDSwgOXOA76BZFWhTFFXk4vebo1FxYrKTX9_9VrnAJvQ")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusOK, gin.H{
			"matches": matches,
		})
	})

	r.Run()
}

func initClientConfig() string {
	viper.SetConfigName("client")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.Get("RIOT_API_KEY").(string)
}
