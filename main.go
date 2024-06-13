package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {
	// set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they dont have to be loaded
	// from the disk again. This makes serving HTML pages very fast
	router.LoadHTMLGlob("./templates/*")

	// (1) Define the route for the index page and display the index.html template
	// (1) To start with, using an inline route handler. Later on, will create
	// (1) standalone functions that will be used as router handler.
	// (2) Handle Index
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	// Start serving the application
	router.Run()
}
