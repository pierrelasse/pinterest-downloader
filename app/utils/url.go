package utils

import (
	_url "net/url"
)

func URL_escape(url string) string {
	return _url.QueryEscape(url)
}
