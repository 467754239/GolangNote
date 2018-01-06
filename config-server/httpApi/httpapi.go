package httpApi

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/467754239/GolangNote/config-server/cfgLoader"
	log "github.com/auxten/logrus"
)

var mux map[string]func(http.ResponseWriter, *http.Request, *myHandler)

type myHandler struct {
	cfgm *cfgLoader.CfgManager
}

func HTTPServerStart(listenPort int, cfgm *cfgLoader.CfgManager) error {
	strListenPort := fmt.Sprintf(":%d", listenPort)
	server := http.Server{
		Addr: strListenPort,
		Handler: &myHandler{
			cfgm: cfgm,
		},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request, *myHandler))
	mux[""] = hello
	mux["conf"] = conf

	return server.ListenAndServe()
}

func (http_handler *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*
	  request URL like "http://127.0.0.1:2120" makes r.URL.String == "/"
	*/

	clean_url := filepath.Clean(r.URL.Path)
	clean_url_list := strings.Split(clean_url, "/")

	url_index := clean_url_list[1]
	//
	//log.Debug(clean_url_list)
	//log.Debug(url_index)

	if h, ok := mux[url_index]; ok {
		h(w, r, http_handler)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, clean_url+" conf not found")
}

func hello(w http.ResponseWriter, r *http.Request, h *myHandler) {
	io.WriteString(w, "<h1>Welconme to cfg-server</h1>")
}

func conf(w http.ResponseWriter, r *http.Request, h *myHandler) {
	clean_url := filepath.Clean(r.URL.Path)
	clean_url_list := strings.Split(clean_url, "/")

	log.Debug(clean_url, clean_url_list)

	req_args := r.URL.Query()
	var resp_str string
	if req_args.Get("flat") == "1" {
		resp_str = string(h.cfgm.GetCfg_json(clean_url_list[2:], cfgLoader.Flat))
	} else {
		resp_str = string(h.cfgm.GetCfg_json(clean_url_list[2:], cfgLoader.Raw))
	}
	log.Debug(resp_str)
	io.WriteString(w, resp_str)
}
