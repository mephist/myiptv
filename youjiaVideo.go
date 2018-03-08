package main

import (
	"io"
	"net/http"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func youjiaVideoHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://27.36.62.161:8060/live/" + strings.TrimPrefix(r.URL.String(), "/youjiaVideo")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
