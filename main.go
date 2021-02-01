package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type client struct {
	ID    int     `json:"id"`
	Name  string  `json:"nome"`
	CPF string `json:"cpf"`
}

type item struct {
	Product  string `json:"produto"`
	Type     string `json:"variedade"`
	Country  string `json:"pais"`
	Category string `json:"categoria"`
	Vintage  string `json:"safra"`
	Price    float64 `json:"preco"`
}

type order struct {
	Code string `json:"codigo"`
	Date string `json:"data"`
	Client string `json:"cliente"`
	Items []item `json:"itens"`
	SubTotal float64 `json:"valorTotal"`
}

func main() {
	url := "http://www.mocky.io/v2/598b16291100004705515ec5"
	url2 := "http://www.mocky.io/v2/598b16861100004905515ec7"

	apiClient := http.Client{
		Timeout: time.Second * 2,
	}
	clients := getClient(&apiClient, url)

	for _, c := range clients {
		fmt.Printf("ID: %d\tNome: %s\tCPF: %s\n", c.ID, c.Name, c.CPF)
	}

	orders := getOrders(&apiClient, url2)

	for _, o := range orders {
		fmt.Printf("Codigo: %s, Data: %s, Cliente: %s, Valor: %.2f\n", o.Code, o.Date, o.Client, o.SubTotal)
		for _, i := range o.Items {
			fmt.Printf("Produto: %s, Pais: %s, Preco: %.2f \n", i.Product, i.Country, i.Price)
		}
		fmt.Println("\n")
	}

	a := App{}
	a.Initialize()
	a.Run(":8010")



}

func getClient(apiClient *http.Client, url string) []client {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var clients []client

	err = json.Unmarshal(body, &clients)
	if err != nil {
		log.Fatal(err)
	}
	return clients

}

func getOrders(apiClient *http.Client, url string) []order {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var orders []order

	err = json.Unmarshal(body, &orders)

	if err != nil {
		log.Fatal(err)
	}
	return orders

}