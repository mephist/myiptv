package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func sctvHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://scgctvshow.sctv.com" + r.URL.Path
	log.Debug("Curl URL:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	if strings.HasSuffix(r.URL.Path, ".ts") {
		io.Copy(w, resp.Body)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	content := strings.Replace(string(body), "3-", "1-", -1)
	w.Write([]byte(content))
}
