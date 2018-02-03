package main

import (
	"io/ioutil"
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func sitvHandler(w http.ResponseWriter, r *http.Request) {
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
	url := "http://www.sitv.com.cn/GetPlayPath/GetPlayPath?type=LIVE&se=sitv&ct=1&code=" + id
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
	w.Header().Set("Location", string(body))
	http.Error(w, http.StatusText(302), 302)
}
