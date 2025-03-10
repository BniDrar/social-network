package server

import "net/http"

// handler used  for status-checking or uptime monitoring of your server.
func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("OK"))
}
