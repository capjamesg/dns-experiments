package main

import (
	"github.com/miekg/dns"
)

func resume(w dns.ResponseWriter, r *dns.Msg) {
	message := new(dns.Msg)

	message.SetReply(r)

	var lines_of_text []string

	lines_of_text = append(lines_of_text, "Hello there, I am James (https://jamesg.blog).")
	lines_of_text = append(lines_of_text, "Projects: IndieWeb Search, implementations of WebSub, IndieAuth, Webmention, Micropub, and Microsub")
	lines_of_text = append(lines_of_text, "More information: https://jamesg.blog/resume")
	lines_of_text = append(lines_of_text, "Contact: jamesg@jamesg.blog")

	for _, line := range lines_of_text {
		message.Answer = append(message.Answer, &dns.TXT{Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{line}})
	}

	w.WriteMsg(message)
}

func most_recent_blog_post(w dns.ResponseWriter, r *dns.Msg) {
	url := "https://jamesg.blog"

	content := parse_mf2(url)

	for _, item := range content.Items {
		if item.Type[0] == "h-feed" {
			for _, article := range item.Children {
				title := article.Properties["name"][0].(string)
				url := article.Properties["url"][0].(string)

				message := title + " - " + url

				send_message(w, r, message)
			}
		}
	}
}
