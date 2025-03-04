package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"pinterest-downloader/app/utils"
)

func GetPin(id string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"source_url": utils.Fmt("/pin/%s/", id),
		"data": map[string]interface{}{
			"options": map[string]interface{}{
				"id":                          id,
				"field_set_key":               "auth_web_main_pin",
				"noCache":                     true,
				"fetch_visual_search_objects": true,
			},
			"context": map[string]interface{}{},
		},
	}

	dataBytes, _ := json.Marshal(params["data"])
	apiURL := fmt.Sprintf("https://in.pinterest.com/resource/PinResource/get/?source_url=%s&data=%s",
		url.QueryEscape(params["source_url"].(string)),
		url.QueryEscape(string(dataBytes)),
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
