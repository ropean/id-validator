package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"os"

	idvalidator "github.com/guanguans/id-validator"
)

//go:embed web
var webFiles embed.FS

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// GET /validate?id=xxx&strict=true
func handleValidate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	strict := r.URL.Query().Get("strict") == "true"
	writeJSON(w, http.StatusOK, response{Success: true, Data: idvalidator.IsValid(id, strict)})
}

// GET /info?id=xxx&strict=true
func handleInfo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	strict := r.URL.Query().Get("strict") == "true"
	info, err := idvalidator.GetInfo(id, strict)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{Success: false, Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response{Success: true, Data: info})
}

// GET /fake
func handleFake(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{Success: true, Data: idvalidator.FakeId()})
}

// GET /upgrade?id=xxx
func handleUpgrade(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result, err := idvalidator.UpgradeId(id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, response{Success: false, Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response{Success: true, Data: result})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var staticHandler http.Handler
	if os.Getenv("DEV") == "1" {
		println("DEV mode: serving static files from ./cmd/server/web")
		staticHandler = http.FileServer(http.Dir("cmd/server/web"))
	} else {
		sub, err := fs.Sub(webFiles, "web")
		if err != nil {
			println("ERROR: failed to load embedded web files:", err.Error())
			os.Exit(1)
		}
		staticHandler = http.FileServer(http.FS(sub))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/validate", handleValidate)
	mux.HandleFunc("/info", handleInfo)
	mux.HandleFunc("/fake", handleFake)
	mux.HandleFunc("/upgrade", handleUpgrade)
	mux.Handle("/", staticHandler)

	println("Server listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		println("ERROR:", err.Error())
	}
}
