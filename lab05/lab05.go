package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type VideoInfo struct {
	Title        string
	ChannelTitle string
	LikeCount    string
	ViewCount    string
	PublishedAt  string
	CommentCount string
	Id           string
}

func init() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func formatNumber(input string) string {
	if len(input) <= 3 {
		return input
	}

	// Split the string into integer and fractional parts (if any)

	// Initialize the result string
	result := input[len(input)-3:]

	// Process the rest of the integer part in reverse, adding commas every three digits
	for i := len(input) - 6; i >= 0; i -= 3 {
		result = input[i:i+3] + "," + result
		if (i-3<0) {
			result = input[0:i] + "," + result
		}
	}
	return result
}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("v")

	if videoID == "" {
		http.ServeFile(w, r, "error.html")
		print("1")
		return
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	dataURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?key=%s&id=%s&part=snippet,statistics", apiKey, videoID)
	
	res, err := http.Get(dataURL)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		print("2")
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		print("3")
		return
	}

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		print("4")
		return
	}
	
	items, itemsExist := m["items"].([]interface{})
	if !itemsExist || len(items) == 0 {
		http.ServeFile(w, r, "error.html")
		print(len(m))
		print(dataURL)
		return
	}

	item, isMap := items[0].(map[string]interface{})
	if !isMap {
		http.ServeFile(w, r, "error.html")
		print("6")
		return
	}

	snippet, snippetExist := item["snippet"].(map[string]interface{})
	if !snippetExist {
		http.ServeFile(w, r, "error.html")
		print("7")
		return
	}

	statistics, statisticsExist := item["statistics"].(map[string]interface{})
	if !statisticsExist {
		http.ServeFile(w, r, "error.html")
		print("8")
		return
	}

	date := strings.Split(snippet["publishedAt"].(string), "T")
	dataArray := strings.Split(date[0], "-")
	dateString := fmt.Sprintf("%s年%s月%s日", dataArray[0], dataArray[1], dataArray[2])

	videoInfo := VideoInfo{
		Title:        snippet["title"].(string),
		ChannelTitle: snippet["channelTitle"].(string),
		LikeCount:    formatNumber(statistics["likeCount"].(string)),
		ViewCount:    formatNumber(statistics["viewCount"].(string)),
		PublishedAt:  dateString,
		CommentCount: formatNumber(statistics["commentCount"].(string)),
		Id:           videoID,
	}

	err = template.Must(template.ParseFiles("index.html")).Execute(w, videoInfo)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		print("9")
		return
	}
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
