package request

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
	"github.com/laraXgram/Laraquest-Go/updates"
)

var update updates.Updates
var startCallback func(updates.Updates)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &update)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	if startCallback != nil {
		startCallback(update)
		update = updates.Updates{}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Serve() {
	http.HandleFunc("/", webhookHandler)
	log.Println("Server is running on 127.0.0.1:9000...")
	if err := http.ListenAndServe("127.0.0.1:9000", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func Start(callback func(updates.Updates)) {
	startCallback = callback
}