package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	router.Use(cors.New(config))

	router.GET("/ping", healthCheck)
	router.GET("/birds", getBirds)
	router.GET("/state", getStateData)

	router.Run("0.0.0.0:8081")
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func healthCheck(c *gin.Context) {
	c.String(200, "Pong")
}

func getBirds(c *gin.Context) {
	regionCode := "US-IN"
	speciesCode := "cangoo"
	period := 7
	url := fmt.Sprintf("https://api.ebird.org/v2/data/obs/%s/recent/%s?back=%d", regionCode, speciesCode, period)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	apiKey := getEnvVar("API-KEY")
	req.Header.Add("X-eBirdApiToken", apiKey)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	c.Data(200, "application/json", b)
}

func getStateData(c *gin.Context) {
	jsonFile, err := os.Open("indianaTopoJson.json")
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Print(err)
	}
	c.Data(200, "application/json", data)
}
