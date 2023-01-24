package main

import (
	"net/http"

	"github.com/erparks/lology/pkg/riotclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	client := riotclient.New("RGAPI-afdf1a00-6a90-4041-90ae-194933967301", riotclient.NA1)

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
