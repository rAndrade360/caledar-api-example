package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/rAndrade360/TestCalendarAPI/gfunctions"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := gfunctions.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	event, err := gfunctions.CreateEvent(srv)
	if err != nil {
		log.Fatal(err)
	}

	gfunctions.GetEvents(srv)

	gfunctions.DeleteEvent(srv, event.Id)

}
