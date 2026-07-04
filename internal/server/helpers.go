package server

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(`{"error": "` + msg + `"}`))
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Internal server error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func wordSwapper(sentence string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	for _, w := range badWords {
		pattern := `(?i)(^|\s)` + regexp.QuoteMeta(w) + `(\s|$)`
		re := regexp.MustCompile(pattern)
		sentence = re.ReplaceAllString(sentence, "${1}****${2}")
	}

	return sentence
}
