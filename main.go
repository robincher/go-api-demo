package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robincher/go-api-example/common"
	configClass "github.com/robincher/go-api-example/config"
	daoClass "github.com/robincher/go-api-example/dao"
	model "github.com/robincher/go-api-example/model"
	"gopkg.in/mgo.v2/bson"
)

var people []model.Person
var movies []model.Movie
var movie model.Movie

var config = configClass.DBConfig{}
var dao = daoClass.MoviesDAO{}

func listAllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, movies)
}

func readMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newMovie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload, please try again")
		return
	}
	//Assign with an ID
	newMovie.ID = bson.NewObjectId()
	if err := dao.Insert(newMovie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var updatedMovie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload, please try again")
		return
	}

	if err := dao.Update(updatedMovie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.Delete(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Main function
func main() {
	router := mux.NewRouter()

	session := common.GetMongoSession()
	defer session.Close()

	//Mock Response for People
	people = append(people, model.Person{ID: "1", Firstname: "John", Lastname: "Cena", Address: &model.Address{City: "City X", State: "State X"}})
	people = append(people, model.Person{ID: "2", Firstname: "Koko", Lastname: "Momo", Address: &model.Address{City: "City Z", State: "State Y"}})

	//Handle People Routes
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	//Handle Movie Routes
	router.HandleFunc("/movies", listAllMovies).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", readMovie).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
