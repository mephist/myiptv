package main

import (
	"io"
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func qmHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://liveal.quanmin.tv" + r.URL.Path
	log.Debug("Curl url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
