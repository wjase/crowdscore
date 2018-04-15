package heatsapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wjase/crowdscore/apis"

	"github.com/wjase/crowdscore/db"
)

// ConfigureRouting - configure paths for this api
func ConfigureRouting(router *mux.Router) {
	router.PathPrefix("/heats/{slug}").
		Methods("GET").
		HandlerFunc(getHeatBySlug)
	router.PathPrefix("/heats").
		Methods("GET").
		HandlerFunc(getHeats)
	router.PathPrefix("/heats").
		Methods("POST").
		HandlerFunc(postHeat)
}

func getHeats(w http.ResponseWriter, r *http.Request) {
	heats := []db.Heat{}
	db.DB.Model(&db.Heat{}).Find(&heats)
	apis.WriteRestResponse(heats, &w)
}

func getHeatBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	heats := []db.Heat{}
	db.DB.Model(&db.Heat{}).Where("slug = ?", vars["slug"]).Find(&heats)
	apis.WriteRestResponse(heats, &w)
}

func postHeat(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var v db.Heat
	err := decoder.Decode(&v)
	if err != nil {
		fmt.Println("Bad request" + fmt.Sprint(err))
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	// now save it
	db.DB.Create(&v)
	apis.WriteRestStatus(true, "", &w)
}
