package web

import (
	"log"
	"net"
	"net/http"
	"path"
	"strconv"

	"github.com/peng225/wolper/service"
)

type WebServer struct {
	port              int
	backendAddr       string
	backendPort       int
	staticContentPath string
}

func NewWebServer(
	port int, backendAddr string, backendPort int, staticContentPath string,
) *WebServer {
	return &WebServer{
		port:              port,
		backendAddr:       backendAddr,
		backendPort:       backendPort,
		staticContentPath: staticContentPath,
	}
}

func (ws *WebServer) Start() {
	http.HandleFunc("/", ws.indexHandler)
	http.HandleFunc("/query", ws.queryHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(ws.port), nil))
}

func (ws *WebServer) indexHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, path.Join(ws.staticContentPath, "index.html"))
}

func (ws *WebServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	key := query.Get("key")
	include := query.Get("include")
	exclude := query.Get("exclude")
	posExclude := query.Get("position exclude")
	uniqStr := query.Get("unique")
	entropySortStr := query.Get("entropy sort")
	uniq := false
	if uniqStr == "on" {
		uniq = true
	}
	entropySort := false
	if entropySortStr == "on" {
		entropySort = true
	}
	words := service.ClientQuery(net.JoinHostPort(ws.backendAddr, strconv.Itoa(ws.backendPort)),
		key, include, exclude, posExclude, uniq, entropySort)
	responseBody := ""
	if len(words) == 0 {
		responseBody = "No results found."
	} else {
		for _, word := range words {
			responseBody += word + "\n"
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}
