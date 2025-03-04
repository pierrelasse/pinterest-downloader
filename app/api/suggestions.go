package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pinterest-downloader/app/fetch"
	"pinterest-downloader/app/utils"
)

type JSON map[string]interface{}

func formatMap(data any) string {
	formatted, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "{}"
	}
	return string(formatted)
}

func Suggestions(id string, bookmark string) (JSON, error) {
	if id == "" {
		return nil, utils.Err("No id specified")
	}

	var bookmarks any
	// if bookmark == "" {
	// 	bookmarks = "[null]"
	// } else {
	// 	bookmarks = []string{bookmark}
	// }
	bookmarks = "[]"

	source_url := fmt.Sprintf("/pin/%s/", id)

	params := JSON{
		"source_url": source_url,
		"data": JSON{
			"options": JSON{
				"pin_id":                 id,
				"context_pin_ids":        "[]",
				"page_size":              12,
				"search_query":           "",
				"source":                 "deep_linking",
				"top_level_source":       "deep_linking",
				"top_level_source_depth": 1,
				"is_pdp":                 false,
				"bookmarks":              bookmarks,
			},
			"context": "{}",
		},
	}

	paramsJSON, _ := json.Marshal(params)
	url := utils.Fmt(
		"%s/resource/RelatedModulesResource/get/?source_url=%s&data=%s",
		fetch.BASE_URL,
		url.QueryEscape(source_url),
		url.QueryEscape(string(paramsJSON)),
	)

	utils.Console_writeln(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data JSON
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	utils.Console_writeln(formatMap(data))

	return parseSuggestions(data), nil
}
