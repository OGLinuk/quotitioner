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

func getQuote() string {
	f, err := os.Open("quotes.txt")
	if err != nil {
		log.Printf("main.go::getQuote::os.Open::ERROR: %s", err.Error())
	}
	defer f.Close()

	var quotes []string
	bs := bufio.NewScanner(f)

	for bs.Scan() {
		quotes = append(quotes, bs.Text())
	}

	return quotes[rand.Intn(len(quotes))]
}

func indexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"quote": getQuote(),
	})
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
