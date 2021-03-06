package teamsapi

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
	router.PathPrefix("/teams/{slug}").
		Methods("GET").
		HandlerFunc(getTeamBySlug)
	router.PathPrefix("/teams").
		Methods("GET").
		HandlerFunc(getTeams)
	router.PathPrefix("/teams").
		Methods("POST").
		HandlerFunc(postTeam)
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	teams := []db.Team{}
	db.DB.Model(&db.Team{}).Find(&teams)
	apis.WriteRestResponse(teams, &w)
}

func getTeamBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teams := []db.Team{}
	db.DB.Model(&db.Team{}).Where("slug = ?", vars["slug"]).Find(&teams)
	apis.WriteRestResponse(teams, &w)
}

func postTeam(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u db.Team
	err := decoder.Decode(&u)
	if err != nil {
		fmt.Println("Bad request" + fmt.Sprint(err))
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	// now save it
	db.DB.Create(&u)
	apis.WriteRestStatus(true, "", &w)
}
