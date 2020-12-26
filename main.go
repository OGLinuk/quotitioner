package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getQuote() (string, error) {
	f, err := os.Open("quotes.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var quotes []string
	bs := bufio.NewScanner(f)

	for bs.Scan() {
		quotes = append(quotes, bs.Text())
	}

	return quotes[rand.Intn(len(quotes))], nil
}

func indexHandler(ctx *gin.Context) {
	quote, err := getQuote()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"quote": quote,
		})
	}
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	g := gin.Default()

	g.GET("/", indexHandler)

	if err := g.Run(fmt.Sprintf("0.0.0.0:%s", PORT)); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
