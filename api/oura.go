package handler

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
)

type Snapshot struct {
	IsSleeping bool `json:"isSleeping"`
}

func getSleepStatus() (bool, error) {
	randomNumber := rand.Intn(10)
	isSleeping := randomNumber == 4
	return isSleeping, nil
}

func getSnapShot() (Snapshot, error) {
	isSleeping, err := getSleepStatus()

	if err != nil {
		return Snapshot{}, errors.New("error retrieving OURA sleep data")
	}

	snapshot := Snapshot{
		IsSleeping: isSleeping,
	}

	return snapshot, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	snapshot, err := getSnapShot()
	if err != nil {
		http.Error(w, "Failure retrieving OURA snapshot", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshot)
}
