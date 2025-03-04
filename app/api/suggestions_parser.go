package api

import (
	"strings"
	"time"
)

type Date struct {
	Formatted string `json:"formatted"`
	Initial   string `json:"initial"`
}

type ISearch struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Pinner struct {
		ID        string `json:"id"`
		Username  string `json:"username"`
		FullName  string `json:"full_name"`
		AvatarURL string `json:"avatarURL"`
		Followers *int   `json:"followers"`
	} `json:"pinner"`
	Date     Date    `json:"date"`
	Type     string  `json:"type"`
	ImageURL string  `json:"imageURL"`
	Video    *string `json:"video"`
}

func formatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

func parseSuggestions(data map[string]interface{}) map[string]interface{} {
	root, ok := data["resource_response"].(map[string]interface{})
	if !ok {
		return nil
	}

	results, _ := root["data"].([]interface{})
	bookmark, _ := root["bookmark"].(string)

	var array []ISearch

	for _, item := range results {

		response, _ := item.(map[string]interface{})

		imageURL := response["images"].(map[string]interface{})["orig"].(map[string]interface{})["url"].(string)
		title := response["title"]
		if title == nil {
			title = response["grid_title"]
		}

		id := response["id"].(string)
		date := response["created_at"].(string)
		typeVal := response["type"].(string)
		pinner, _ := response["pinner"].(map[string]interface{})
		initialDate, _ := time.Parse(time.RFC3339, date)
		formattedDate := formatDate(initialDate, "2006-01-02")

		var cleanURL *string
		videos := response["videos"].(map[string]interface{})
		if videoList, ok := videos["video_list"].(map[string]interface{}); ok {
			if vHLSV4, ok := videoList["V_HLSV4"].(map[string]interface{}); ok {
				url := vHLSV4["url"].(string)
				replacement := strings.Replace(url, "/hls/", "/hevcMp4V2/", 1)
				cleanURLValue := strings.Replace(replacement, ".m3u8", "_t5.mp4", 1)
				cleanURL = &cleanURLValue
			}
		}

		pinnerData := struct {
			ID        string `json:"id"`
			Username  string `json:"username"`
			FullName  string `json:"full_name"`
			AvatarURL string `json:"avatarURL"`
			Followers *int   `json:"followers"`
		}{
			ID:        pinner["id"].(string),
			Username:  pinner["username"].(string),
			FullName:  pinner["full_name"].(string),
			AvatarURL: pinner["image_medium_url"].(string),
		}
		if followerCount, ok := pinner["follower_count"].(float64); ok && followerCount > 0 {
			followersCount := int(followerCount)
			pinnerData.Followers = &followersCount
		}

		array = append(array, ISearch{
			ID:       id,
			Title:    title.(string),
			Pinner:   pinnerData,
			Date:     Date{formattedDate, date},
			Type:     typeVal,
			ImageURL: imageURL,
			Video:    cleanURL,
		})
	}

	return map[string]interface{}{
		"bookmark": bookmark,
		"response": array,
	}
}
