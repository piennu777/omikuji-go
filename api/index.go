package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type LotteryResult struct {
	Prize string `json:"prize"`
}

var prizes = []string{
	"大当たり",
	"中当たり",
	"小当たり",
	"大外れ",
	"中外れ",
	"小外れ",
	"帰れ",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	result := LotteryResult{
		Prize: prizes[rand.Intn(len(prizes))],
	}

	json.NewEncoder(w).Encode(result)
}
