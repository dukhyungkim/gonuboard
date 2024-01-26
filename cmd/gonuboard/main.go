package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/dukhyungkim/gonuboard/install"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

var (
	NeedInstall = false
)

func main() {
	parseFlags()

	if FlagVersion {
		printVersion()
		return
	}

	if FlagHelp {
		flag.Usage()
		return
	}

	err := loadEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run() error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", defaultHandler)
	r.Route("/install", install.DefaultRouter)

	fileServer := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	addr := ":8080"
	fmt.Printf("running on %s\n", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		return err
	}
	return nil
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			NeedInstall = true
		} else {
			return err
		}
	}
	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if NeedInstall {
		http.Redirect(w, r, "/install", http.StatusMovedPermanently)
		return
	}

	_, _ = w.Write([]byte("hello gnuboard"))
}
