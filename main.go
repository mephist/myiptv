package main

import (
	"net/http"

	log "github.com/tominescu/double-golang/simplelog"
)

func main() {
	log.SetLevel(log.DEBUG)
	http.HandleFunc("/sitv.m3u8", sitvHandler)
	http.HandleFunc("/byr.m3u8", indexHandler)
	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8090", nil)
}
