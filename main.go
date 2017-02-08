package main

import (
	"flag"
	"log"

	"github.com/strava/go.strava"
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

func main() {

	var authToken string
	var athleteID int64

	flag.StringVar(&authToken, "authtoken", "", "Authorization Token")
	flag.Int64Var(&athleteID, "athid", 0, "Athlete ID")

	flag.Parse()

	if authToken == "" {
		flag.PrintDefaults()
		log.Fatalln("Please provide an auth token, setup one at https://www.strava.com/settings/api")
	}

	if athleteID == 0 {
		flag.PrintDefaults()
		log.Fatalln("Please provide an athlete ID.")
	}

	client := strava.NewClient(authToken)

	ListPastRuns(client, athleteID)
}
