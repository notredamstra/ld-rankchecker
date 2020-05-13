package main

import (
	"fmt"
	"ldjam-rank/pkg/ldmdare"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func getRank(c *gin.Context) {
	var game *ldmdare.LDGame
	var event *ldmdare.LDEvent
	var message string

	response := make(map[string]interface{})

	client := ldmdare.NewClient(baseURL, httpClient)
	userInput := c.PostForm("url")
	response["input"] = userInput

	userURL, err := url.ParseRequestURI(userInput)
	if err != nil {
		fmt.Println(err)
		client.BadRequest(c, response, "Invalid URL")
		return
	}
	game, err = client.GetGameFromURL(userURL)
	if err != nil {
		fmt.Println(err)
		client.BadRequest(c, response, "Could not get game url")
		return
	}
	response["game"] = game
	event, err = client.GetEventStatsFromGame(game)
	if err != nil {
		fmt.Println(err)
		client.BadRequest(c, response, "Could not get event stats from game")
		return
	}

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"input":   userInput,
		"game":    game,
		"event":   event,
		"message": message,
	})

}
