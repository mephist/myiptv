package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func neu6tsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://ts.media.neu6.edu.cn" + r.URL.Path
	log.Debug("Curl url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	//io.Copy(w, resp.Body)
	buf := make([]byte, 8192)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			break
		}
		w.Write(buf[0:n])
	}
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
