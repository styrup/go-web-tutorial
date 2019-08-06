package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Println("Environment variable " + key + " not found, using default value: " + fallback)
		value = fallback
	}
	return value
}

func main() {
	serverPort := getEnv("SERVERPORT", "8000")
	serverAddr := ":" + serverPort
	router := mux.NewRouter()
	router.HandleFunc("/api/health", serverStatus)
	router.HandleFunc("/", getDataserviceValue)
	srv := &http.Server{
		Handler:      router,
		Addr:         serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server at port " + serverPort)
	log.Fatal(srv.ListenAndServe())

}

func serverStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getDataserviceValue(w http.ResponseWriter, r *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i1 := r1.Intn(100)
	dataserviceValue := strconv.Itoa(i1)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"dataserviceValue": dataserviceValue})
}
