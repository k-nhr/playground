package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

type search struct {
	Key      string
	EngineID string
	Type     string
	Count    string
}

type result struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind        string `json:"kind"`
		Title       string `json:"title"`
		HTMLTitle   string `json:"htmlTitle"`
		Link        string `json:"link"`
		DisplayLink string `json:"displayLink"`
		Snippet     string `json:"snippet"`
		HTMLSnippet string `json:"htmlSnippet"`
		Mime        string `json:"mime"`
		Image       struct {
			ContextLink     string `json:"contextLink"`
			Height          int    `json:"height"`
			Width           int    `json:"width"`
			ByteSize        int    `json:"byteSize"`
			ThumbnailLink   string `json:"thumbnailLink"`
			ThumbnailHeight int    `json:"thumbnailHeight"`
			ThumbnailWidth  int    `json:"thumbnailWidth"`
		} `json:"image"`
	} `json:"items"`
}

func (b *Bot) image(word string) ([]slack.Attachment, error) {
	var attachments []slack.Attachment

	baseURL := "https://www.googleapis.com/customsearch/v1"
	s := search{os.Getenv("CUSTOM_SEARCH_APIKEY"), os.Getenv("CUSTOM_SEARCH_ENGINE_ID"), "image", "10"}
	word = strings.TrimSpace(word)
	url := baseURL + "?key=" + s.Key + "&cx=" + s.EngineID + "&searchType=" + s.Type + "&num=" + s.Count + "&q=" + word

	imageURL, err := getImage(url)
	if err != nil {
		attachments = []slack.Attachment{slack.Attachment{
			Pretext: imageURL,
		}}
		return attachments, err
	}

	attachments = []slack.Attachment{slack.Attachment{
		ImageURL: imageURL,
	}}
	return attachments, nil
}

func getImage(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return "", err
	}
	if data.Items == nil {
		return "", fmt.Errorf("could not retrieve images")
	}

	rand.Seed(time.Now().UnixNano())
	return data.Items[rand.Intn(len(data.Items))].Link, nil
}
