package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LYH2263/go-cronexpr/planner"
)

func main() {
	addr := flag.String("addr", ":8222", "listen address")
	tzName := flag.String("tz", "Local", "default timezone")
	flag.Parse()
	loc := time.Local
	if *tzName != "Local" {
		if l, err := time.LoadLocation(*tzName); err == nil {
			loc = l
		}
	}
	api := &planner.API{DefaultTZ: loc}
	srv := planner.New(api)
	log.Printf("cronexpr planner on %s", *addr)
	if err := http.ListenAndServe(*addr, srv.Handler); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
