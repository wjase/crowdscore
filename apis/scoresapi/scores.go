package scoresapi

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
	router.PathPrefix("/scores/").
		Methods("POST").
		HandlerFunc(postVote)
	router.PathPrefix("/scores/summary/{heatslug}/{teamslug}").
		Methods("GET").
		HandlerFunc(getVoteSummaryByheat)
}

/*
	type Score struct {
	UserID     int
	HeatID     int
	CategoryID int
	ProjectID  int
	Rating     int
}
*/

type ScoreRequest struct {
	Heat     string
	Team     string
	Category string
	Score    int
}

func postVote(w http.ResponseWriter, r *http.Request) {
	user, err := apis.PrincipalFromContext(r)
	if err != nil {
		fmt.Println("ERROR: Couldn't save vote.No valid user in context")
		fmt.Println(err)
		apis.WriteRestStatus(false, "Couldn't save vote.Try again soon.", &w)
	}

	scoreRequest := readScoreRequestFromBody(w, r)
	previousVote := db.SingleVoteRegister{Team: scoreRequest.Team, User: user.Username, Heat: scoreRequest.Heat}
	if db.DB.Where(&previousVote).First(&previousVote).RecordNotFound() {
		scoringCategory := db.ScoringCategory{Slug: scoreRequest.Category}
		db.DB.Where(&scoringCategory).Find(&scoringCategory)
		team := db.Team{Slug: scoreRequest.Team}
		db.DB.Where(&team).Find(&team)
		heat := db.Heat{Slug: scoreRequest.Heat}
		heat := db.Heat{Slug: scoreRequest.Heat}
		db.DB.Where(&heat).Find(&heat)

		if db.DB.Error != nil {
			fmt.Println("ERROR: Couldn't save vote.")
			fmt.Println(db.DB.Error)
			apis.WriteRestStatus(false, "Couldn't save vote.Try again soon.", &w)
		}

		db.DB.Begin()
		defer db.DB.Commit()
		db.DB.
			Create(&db.Score{HeatID: heat.ID,
				CategoryID: scoringCategory.ID,
				TeamID:     team.ID,
				Rating:     scoreRequest.Score,
			}).
			Create(&previousVote)
		if db.DB.Error != nil {
			fmt.Println("ERROR: Couldn't save vote.")
			fmt.Println(db.DB.Error)
			apis.WriteRestStatus(false, "Couldn't save vote.Try again soon.", &w)
		}
	}

	apis.WriteRestStatus(true, "", &w)
}

func readScoreRequestFromBody(w http.ResponseWriter, r *http.Request) ScoreRequest {
	decoder := json.NewDecoder(r.Body)

	v := ScoreRequest{}
	err := decoder.Decode(&v)
	if err != nil {
		fmt.Println("Bad request" + fmt.Sprint(err))
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	return v
}
