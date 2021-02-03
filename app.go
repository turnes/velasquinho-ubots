package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/turnes/velasquinho-ubots/api/report"
	"github.com/turnes/velasquinho-ubots/data"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router          *mux.Router
	VelasquinhoData *data.VelasquinhoData
}

type Message struct {
	Message string
}

func (a *App) Initialize() {
	OrdersURL := "http://www.mocky.io/v2/598b16861100004905515ec7"
	ClientUrl := "http://www.mocky.io/v2/598b16291100004705515ec5"
	a.VelasquinhoData = &data.VelasquinhoData{}
	a.VelasquinhoData.Initialize(ClientUrl, OrdersURL)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	a.VelasquinhoData.Run()
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Println(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

func (a *App) getSpendingByClient(w http.ResponseWriter, r *http.Request) {

	payload := report.AllTime(a.VelasquinhoData.Clients, a.VelasquinhoData.Orders)
	respondWithJSON(w, http.StatusOK, payload)

}

func (a *App) getSpendingByYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["year"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid year")
		return
	}
	payload := report.ByYear(a.VelasquinhoData.Clients, a.VelasquinhoData.Orders, vars["year"])
	respondWithJSON(w, http.StatusOK, payload)
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, nil)
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/health", a.health).Methods("GET")
	a.Router.HandleFunc("/api/v1/report/orders", a.getSpendingByClient).Methods("GET")
	a.Router.HandleFunc("/api/v1/report/orders/year/{year:[0-9]{4}}", a.getSpendingByYear).Methods("GET")
}
