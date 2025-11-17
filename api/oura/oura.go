package oura

import (
	"errors"
	"math/rand"
)

type Snapshot struct {
	IsSleeping bool `json:"isSleeping"`
}

func getSleepStatus() (bool, error) {
	randomNumber := rand.Intn(10)

	var isSleeping bool

	// Add error checking with request op
	isSleeping = randomNumber == 4

	return isSleeping, nil
}

func GetSnapshot() (Snapshot, error) {
	var isSleeping bool
	isSleeping, err := getSleepStatus()

	if err != nil {
		return Snapshot{}, errors.New("error retrieving OURA sleep data")
	}

	snapshot := Snapshot{
		IsSleeping: isSleeping,
	}

	return snapshot, nil
}
