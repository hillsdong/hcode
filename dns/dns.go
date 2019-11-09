package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/miekg/dns"
)

func main() {
	srv := &dns.Server{Addr: ":53", Net: "udp"}

	go srv.ListenAndServe()

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Authoritative = true
		w.WriteMsg(m)

		fmt.Printf("Listen %v", m)
	})

	fmt.Printf("Started On %s", "53")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sig
	fmt.Println("Stoped")
}
