// tombo far off
// AQ <aq@okaq.com>
// 2020-10-04
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "zoho.html"
	NANO = "nano/"
)

func motd() {
	fmt.Println("web localhost:8080")
	fmt.Println(time.Now().String())
}

func TomboHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", TomboHandler)
	http.Handle("/nano/", http.StripPrefix("/nano/", http.FileServer(http.Dir(NANO))))
	http.ListenAndServe(":8080", nil)
}


