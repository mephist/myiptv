package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(503), 503)
		return
	}
	id := r.Form.Get("id")
	if id == "" {
		http.Error(w, http.StatusText(503), 503)
		return
	}
	url := "http://tv.byr.cn/tv-show-detail/" + id
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	re := regexp.MustCompile(`http://.*/index.m3u8`)
	hls := re.Find(body)
	if len(hls) < 8 {
		http.Error(w, "Cant't find m3u8 url", 503)
		return
	}
	dst := strings.Replace(string(hls), "http://tv.byr.cn:8888", "http://"+r.Host, 1)
	w.Header().Set("Location", dst)
	http.Error(w, http.StatusText(302), 302)
}
