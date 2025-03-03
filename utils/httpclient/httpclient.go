package httpclient

import (
	"bytes"
	"net/http"
)

// Get issues the standard http.Get request but with a custom User-Agent header
func Get(url string) (*http.Response, error) {
	client := &http.Client{}
	var data bytes.Buffer
	if request, err := http.NewRequest("GET", url, &data); err == nil {
		request.Header.Set("User-Agent", "ada-bot / https://github.com/Meldream/ada-bot")
		if response, err := client.Do(request); err == nil {
			return response, nil
		} else {
			return nil, err // Error at client.Do() call
		}
	} else {
		return nil, err // Error at http.NewRequest() call
	}
}
