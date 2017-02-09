package main

import (
	"flag"
	"log"
	"net/http"
)

func init() {
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

}

func main() {
	router := Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
