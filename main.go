package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/torre/data"
	"github.com/torre/handler"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	f := data.NewFetch()

	r := mux.NewRouter()
	s := r.PathPrefix("/torre").Subrouter()
	// Add your routes as needed
	bh := handler.NewBioHandler(f)
	s.HandleFunc("/bio/{username}", bh.GetBioByUsername).Methods(http.MethodGet)

	jh := handler.NewJobHandler(f)
	s.HandleFunc("/job/{id}", jh.GetJobByID).Methods(http.MethodGet)

	jsh := handler.NewJobSearchHandler(f)
	s.HandleFunc("/search/job", jsh.Find).Methods(http.MethodPost)

	psh := handler.NewPeopleSearchHandler(f)
	s.HandleFunc("/search/people", psh.Find).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/jobs.html"
		}

		if strings.Contains(r.URL.Path, ".html") {
			r.URL.Path = `/examples` + r.URL.Path
		}

		fs.ServeHTTP(w, r)
	})).Methods("GET")

	port := os.Getenv("PORT")
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 120,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
