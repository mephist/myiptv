package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/tominescu/double-golang/simplelog"
)

const AUTH_URL = "http://portal.gcable.cn:8080/PortalServer-App/new/aaa_aut_aut002"

type Result struct {
	Data Data `json:"data"`
}

type Data struct {
	AuthResult string `json:"authResult"`
}

func youjiaHandler(w http.ResponseWriter, r *http.Request) {
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
	cid := r.Form.Get("cid")
	if cid == "" {
		http.Error(w, http.StatusText(503), 503)
		return
	}

	postData := "ptype=8&plocation=1601&puser=freeuser&ptoken=c1b3535d893682e6&pversion=030101&pserverAddress=http%3A%2F%2Fportal.gcable.cn%3A8080&pserialNumber=c1b3535d893682e6&DRMtoken=&epgID=&authType=0&secondAuthid=&t=c1b3535d893682e6&pid=&cid=" + cid + "&u=freeuser&p=8&l=1601&d=c1b3535d893682e6&n=" + id + "&v=2"
	resp, err := http.Post(AUTH_URL, "application/x-www-form-urlencoded", strings.NewReader(postData))
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
	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	log.Debug("AuthResult:%s", result.Data.AuthResult)
	u, err := url.Parse(result.Data.AuthResult)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	q := u.Query()
	dst := "http://gslb.gcable.cn:8070/live/" + id + ".m3u8?t=c1b3535d893682e6&u=freeuser&p=8&pid=&cid=" + cid + "&l=1601&d=c1b3535d893682e6&sid=" + q.Get("sid") + "&r=" + q.Get("r") + "&e=" + q.Get("e") + "&nc=" + q.Get("nc") + "&a=" + q.Get("a") + "&v=2"
	w.Header().Set("Location", dst)
	http.Error(w, dst, 302)
}
