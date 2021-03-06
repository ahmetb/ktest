package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/mchmarny/ktest/utils"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"name":    "ktest",
		"on":      time.Now().String(),
		"version": utils.MustGetEnv("RELEASE", "RELEASE not set"),
	}

	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
		log.Printf("Error in home template: %s", err)
	}

}
