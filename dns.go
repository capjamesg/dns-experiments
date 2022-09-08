package main

import (
	"github.com/miekg/dns"
)

func send_message(w dns.ResponseWriter, r *dns.Msg, contents string) {
	message := new(dns.Msg)

	message.SetReply(r)

	message.Answer = append(message.Answer, &dns.TXT{Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{contents}})

	w.WriteMsg(message)
}

func main() {
	handler := dns.NewServeMux()

	handler.HandleFunc("resume", resume)
	handler.HandleFunc("is.it.hwc.day", is_hwc_day)
	handler.HandleFunc("next.indieweb.event", next_indieweb_event)
	handler.HandleFunc("recent.blog", most_recent_blog_post)

	server := &dns.Server{Addr: ":5003", Net: "udp", Handler: handler}

	server.ListenAndServe()

	defer server.Shutdown()
}
