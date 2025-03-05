package api

import (
	"pinterest-downloader/app/utils"
	"strings"
	"time"
)

type Date struct {
	Formatted string
	Initial   string
}

type ISearch_Pinner struct {
	ID        string
	Username  string
	FullName  string
	AvatarURL string
	Followers *int
}

type ISearch struct {
	ID       string
	Title    string
	Pinner   ISearch_Pinner
	Date     Date
	Type     string
	ImageURL string
	Video    *string
}

type parseSuggestions_Result struct {
	Bookmark string
	Response []ISearch
}

func formatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

func parseSuggestions(d utils.JSON) *parseSuggestions_Result {
	resourceResponse, ok := d["resource_response"].(utils.JSON)
	if !ok {
		return nil
	}

	data := resourceResponse["data"].([]interface{})
	bookmark, ok := resourceResponse["bookmark"].(string)
	if !ok {
		return nil
	}

	var array []ISearch

	for _, i := range data {
		item := i.(utils.JSON)

		images, ok := item["images"]
		if !ok {
			continue
		}

		image := images.(utils.JSON)["orig"].(utils.JSON)
		imageURL := image["url"].(string)

		title := item["title"]
		if title == nil {
			title = item["grid_title"]
		}

		id := item["id"].(string)
		date := item["created_at"].(string)
		typeVal := item["type"].(string)
		pinner, _ := item["pinner"].(utils.JSON)
		initialDate, _ := time.Parse(time.RFC3339, date)
		formattedDate := formatDate(initialDate, "2006-01-02")

		var cleanURL *string
		videos, ok := item["videos"].(utils.JSON)
		if ok {
			if videoList, ok := videos["video_list"].(utils.JSON); ok {
				if vHLSV4, ok := videoList["V_HLSV4"].(utils.JSON); ok {
					url := vHLSV4["url"].(string)
					replacement := strings.Replace(url, "/hls/", "/hevcMp4V2/", 1)
					cleanURLValue := strings.Replace(replacement, ".m3u8", "_t5.mp4", 1)
					cleanURL = &cleanURLValue
				}
			}
		}

		pinnerData := ISearch_Pinner{
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

	return &parseSuggestions_Result{
		Bookmark: bookmark,
		Response: array,
	}
}
