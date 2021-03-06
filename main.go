package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	rand.Seed(time.Now().Unix())

	router := gin.Default()

	router.Static("/assets", "web/assets")

	router.GET("/vlad/:id", getAnime) // https://yummyanime.club/users/get-remote-list?user_id=107327&list_id=3
	router.GET("/ivan/:id", nil)      // https://shikimori.one/baho1015/list/anime/mylist/completed
	router.GET("/", func(c *gin.Context) {
		c.File("web/index.html")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "-_- такого тут нет",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not set")
	}

	router.Run(":" + port)
}

func getAnime(c *gin.Context) {
	type AnimeStruct struct {
		AnimeID int    `json:"anime_id"`
		Name    string `json:"name"`
		Alias   string `json:"alias"`
		Image   string `json:"image"`
	}
	id := c.Param("id")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://yummyanime.club/users/get-remote-list?user_id=107327&list_id="+id, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")
	req.Header.Add("Host", "localhost")

	list, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": "Ошибка запроса",
			"err":   err.Error(),
		})
		return
	}
	defer list.Body.Close()
	b, _ := ioutil.ReadAll(list.Body)
	var ListAnime []AnimeStruct
	if err := json.Unmarshal(b, &ListAnime); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": "Ошибка парсинга",
			"err":   err.Error(),
			"res":   string(b),
		})
		return
	}
	index := rand.Intn(len(ListAnime))
	c.JSON(http.StatusOK, ListAnime[index])
}
