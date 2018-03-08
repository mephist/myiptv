package main

import (
	"bufio"
	"net/http"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

func youjiaApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Request URL:%s", r.URL)
	url := "http://gslb.gcable.cn:8070" + strings.TrimPrefix(r.URL.String(), "/youjia")
	log.Debug("New Url:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-type", "application/vnd.apple.mpegurl")
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "#") {
			line = "http://" + r.Host + "/youjiaVideo/" + line
		}
		w.Write([]byte(line))
		w.Write([]byte("\r\n"))
		continue
	}
}
