package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func grtnHandler(w http.ResponseWriter, r *http.Request) {
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
	uri := "http://www.gdtv.cn/m2o/channel/channel_info.php?id=" + id
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", "curl/7.52.1")
	resp, err := client.Do(req)
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
	re := regexp.MustCompile(`http:[^\"]*\.m3u8\?_upt=\w*`)
	hls := re.Find(body)
	dst := strings.Replace(string(hls), "\\", "", -1)
	req, _ = http.NewRequest("GET", dst, nil)
	req.Header.Set("User-Agent", "curl/7.52.1")
	resp, err = client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	re = regexp.MustCompile(`.*\.m3u8\?_upt=.*`)
	hls = re.Find(body)
	u, err := url.Parse(string(hls))
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	base, err := url.Parse(dst)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	w.Header().Set("Location", base.ResolveReference(u).String())
	http.Error(w, http.StatusText(302), 302)
}
