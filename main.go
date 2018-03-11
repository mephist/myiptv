package main

import (
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func main() {
	log.SetLevel(log.DEBUG)
	http.HandleFunc("/sitv.m3u8", sitvHandler)
	http.HandleFunc("/byr.m3u8", byrHandler)
	http.HandleFunc("/youtube.m3u8", youtubeIndexHandler)
	http.HandleFunc("/litv.m3u8", litvHandler)
	http.HandleFunc("/fjtv.m3u8", fjtvHandler)
	http.HandleFunc("/grtn.m3u8", grtnHandler)
	http.HandleFunc("/youjia.m3u8", youjiaHandler)
	http.HandleFunc("/live/", qmHandler)
	http.HandleFunc("/sdlive/", sctvHandler)
	http.HandleFunc("/hdlive/", sctvHandler)
	http.HandleFunc("/haixia/", fjtvApiHandler)
	http.HandleFunc("/haixia_sd/", fjtvApiHandler)
	http.HandleFunc("/youjia/", youjiaApiHandler)
	http.HandleFunc("/api/", youtubeApiHandler)
	http.HandleFunc("/videoplayback/", youtubeVideoHandler)
	http.HandleFunc("/", byrApiHandler)
	http.ListenAndServe(":8090", nil)
}
