package main

import (
	"fmt"
	"strconv"

	"github.com/miekg/dns"
)

func send_message(w dns.ResponseWriter, r *dns.Msg, contents string) {
	message := new(dns.Msg)

	message.SetReply(r)

	message.Answer = append(message.Answer, &dns.TXT{Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{contents}})

	w.WriteMsg(message)
}

func temperature_conversion(w dns.ResponseWriter, r *dns.Msg, formula string) {
	temperature := r.Question[0].Name

	input_temp, err := strconv.ParseFloat(temperature[2:], 32)

	if err != nil {
		send_message(w, r, "There was an error.")
	}

	var output_temp float64

	if formula == "celsius" {
		output_temp = (input_temp * float64(1.8)) + 32
	} else {
		output_temp = (float64(5) / float64(9)) * (input_temp - 32)
	}

	send_message(w, r, fmt.Sprintf("%f", output_temp))
}

func main() {
	handler := dns.NewServeMux()

	handler.HandleFunc("resume", resume)
	handler.HandleFunc("is.it.hwc.day", is_hwc_day)
	handler.HandleFunc("next.indieweb.event", next_indieweb_event)
	handler.HandleFunc("recent.blog", most_recent_blog_post)
	handler.HandleFunc("is.it.newsletter.day", is_newsletter_day)

	// fallback handler
	handler.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		fmt.Println(r.Question[0].Name[0:2])
		if r.Question[0].Name[0:2] == "cf" {
			temperature_conversion(w, r, "celsius")
		} else if r.Question[0].Name[0:2] == "fc" {
			temperature_conversion(w, r, "fahrenheit")
		}
	})

	server := &dns.Server{Addr: ":5003", Net: "udp", Handler: handler}

	server.ListenAndServe()

	defer server.Shutdown()
}
