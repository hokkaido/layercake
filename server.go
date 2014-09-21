package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/hokkaido/blink"
	_ "github.com/hokkaido/blink-mbtiles"
	"net/http"
	_ "net/url"
	"strconv"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func tileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//path, _ := vars["path"]
	zoom, _ := strconv.Atoi(vars["zoom"])
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	layer, err := blink.GetLayer("moots")

	if err != nil {
		fmt.Fprintf(w, "Error: %s!", err)
	}
	if layer != nil {
		tile, err := layer.GetTile(zoom, x, y)
		if err != nil {
			fmt.Fprintf(w, "Error: %s!", err)
		}
		w.Write(tile)
	} else {
		fmt.Fprintf(w, "Error: %s!", err)
	}
}

func main() {

	glog.Info("Startup...")
	blink.Load("config.json")

	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/map/{path}/{zoom}/{x}/{y}", tileHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3333", nil)
}
