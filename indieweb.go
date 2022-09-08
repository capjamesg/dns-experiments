package main

import (
	"log"
	"net/http"
	"time"

	"github.com/miekg/dns"
	"willnorris.com/go/microformats"
)

func parse_mf2(url string) *microformats.Data {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	content := microformats.Parse(resp.Body, nil)

	return content
}

func is_hwc_day(w dns.ResponseWriter, r *dns.Msg) {
	var message string

	if time.Now().Weekday() == time.Wednesday {
		message = "It's HWC Day!"
	} else {
		message = "It's not HWC day."
	}

	send_message(w, r, message)
}

func next_indieweb_event(w dns.ResponseWriter, r *dns.Msg) {
	url := "https://events.indieweb.org"

	content := parse_mf2(url)

	currentTime := time.Now()

	for _, item := range content.Items {
		if item.Type[0] == "h-feed" {
			for _, event := range item.Children {
				event_name := event.Properties["name"][0].(string)
				event_date := event.Properties["start"][0].(string)

				parsed_event_date, _ := time.Parse(time.RFC3339, event_date)

				if time.Time.After(parsed_event_date, currentTime) {
					send_message(w, r, event_name+" - "+parsed_event_date.String())
				}
			}
		}
	}

	send_message(w, r, "No new events found.")
}
