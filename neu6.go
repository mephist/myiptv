package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/tominescu/double-golang/simplelog"
)

func neu6tsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://ts.media.neu6.edu.cn" + r.URL.Path
	bt := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	n, _ := MyCopy(w, resp.Body)
	et := time.Now()
	cost := et.Sub(bt).Nanoseconds()
	log.Info("Curl url: %s cost: %.2fs download: %.2fMB speed:%.2fMbps", url, float64(cost)/1e9, float64(n)/1024/1024, float64(n)/float64(cost)*1e9/1024/1024*8)
}

func neu6Handler(w http.ResponseWriter, r *http.Request) {
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
	url := "http://media2.neu6.edu.cn/hls/" + id + ".m3u8"
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
	newBody := bytes.Replace(body, []byte("ts.media.neu6.edu.cn"), []byte(r.Host), -1)
	w.Write(newBody)
}
