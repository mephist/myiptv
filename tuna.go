package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/tominescu/double-golang/simplelog"
)

func tunaTsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "https://iptv.tsinghua.edu.cn" + strings.TrimPrefix(r.URL.Path, "/tuna")
	bt := time.Now()
	n, err := MultiDownload(w, url, 5)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	et := time.Now()
	cost := et.Sub(bt).Nanoseconds()
	log.Info("Curl url: %s download: %.2fMB cost: %.2fs speed:%.2fMbps threadNum:%d", url, float64(n)/1024/1024, float64(cost)/1e9, float64(n)/float64(cost)*1e9/1024/1024*8, DefaultThreadNum)
}

func tunaHandler(w http.ResponseWriter, r *http.Request) {
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
	url := "https://iptv.tsinghua.edu.cn/hls/" + id + ".m3u8"
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
	newBody := bytes.Replace(body, []byte("https://iptv.tsinghua.edu.cn"), []byte("http://"+r.Host+"/tuna"), -1)
	w.Write(newBody)
}
