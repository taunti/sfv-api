package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taunti/sfv-api/pkg/cfn"
)

const token = ""

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["id"]

	cfn := cfn.NewCFN(token)
	profile := cfn.GetProfile(alias)

	fmt.Printf(
		"%s - %s (%v LP) - %d matches\n",
		profile.GetFighterId(),
		profile.GetLeague(),
		profile.GetLeaguePoints(),
		profile.GetTotalMatches(),
	)

	w.Write(profile.ToJSON())
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/profile/{id}", ProfileHandler)

	srv := http.Server{
		Handler: r,
		Addr:    "0.0.0.0:7777",
	}

	log.Fatal(srv.ListenAndServe())
}
