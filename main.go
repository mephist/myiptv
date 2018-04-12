package main

import (
	"flag"
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

var DefaultThreadNum int

func main() {
	log.SetLevel(log.DEBUG)

	flag.IntVar(&DefaultThreadNum, "t", 5, "Multi Download Thread Num")
	flag.Parse()

	http.HandleFunc("/sitv.m3u8", sitvHandler)
	http.HandleFunc("/byr.m3u8", byrHandler)
	http.HandleFunc("/youtube.m3u8", youtubeIndexHandler)
	http.HandleFunc("/litv.m3u8", litvHandler)
	http.HandleFunc("/4gtv.m3u8", fourgtvHandler)
	http.HandleFunc("/fjtv.m3u8", fjtvHandler)
	http.HandleFunc("/grtn.m3u8", grtnHandler)
	http.HandleFunc("/youjia.m3u8", youjiaHandler)
	http.HandleFunc("/inews.m3u8", inewsHandler)
	http.HandleFunc("/douyu.m3u8", douyuHandler)
	http.HandleFunc("/neu6.m3u8", neu6Handler)
	http.HandleFunc("/neu.m3u8", neuHandler)
	http.HandleFunc("/tuna.m3u8", tunaHandler)
	http.HandleFunc("/live/pool/", fourgtvTsHandler)
	http.HandleFunc("/live/", qmHandler)
	http.HandleFunc("/hls/", neu6tsHandler)
	http.HandleFunc("/neu/hls/", neutsHandler)
	http.HandleFunc("/tuna/hls/", tunaTsHandler)
	http.HandleFunc("/sdlive/", sctvHandler)
	http.HandleFunc("/hdlive/", sctvHandler)
	http.HandleFunc("/haixia/", fjtvApiHandler)
	http.HandleFunc("/haixia_sd/", fjtvApiHandler)
	http.HandleFunc("/youjia/", youjiaApiHandler)
	http.HandleFunc("/api/", youtubeApiHandler)
	http.HandleFunc("/videoplayback/", youtubeVideoHandler)
	http.HandleFunc("/hi/vod/", fourgtvApiHandler)
	http.HandleFunc("/", byrApiHandler)
	http.ListenAndServe(":8090", nil)
}
