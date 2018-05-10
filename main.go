//go:generate go-bindata-assetfs -pkg main -o files.go skeleton.css

package main

import (
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

func main() {
	srvmux := http.NewServeMux()

	assFS := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: ""}
	srvmux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assFS)))

	srvmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path == "/" {
			IndexPage(w, ControlForm())
		} else {
			http.NotFound(w, r)
		}
	})

	httpsrv := http.Server{Addr: ":8000", Handler: srvmux}
	defer httpsrv.Close()

	log.Println(httpsrv.ListenAndServe())
}
