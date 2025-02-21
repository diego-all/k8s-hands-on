package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}

	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// resp := fmt.Sprintf("La hora es %v y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))
	// w.Write([]byte(`{"message": "hello world"}`))

	// bytes array
	// w.Write([]byte(resp))
	w.Write([]byte(jsonResp))
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
