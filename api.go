package main

// aket pengolahan JSON dan protokol HTTP
import (
	"encoding/json"
	"net/http"
)

// deklarasi user
type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// inisialisasi user
var users = []user{
	{ID: "1", Name: "Heru Purnama", Email: "heru@example.com", Role: "Driver"},
	{ID: "2", Name: "Okky Rafa Nuggraha", Email: "okki@example.com", Role: "Direktur"},
}

// handler utk menangani permintaan GET ke /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// handler untk menangani permintaan POST ke /adduser
func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// membuat ID
	newUser.ID = "3"

	// menambahkan user baru
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/adduser", addUser)

	http.ListenAndServe(":8000", nil)
}
