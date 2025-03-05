package api

import (
	"pinterest-downloader/app/fetch"
	"pinterest-downloader/app/utils"
)

func Suggestions(id string, bookmark string) (*parseSuggestions_Result, error) {
	data := utils.JSON{
		"options": utils.JSON{
			"pin_id":                 id,
			"context_pin_ids":        []string{},
			"page_size":              12,
			"search_query":           "",
			"source":                 "deep_linking",
			"top_level_source":       "deep_linking",
			"top_level_source_depth": 1,
			"is_pdp":                 false,
		},
		"context": utils.JSON{},
	}
	if bookmark != "" {
		data["options"].(utils.JSON)["bookmarks"] = []string{bookmark}
	}
	resp, err := fetch.Get("/resource/RelatedModulesResource/get/", utils.Fmt("/pin/%s/", id), data)
	if err != nil {
		return nil, err
	}
	result := parseSuggestions(resp)
	if result == nil {
		return nil, utils.Err("could not parse response")
	}
	return result, nil
}
