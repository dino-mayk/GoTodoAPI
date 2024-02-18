package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Activity struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Person *Person `json:"person"`
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

func deleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, activity := range activities {
		if activity.ID == params["id"] {
			activities = append(activities[:index], activities[index + 1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(activities)
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, activity := range activities {
		if activity.ID == params["id"] {
			json.NewEncoder(w).Encode(activity)
			return
		}
	} 
}

func createActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activity Activity
	_ = json.NewDecoder(r.Body).Decode(&activity)
	activity.ID = strconv.Itoa(rand.Intn(1000000000))
	activities = append(activities, activity)
	json.NewEncoder(w).Encode(activity)
}

func updateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, acactivity := range activities {
		if acactivity.ID == params["id"] {
			activities = append(activities, activities[index + 1:]...)
			var acactivity Activity
			_ = json.NewDecoder(r.Body).Decode(&acactivity)
			acactivity.ID = params["id"]
			activities = append(activities, acactivity)
			json.NewEncoder(w).Encode(acactivity)
		}
	}
}

func main() {
	r := mux.NewRouter()

	activities = append(activities, Activity{ID: "1", Title: "Activity One", Description: "play dotka", Person: &Person{Firstname: "Nikita", Lastname: "Krasnov"}})
	activities = append(activities, Activity{ID: "2", Title: "Activity Two", Description: "drink vodka", Person: &Person{Firstname: "Nikitiy", Lastname: "Makakin"}})

	r.HandleFunc("/activities", getActivities).Methods("GET")
	r.HandleFunc("/activities/{id}", getActivity).Methods("GET")
	r.HandleFunc("/activities", createActivity).Methods("POST")
	r.HandleFunc("/activities/{id}", updateActivity).Methods("PUT")
	r.HandleFunc("/activities/{id}", deleteActivity).Methods("DELETE")

	fmt.Printf("Server starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
