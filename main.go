package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	ID       int    `json:"id"`
	Response string `json:"response"`
}

var responses = []response{
	{ID: 0, Response: "Тодорхой, тийм"},
	{ID: 1, Response: "Тийм байдаг шүү"},
	{ID: 2, Response: "Эргэлзээгүй тийм"},
	{ID: 3, Response: "Тийм, Тийм"},
	{ID: 4, Response: "Найдаж болно"},
	{ID: 5, Response: "Би тийм гэж харж байна"},
	{ID: 6, Response: "Тийм байх өндөр магадлалтай"},
	{ID: 7, Response: "Дажгүй харагдаж байна"},
	{ID: 8, Response: "Тийм"},
	{ID: 9, Response: "Зөв чиглэл рүү зааж байна"},
	{ID: 10, Response: "Тодорхойгүй байна, дахиад үзнэ үү"},
	{ID: 11, Response: "Дараа дахин үзнэ үү"},
	{ID: 12, Response: "Одоо хариу хэлэхгүй байх нь зөв"},
	{ID: 13, Response: "Зөгнөх боломжгүй байна"},
	{ID: 14, Response: "Төвлөр, тэгээд ахиад асуу"},
	{ID: 15, Response: "Найдах хэрэггүй"},
	{ID: 16, Response: "Миний хариулт бол үгүй"},
	{ID: 17, Response: "Үгүй юм байна"},
	{ID: 18, Response: "Муухан харагдаж байна"},
	{ID: 19, Response: "Маш эргэлзээтэй"},
}

func getAllAnswers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, responses)
}

func getRandomAnswer(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	c.IndentedJSON(http.StatusOK, responses[rand.Intn(len(responses))])
}

func healthCheck(c *gin.Context) {
	c.Writer.WriteHeader(200)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/answers", getAllAnswers)
	router.GET("/answer", getRandomAnswer)
	router.GET("/health", healthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("Error: %s", err)
	}
}
