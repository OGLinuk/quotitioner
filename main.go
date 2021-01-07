package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// getQuote loads quotes and shuffles the list of quotes,
// then returns a pseudo-random quote from the list.
func getQuote() (string, error) {
	// TODO: quotes.txt will eventually need to be compressed or hosted
	// due to its size. Currently 7.5K (66 quotes).
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

	rand.Shuffle(len(quotes), func(i, j int) {
		quotes[i], quotes[j] = quotes[j], quotes[i]
	})

	return quotes[rand.Intn(len(quotes))], nil
}

// restHandler calls and returns getQuote as JSON
func restHandler(ctx *gin.Context) {
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

// indexHandler calls and returns getQuote as a raw string
func indexHandler(ctx *gin.Context) {
	quote, err := getQuote()
	if err != nil {
		// TODO: ...
		// http error
	} else {
		fmt.Fprintf(ctx.Writer, quote)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9429"
	}

	g := gin.Default()

	g.GET("/", indexHandler)
	g.GET("/rest", restHandler)

	if err := g.Run(fmt.Sprintf("0.0.0.0:%s", PORT)); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
