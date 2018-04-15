package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/wjase/crowdscore/apis/heatsapi"
	"github.com/wjase/crowdscore/apis/teamsapi"
	"github.com/wjase/crowdscore/db"
	"github.com/wjase/crowdscore/db/sqlite"

	"github.com/gorilla/mux"
)

var apiHandlers = make(map[string]http.HandlerFunc)

func main() {

	initSQLDB()

	router := mux.NewRouter()
	handleStatics(router, "css", "js", "img")

	router.PathPrefix("/api").
		Subrouter().
		HandleFunc("/jason", serveJason)
	initialiseAPIRouting(router)
	router.HandleFunc("/app/", serveTemplate)
	router.HandleFunc("/", redirectHome)
	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

func handleStatics(router *mux.Router, resourcePrefix ...string) {

	for _, prefix := range resourcePrefix {
		handleStatic(router, "/app/"+prefix+"/", "static/"+prefix)
	}
}

func handleStatic(router *mux.Router, urlToStrip string, staticRoot string) {

	staticFolder := http.Dir(staticRoot)
	fs := http.FileServer(staticFolder)
	router.NewRoute().PathPrefix(urlToStrip).Handler(http.StripPrefix(urlToStrip, fs))
}

func initialiseAPIRouting(router *mux.Router) {
	apiRouter := router.NewRoute().PathPrefix("/api").Subrouter()
	teamsapi.ConfigureRouting(apiRouter)
	heatsapi.ConfigureRouting(apiRouter)
	//apiRouter.PathPrefix("/").HandlerFunc(serveAPI)
}

func serveAPI(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	path = strings.TrimPrefix(path, "/api/")
	fmt.Fprintf(w, "A message from the other side:"+path)

}

func serveJason(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	path = strings.TrimPrefix(path, "/jason/")
	fmt.Fprintf(w, "A jason message from the other side:"+path)

}

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "app/", 301)

}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	path = strings.TrimPrefix(path, "/app")
	if path == "/" || path == "" {
		path = "/index.html"
	}
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", path)

	if fp == "templates" {
		fp = "templates/index.html"
	}

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			fp = "templates/index.html"
			info, err = os.Stat(fp)
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}
	var templateEngine = template.New("NewDelims").Delims("%{", "}%")
	tmpl, err := templateEngine.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func initSQLDB() {
	dbpath := "crowdscore" + ".db"
	dbToUse, err := sqlite.InitDB(dbpath)
	if err != nil || dbToUse == nil {
		panic(err)
	}

	db.Init(dbToUse)
	db.BootstrapData(dbToUse)

}
