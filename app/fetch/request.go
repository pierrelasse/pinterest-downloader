package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	net_url "net/url"
	"pinterest-downloader/app/utils"
	"time"
)

var counter uint16

func Request(url string, method string) (*http.Response, error) {
	counter++

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		utils.Console_writeln(utils.Fmt(utils.FRed+"[fetch #%d] new-request %v"+utils.Reset, counter, err))
		return nil, utils.Err("fetch #%d (new-request)", counter)
	}

	req.Header.Set("x-pinterest-pws-handler", "www/pin/[id].js")
	req.Header.Set("x-pinterest-source-url", url)
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("User-Agent", USER_AGENT)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.Console_writeln(utils.Fmt(utils.FRed+"[fetch #%d] do %v"+utils.Reset, counter, err))
		return nil, utils.Err("fetch #%d (do)", counter)
	}

	if resp.StatusCode != 200 {
		utils.Console_writeln(utils.Fmt(utils.FRed+"[fetch #%d] status %s"+utils.Reset, counter, resp.Status))
		return nil, utils.Err("fetch #%d (status)", counter)
	}

	return resp, nil
}

func Get(endpoint string, sourceURL string, data utils.JSON) (utils.JSON, error) {
	utils.Console_writeln(utils.FGray + utils.Fmt("[fetch #%d] %s", counter, endpoint) + utils.Reset)

	encodedData, _ := json.Marshal(data)
	encodedDataStr := string(encodedData)

	url := utils.Fmt(
		"%s%s?source_url=%s&data=%s&_=%d",
		BASE_URL,
		endpoint,
		net_url.QueryEscape(sourceURL),
		net_url.QueryEscape(encodedDataStr),
		time.Now().UnixMilli(),
	)

	resp, err := Request(url, "GET")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respJSON utils.JSON
	if err := json.NewDecoder(resp.Body).Decode(&respJSON); err != nil {
		body, _ := io.ReadAll(resp.Body)
		utils.Console_writeln(utils.Fmt(utils.FRed+"[fetch #%d] decode-body %v\n> %s"+utils.Reset, counter, err, string(body))) // [:100]
		return nil, utils.Err("fetch #%d (decode-body)", counter)
	}

	return respJSON, nil
}
