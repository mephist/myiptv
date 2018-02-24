package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	log "github.com/tominescu/double-golang/simplelog"
)

var gCurlTimeMap = make(map[string]int64)

func byrApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://tv.byr.cn:8888" + r.URL.Path
	log.Debug("Curl url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	if strings.HasSuffix(r.URL.Path, "index.m3u8") {
		go CurlCount(r.URL.Path)
	}
	io.Copy(w, resp.Body)
}

func CurlCount(path string) {
	ts := gCurlTimeMap[path]
	now := time.Now().Unix()
	if now-ts >= 25 {
		gCurlTimeMap[path] = now
		url := fmt.Sprintf("http://tv.byr.cn/player_count/res.gif?play_url=http://tv.byr.cn:8888%s&refer=http://tv.byr.cn/tv-show&title=BYR-IPTV", path)
		log.Debug("Curl count url: %s", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Warn("curl count url: %s error: %s", url, err)
		} else {
			resp.Body.Close()
		}
		if len(gCurlTimeMap) > 1000 {
			gCurlTimeMap = make(map[string]int64)
		}
	}
}
