package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Activity struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Person *Person `json:"director"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var activities []Activity

func getActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

func main() {
	r := mux.NewRouter()

	activities = append(activities, Activity{ID: "1", Title: "Activity One", Description: "play dotka", Person: &Person{Firstname: "Nikita", Lastname: "Krasnov"}})
	activities = append(activities, Activity{ID: "2", Title: "Activity Two", Description: "drink vodka", Person: &Person{Firstname: "Nikitiy", Lastname: "Makakin"}})

	r.HandleFunc("/activities", getActivities).Methods("GET")

	fmt.Printf("Server starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
