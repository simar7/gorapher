package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	strava "github.com/strava/go.strava"
)

// ListPastRuns lists the latest runs
func ListPastRuns(client *strava.Client, athleteID int64) {
	log.Println("Fetching latest runs...")
	activity, err := strava.NewAthletesService(client).ListActivities(athleteID).Do()
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range activity {
		if activity[key].Type.String() == "Run" {
			log.Printf("Run: %d | HR: %f | Distance: %2f", key, value.AverageHeartrate, value.Distance/1000)
		}
	}
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("you chose: ", vars["type"])

	client := strava.NewClient(authToken)

	ListPastRuns(client, athleteID)
}
