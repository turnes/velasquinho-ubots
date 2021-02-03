package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


type App struct {
	Router *mux.Router
}

type Message struct {
	Message string
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()



	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatalln(http.ListenAndServe(addr, a.Router))
}


func (a *App) InitializeRoutes(){
	a.Router.HandleFunc("/health", a.health ).Methods("GET")
	a.Router.HandleFunc("/api/v1/orders/report/all", a.getSpendingByClient ).Methods("GET")
	a.Router.HandleFunc("/api/v1/orders/report/year/{year:[0-9]{4}}", a.getSpendingByYear ).Methods("GET")


}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	log.Fatalln(w.Write(response))
}

func (a *App) getSpendingByClient(w http.ResponseWriter, r *http.Request) {
	m := Message{
		"Hello World",
	}
	respondWithJSON(w, http.StatusOK, m)

}

func (a *App) getSpendingByYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid year")
		return
	}
	m := Message{
		strconv.Itoa(year),
	}
	respondWithJSON(w, http.StatusOK, m)
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w,http.StatusOK, nil)
}





