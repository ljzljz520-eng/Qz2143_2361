package main

import (
	"community50/api"
	"community50/audit"
	"community50/config"
	"community50/intake"
	"community50/query"
	"community50/store"
	"community50/workflow"
	"log"
	"net/http"
)

func main() {
	c := config.Load()
	db, e := store.Open(c.DataPath)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	i := intake.New(db)
	f := workflow.New(db)
	q := query.New(db)
	_ = audit.New(db)
	log.Printf("community50 listening on %s", c.Addr)
	log.Fatal(http.ListenAndServe(c.Addr, api.New(i, f, q).Routes()))
}
